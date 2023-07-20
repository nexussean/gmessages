// mautrix-gmessages - A Matrix-Google Messages puppeting bridge.
// Copyright (C) 2023 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"strings"
	"sync"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/rs/zerolog"
	mutil "maunium.net/go/mautrix/util"
	"maunium.net/go/mautrix/util/variationselector"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/appservice"
	"maunium.net/go/mautrix/bridge"
	"maunium.net/go/mautrix/crypto/attachment"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"

	"go.mau.fi/mautrix-gmessages/database"
	"go.mau.fi/mautrix-gmessages/libgm"
	"go.mau.fi/mautrix-gmessages/libgm/gmproto"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

func (br *GMBridge) GetPortalByMXID(mxid id.RoomID) *Portal {
	br.portalsLock.Lock()
	defer br.portalsLock.Unlock()
	portal, ok := br.portalsByMXID[mxid]
	if !ok {
		dbPortal, err := br.DB.Portal.GetByMXID(context.TODO(), mxid)
		if err != nil {
			br.ZLog.Err(err).Str("mxid", mxid.String()).Msg("Failed to get portal from database")
			return nil
		}
		return br.loadDBPortal(dbPortal, nil)
	}
	return portal
}

func (br *GMBridge) GetIPortal(mxid id.RoomID) bridge.Portal {
	p := br.GetPortalByMXID(mxid)
	if p == nil {
		return nil
	}
	return p
}

func (portal *Portal) IsEncrypted() bool {
	return portal.Encrypted
}

func (portal *Portal) MarkEncrypted() {
	portal.Encrypted = true
	err := portal.Update(context.TODO())
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to save portal to database after marking it as encrypted")
	}
}

func (portal *Portal) ReceiveMatrixEvent(brUser bridge.User, evt *event.Event) {
	user := brUser.(*User)
	if user.RowID == portal.Receiver {
		portal.matrixMessages <- PortalMatrixMessage{user: user, evt: evt, receivedAt: time.Now()}
	}
}

func (br *GMBridge) GetPortalByKey(key database.Key) *Portal {
	br.portalsLock.Lock()
	defer br.portalsLock.Unlock()
	portal, ok := br.portalsByKey[key]
	if !ok {
		dbPortal, err := br.DB.Portal.GetByKey(context.TODO(), key)
		if err != nil {
			br.ZLog.Err(err).Object("portal_key", key).Msg("Failed to get portal from database")
			return nil
		}
		return br.loadDBPortal(dbPortal, &key)
	}
	return portal
}

func (br *GMBridge) GetExistingPortalByKey(key database.Key) *Portal {
	br.portalsLock.Lock()
	defer br.portalsLock.Unlock()
	portal, ok := br.portalsByKey[key]
	if !ok {
		dbPortal, err := br.DB.Portal.GetByKey(context.TODO(), key)
		if err != nil {
			br.ZLog.Err(err).Object("portal_key", key).Msg("Failed to get portal from database")
			return nil
		}
		return br.loadDBPortal(dbPortal, nil)
	}
	return portal
}

func (br *GMBridge) GetAllPortals() []*Portal {
	return br.loadManyPortals(br.DB.Portal.GetAll)
}

func (br *GMBridge) GetAllPortalsForUser(userID int) []*Portal {
	return br.loadManyPortals(func(ctx context.Context) ([]*database.Portal, error) {
		return br.DB.Portal.GetAllForUser(ctx, userID)
	})
}

func (br *GMBridge) GetAllIPortals() (iportals []bridge.Portal) {
	portals := br.GetAllPortals()
	iportals = make([]bridge.Portal, len(portals))
	for i, portal := range portals {
		iportals[i] = portal
	}
	return iportals
}

func (br *GMBridge) loadManyPortals(query func(ctx context.Context) ([]*database.Portal, error)) []*Portal {
	br.portalsLock.Lock()
	defer br.portalsLock.Unlock()
	dbPortals, err := query(context.TODO())
	if err != nil {
		br.ZLog.Err(err).Msg("Failed to load all portals from database")
		return []*Portal{}
	}
	output := make([]*Portal, len(dbPortals))
	for index, dbPortal := range dbPortals {
		if dbPortal == nil {
			continue
		}
		portal, ok := br.portalsByKey[dbPortal.Key]
		if !ok {
			portal = br.loadDBPortal(dbPortal, nil)
		}
		output[index] = portal
	}
	return output
}

func (br *GMBridge) loadDBPortal(dbPortal *database.Portal, key *database.Key) *Portal {
	if dbPortal == nil {
		if key == nil {
			return nil
		}
		dbPortal = br.DB.Portal.New()
		dbPortal.Key = *key
		err := dbPortal.Insert(context.TODO())
		if err != nil {
			br.ZLog.Err(err).Object("portal_key", key).Msg("Failed to insert portal")
			return nil
		}
	}
	portal := br.NewPortal(dbPortal)
	br.portalsByKey[portal.Key] = portal
	if len(portal.MXID) > 0 {
		br.portalsByMXID[portal.MXID] = portal
	}
	return portal
}

func (portal *Portal) GetUsers() []*User {
	return nil
}

func (br *GMBridge) newBlankPortal(key database.Key) *Portal {
	portal := &Portal{
		bridge: br,
		zlog:   br.ZLog.With().Str("portal_id", key.ID).Int("portal_receiver", key.Receiver).Logger(),

		messages:       make(chan PortalMessage, br.Config.Bridge.PortalMessageBuffer),
		matrixMessages: make(chan PortalMatrixMessage, br.Config.Bridge.PortalMessageBuffer),

		outgoingMessages: make(map[string]*outgoingMessage),
	}
	go portal.handleMessageLoop()
	return portal
}

func (br *GMBridge) NewPortal(dbPortal *database.Portal) *Portal {
	portal := br.newBlankPortal(dbPortal.Key)
	portal.Portal = dbPortal
	return portal
}

const recentlyHandledLength = 100

type PortalMessage struct {
	evt    *gmproto.Message
	source *User
}

type PortalMatrixMessage struct {
	evt        *event.Event
	user       *User
	receivedAt time.Time
}

type outgoingMessage struct {
	*event.Event
	Saved        bool
	Checkpointed bool
}

type Portal struct {
	*database.Portal

	bridge *GMBridge
	zlog   zerolog.Logger

	roomCreateLock sync.Mutex
	encryptLock    sync.Mutex
	backfillLock   sync.Mutex
	avatarLock     sync.Mutex

	forwardBackfillLock sync.Mutex
	lastMessageTS       time.Time

	recentlyHandled      [recentlyHandledLength]string
	recentlyHandledLock  sync.Mutex
	recentlyHandledIndex uint8

	outgoingMessages     map[string]*outgoingMessage
	outgoingMessagesLock sync.Mutex

	currentlyTyping     []id.UserID
	currentlyTypingLock sync.Mutex

	messages       chan PortalMessage
	matrixMessages chan PortalMatrixMessage
}

var (
	_ bridge.Portal                    = (*Portal)(nil)
	_ bridge.ReadReceiptHandlingPortal = (*Portal)(nil)
	//_ bridge.MembershipHandlingPortal  = (*Portal)(nil)
	//_ bridge.MetaHandlingPortal        = (*Portal)(nil)
	//_ bridge.TypingPortal              = (*Portal)(nil)
)

func (portal *Portal) handleMessageLoopItem(msg PortalMessage) {
	if len(portal.MXID) == 0 {
		return
	}
	portal.forwardBackfillLock.Lock()
	defer portal.forwardBackfillLock.Unlock()
	switch {
	case msg.evt != nil:
		portal.handleMessage(msg.source, msg.evt)
	//case msg.receipt != nil:
	//	portal.handleReceipt(msg.receipt, msg.source)
	default:
		portal.zlog.Warn().Interface("portal_message", msg).Msg("Unexpected PortalMessage with no message")
	}
}

func (portal *Portal) handleMatrixMessageLoopItem(msg PortalMatrixMessage) {
	portal.forwardBackfillLock.Lock()
	defer portal.forwardBackfillLock.Unlock()
	evtTS := time.UnixMilli(msg.evt.Timestamp)
	timings := messageTimings{
		initReceive:  msg.evt.Mautrix.ReceivedAt.Sub(evtTS),
		decrypt:      msg.evt.Mautrix.DecryptionDuration,
		portalQueue:  time.Since(msg.receivedAt),
		totalReceive: time.Since(evtTS),
	}
	switch msg.evt.Type {
	case event.EventMessage, event.EventSticker:
		portal.HandleMatrixMessage(msg.user, msg.evt, timings)
	case event.EventReaction:
		portal.HandleMatrixReaction(msg.user, msg.evt)
	case event.EventRedaction:
		portal.HandleMatrixRedaction(msg.user, msg.evt)
	default:
		portal.zlog.Warn().
			Str("event_type", msg.evt.Type.Type).
			Msg("Unsupported event type in portal message channel")
	}
}

func (portal *Portal) handleMessageLoop() {
	for {
		select {
		case msg := <-portal.messages:
			portal.handleMessageLoopItem(msg)
		case msg := <-portal.matrixMessages:
			portal.handleMatrixMessageLoopItem(msg)
		}
	}
}

func (portal *Portal) isOutgoingMessage(msg *gmproto.Message) id.EventID {
	portal.outgoingMessagesLock.Lock()
	defer portal.outgoingMessagesLock.Unlock()
	out, ok := portal.outgoingMessages[msg.TmpID]
	if ok {
		if !out.Saved {
			portal.markHandled(&ConvertedMessage{
				ID:        msg.MessageID,
				Timestamp: time.UnixMicro(msg.GetTimestamp()),
				SenderID:  msg.ParticipantID,
			}, out.ID, true)
			out.Saved = true
		}
		switch msg.GetMessageStatus().GetStatus() {
		case gmproto.MessageStatusType_OUTGOING_DELIVERED, gmproto.MessageStatusType_OUTGOING_COMPLETE, gmproto.MessageStatusType_OUTGOING_DISPLAYED:
			delete(portal.outgoingMessages, msg.TmpID)
			go portal.sendStatusEvent(out.ID, "", nil)
		case gmproto.MessageStatusType_OUTGOING_FAILED_GENERIC,
			gmproto.MessageStatusType_OUTGOING_FAILED_EMERGENCY_NUMBER,
			gmproto.MessageStatusType_OUTGOING_CANCELED,
			gmproto.MessageStatusType_OUTGOING_FAILED_TOO_LARGE,
			gmproto.MessageStatusType_OUTGOING_FAILED_RECIPIENT_LOST_RCS,
			gmproto.MessageStatusType_OUTGOING_FAILED_NO_RETRY_NO_FALLBACK,
			gmproto.MessageStatusType_OUTGOING_FAILED_RECIPIENT_DID_NOT_DECRYPT,
			gmproto.MessageStatusType_OUTGOING_FAILED_RECIPIENT_LOST_ENCRYPTION,
			gmproto.MessageStatusType_OUTGOING_FAILED_RECIPIENT_DID_NOT_DECRYPT_NO_MORE_RETRY:
			err := OutgoingStatusError(msg.GetMessageStatus().GetStatus())
			go portal.sendStatusEvent(out.ID, "", err)
			// TODO error notice
		}
		return out.ID
	}
	return ""
}
func hasInProgressMedia(msg *gmproto.Message) bool {
	for _, part := range msg.MessageInfo {
		media, ok := part.GetData().(*gmproto.MessageInfo_MediaContent)
		if ok && media.MediaContent.GetMediaID() == "" {
			return true
		}
	}
	return false
}

func (portal *Portal) handleMessage(source *User, evt *gmproto.Message) {
	if len(portal.MXID) == 0 {
		portal.zlog.Warn().Msg("handleMessage called even though portal.MXID is empty")
		return
	}
	eventTS := time.UnixMicro(evt.GetTimestamp())
	if eventTS.After(portal.lastMessageTS) {
		portal.lastMessageTS = eventTS
	}
	log := portal.zlog.With().
		Str("message_id", evt.MessageID).
		Str("participant_id", evt.ParticipantID).
		Str("status", evt.GetMessageStatus().GetStatus().String()).
		Str("action", "handle google message").
		Logger()
	ctx := log.WithContext(context.TODO())
	switch evt.GetMessageStatus().GetStatus() {
	case gmproto.MessageStatusType_INCOMING_AUTO_DOWNLOADING, gmproto.MessageStatusType_INCOMING_RETRYING_AUTO_DOWNLOAD:
		log.Debug().Msg("Not handling incoming message that is auto downloading")
		return
	case gmproto.MessageStatusType_MESSAGE_DELETED:
		portal.handleGoogleDeletion(ctx, evt.MessageID)
		return
	}
	if hasInProgressMedia(evt) {
		log.Debug().Msg("Not handling incoming message that doesn't have full media yet")
		return
	}
	if evtID := portal.isOutgoingMessage(evt); evtID != "" {
		log.Debug().Str("event_id", evtID.String()).Msg("Got echo for outgoing message")
		return
	}
	existingMsg, err := portal.bridge.DB.Message.GetByID(ctx, portal.Key, evt.MessageID)
	if err != nil {
		log.Err(err).Msg("Failed to check if message is duplicate")
	} else if existingMsg != nil {
		log.Debug().Msg("Not handling duplicate message")
		portal.syncReactions(ctx, source, existingMsg, evt.Reactions)
		return
	}

	converted := portal.convertGoogleMessage(ctx, source, evt, false)
	if converted == nil {
		return
	}

	eventIDs := make([]id.EventID, 0, len(converted.Parts))
	for _, part := range converted.Parts {
		resp, err := portal.sendMessage(converted.Intent, event.EventMessage, part.Content, part.Extra, converted.Timestamp.UnixMilli())
		if err != nil {
			log.Err(err).Msg("Failed to send message")
		} else {
			eventIDs = append(eventIDs, resp.EventID)
		}
	}
	portal.markHandled(converted, eventIDs[0], true)
	portal.sendDeliveryReceipt(eventIDs[len(eventIDs)-1])
	log.Debug().Interface("event_ids", eventIDs).Msg("Handled message")
}

func (portal *Portal) handleGoogleDeletion(ctx context.Context, messageID string) {
	log := zerolog.Ctx(ctx)
	msg, err := portal.bridge.DB.Message.GetByID(ctx, portal.Key, messageID)
	if err != nil {
		log.Err(err).Msg("Failed to get deleted message from database")
	} else if msg == nil {
		log.Debug().Msg("Didn't find deleted message in database")
	} else {
		if _, err = portal.MainIntent().RedactEvent(portal.MXID, msg.MXID); err != nil {
			log.Err(err).Msg("Faield to redact deleted message")
		}
		if err = msg.Delete(ctx); err != nil {
			log.Err(err).Msg("Failed to delete message from database")
		}
		log.Debug().Msg("Handled message deletion")
	}
}

func (portal *Portal) syncReactions(ctx context.Context, source *User, message *database.Message, reactions []*gmproto.ReactionEntry) {
	log := zerolog.Ctx(ctx)
	existing, err := portal.bridge.DB.Reaction.GetAllByMessage(ctx, portal.Key, message.ID)
	if err != nil {
		log.Err(err).Msg("Failed to get existing reactions from db to sync reactions")
		return
	}
	remove := make(map[string]*database.Reaction, len(existing))
	for _, reaction := range existing {
		remove[reaction.Sender] = reaction
	}
	for _, reaction := range reactions {
		emoji := reaction.GetData().GetUnicode()
		if emoji == "" {
			emoji = reaction.GetData().GetType().Unicode()
			if emoji == "" {
				continue
			}
		}
		for _, participant := range reaction.GetParticipantIDs() {
			dbReaction, ok := remove[participant]
			if ok {
				delete(remove, participant)
				if dbReaction.Reaction == emoji {
					continue
				}
			}
			intent := portal.getIntent(ctx, source, participant)
			if intent == nil {
				continue
			}
			var resp *mautrix.RespSendEvent
			resp, err = intent.SendReaction(portal.MXID, message.MXID, variationselector.Add(emoji))
			if err != nil {
				log.Err(err).Str("reaction_sender_id", participant).Msg("Failed to send reaction")
				continue
			}
			if dbReaction == nil {
				dbReaction = portal.bridge.DB.Reaction.New()
				dbReaction.Chat = portal.Key
				dbReaction.Sender = participant
				dbReaction.MessageID = message.ID
			} else if _, err = intent.RedactEvent(portal.MXID, dbReaction.MXID); err != nil {
				log.Err(err).Str("reaction_sender_id", participant).Msg("Failed to redact old reaction after adding new one")
			}
			dbReaction.Reaction = emoji
			dbReaction.MXID = resp.EventID
			err = dbReaction.Insert(ctx)
			if err != nil {
				log.Err(err).Str("reaction_sender_id", participant).Msg("Failed to insert added reaction into db")
			}
		}
	}
	for _, reaction := range remove {
		intent := portal.getIntent(ctx, source, reaction.Sender)
		if intent == nil {
			continue
		}
		_, err = intent.RedactEvent(portal.MXID, reaction.MXID)
		if err != nil {
			log.Err(err).Str("reaction_sender_id", reaction.Sender).Msg("Failed to redact removed reaction")
		} else if err = reaction.Delete(ctx); err != nil {
			log.Err(err).Str("reaction_sender_id", reaction.Sender).Msg("Failed to remove reaction from db")
		}
	}
}

type ConvertedMessagePart struct {
	Content *event.MessageEventContent
	Extra   map[string]any
}

type ConvertedMessage struct {
	ID       string
	SenderID string

	Intent    *appservice.IntentAPI
	Timestamp time.Time
	ReplyTo   string
	Parts     []ConvertedMessagePart
}

func (portal *Portal) getIntent(ctx context.Context, source *User, participant string) *appservice.IntentAPI {
	if source.IsSelfParticipantID(participant) {
		intent := source.DoublePuppetIntent
		if intent == nil {
			zerolog.Ctx(ctx).Debug().Msg("Dropping message from self as double puppeting is not enabled")
			return nil
		}
		return intent
	} else {
		puppet := source.GetPuppetByID(participant, "")
		if puppet == nil {
			zerolog.Ctx(ctx).Debug().Str("participant_id", participant).Msg("Dropping message from unknown participant")
			return nil
		}
		return puppet.IntentFor(portal)
	}
}

func (portal *Portal) convertGoogleMessage(ctx context.Context, source *User, evt *gmproto.Message, backfill bool) *ConvertedMessage {
	log := zerolog.Ctx(ctx)

	var cm ConvertedMessage
	cm.SenderID = evt.ParticipantID
	cm.ID = evt.MessageID
	cm.Timestamp = time.UnixMicro(evt.Timestamp)
	cm.Intent = portal.getIntent(ctx, source, evt.ParticipantID)
	if cm.Intent == nil {
		return nil
	}

	var replyTo id.EventID
	if evt.GetReplyMessage() != nil {
		cm.ReplyTo = evt.GetReplyMessage().GetMessageID()
		msg, err := portal.bridge.DB.Message.GetByID(ctx, portal.Key, cm.ReplyTo)
		if err != nil {
			log.Err(err).Str("reply_to_id", cm.ReplyTo).Msg("Failed to get reply target message")
		} else if msg == nil {
			if backfill {
				replyTo = portal.deterministicEventID(cm.ReplyTo, 0)
			} else {
				log.Warn().Str("reply_to_id", cm.ReplyTo).Msg("Reply target message not found")
			}
		} else {
			replyTo = msg.MXID
		}
	}

	subject := evt.GetSubject()
	for _, part := range evt.MessageInfo {
		var content event.MessageEventContent
		switch data := part.GetData().(type) {
		case *gmproto.MessageInfo_MessageContent:
			content = event.MessageEventContent{
				MsgType: event.MsgText,
				Body:    data.MessageContent.GetContent(),
			}
			if subject != "" {
				content.Format = event.FormatHTML
				content.FormattedBody = fmt.Sprintf("<strong>%s</strong><br>%s", event.TextToHTML(subject), event.TextToHTML(content.Body))
				content.Body = fmt.Sprintf("**%s**\n%s", subject, content.Body)
				subject = ""
			}
		case *gmproto.MessageInfo_MediaContent:
			contentPtr, err := portal.convertGoogleMedia(source, cm.Intent, data.MediaContent)
			if err != nil {
				log.Err(err).Msg("Failed to copy attachment")
				content = event.MessageEventContent{
					MsgType: event.MsgNotice,
					Body:    fmt.Sprintf("Failed to transfer attachment %s", data.MediaContent.GetMediaName()),
				}
			} else {
				content = *contentPtr
			}
		}
		if replyTo != "" {
			content.RelatesTo = &event.RelatesTo{InReplyTo: &event.InReplyTo{EventID: replyTo}}
		}
		cm.Parts = append(cm.Parts, ConvertedMessagePart{
			Content: &content,
		})
	}
	if subject != "" {
		cm.Parts = append(cm.Parts, ConvertedMessagePart{
			Content: &event.MessageEventContent{
				MsgType: event.MsgText,
				Body:    subject,
			},
		})
	}
	if portal.bridge.Config.Bridge.CaptionInMessage {
		cm.MergeCaption()
	}
	return &cm
}

func (msg *ConvertedMessage) MergeCaption() {
	if len(msg.Parts) != 2 {
		return
	}

	var textPart, filePart ConvertedMessagePart
	if msg.Parts[0].Content.MsgType == event.MsgText {
		textPart = msg.Parts[0]
		filePart = msg.Parts[1]
	} else {
		textPart = msg.Parts[1]
		filePart = msg.Parts[0]
	}

	if textPart.Content.MsgType != event.MsgText {
		return
	}
	switch filePart.Content.MsgType {
	case event.MsgImage, event.MsgVideo, event.MsgAudio, event.MsgFile:
	default:
		return
	}

	filePart.Content.FileName = filePart.Content.Body
	filePart.Content.Body = textPart.Content.Body
	filePart.Content.Format = textPart.Content.Format
	filePart.Content.FormattedBody = textPart.Content.FormattedBody
	msg.Parts = []ConvertedMessagePart{filePart}
}

func (portal *Portal) convertGoogleMedia(source *User, intent *appservice.IntentAPI, msg *gmproto.MediaContent) (*event.MessageEventContent, error) {
	var data []byte
	var err error
	data, err = source.Client.DownloadMedia(msg.MediaID, msg.DecryptionKey)
	if err != nil {
		return nil, err
	}
	mime := libgm.FormatToMediaType[msg.GetFormat()].Format
	if mime == "" {
		mime = mimetype.Detect(data).String()
	}
	msgtype := event.MsgFile
	switch strings.Split(mime, "/")[0] {
	case "image":
		msgtype = event.MsgImage
	case "video":
		msgtype = event.MsgVideo
		// TODO convert weird formats to mp4
	case "audio":
		msgtype = event.MsgAudio
		// TODO convert everything to ogg and include voice message metadata
	}
	content := &event.MessageEventContent{
		MsgType: msgtype,
		Body:    msg.MediaName,
		Info: &event.FileInfo{
			MimeType: mime,
			Size:     len(data),
		},
	}
	return content, portal.uploadMedia(intent, data, content)
}

func (portal *Portal) isRecentlyHandled(id string) bool {
	start := portal.recentlyHandledIndex
	for i := start; i != start; i = (i - 1) % recentlyHandledLength {
		if portal.recentlyHandled[i] == id {
			return true
		}
	}
	return false
}

func (portal *Portal) markHandled(cm *ConvertedMessage, eventID id.EventID, recent bool) *database.Message {
	msg := portal.bridge.DB.Message.New()
	msg.Chat = portal.Key
	msg.ID = cm.ID
	msg.MXID = eventID
	msg.Timestamp = cm.Timestamp
	msg.Sender = cm.SenderID
	err := msg.Insert(context.TODO())
	if err != nil {
		portal.zlog.Err(err).Str("message_id", cm.ID).Msg("Failed to insert message to database")
	}

	if recent {
		portal.recentlyHandledLock.Lock()
		index := portal.recentlyHandledIndex
		portal.recentlyHandledIndex = (portal.recentlyHandledIndex + 1) % recentlyHandledLength
		portal.recentlyHandledLock.Unlock()
		portal.recentlyHandled[index] = cm.ID
	}
	return msg
}

func (portal *Portal) SyncParticipants(source *User, metadata *gmproto.Conversation) (userIDs []id.UserID, changed bool) {
	var firstParticipant *gmproto.Participant
	var manyParticipants bool
	for _, participant := range metadata.Participants {
		if participant.IsMe {
			err := source.AddSelfParticipantID(context.TODO(), participant.ID.ParticipantID)
			if err != nil {
				portal.zlog.Warn().Err(err).
					Str("participant_id", participant.ID.ParticipantID).
					Msg("Failed to save self participant ID")
			}
			continue
		} else if participant.ID.Number == "" {
			portal.zlog.Warn().Interface("participant", participant).Msg("No number found in non-self participant entry")
			continue
		}
		if firstParticipant == nil {
			firstParticipant = participant
		} else {
			manyParticipants = true
		}
		portal.zlog.Debug().Interface("participant", participant).Msg("Syncing participant")
		puppet := source.GetPuppetByID(participant.ID.ParticipantID, participant.ID.Number)
		if puppet == nil {
			portal.zlog.Error().Any("participant_id", participant.ID).Msg("Failed to get puppet for participant")
			continue
		}
		userIDs = append(userIDs, puppet.MXID)
		puppet.Sync(source, participant)
		if portal.MXID != "" {
			err := puppet.IntentFor(portal).EnsureJoined(portal.MXID)
			if err != nil {
				portal.zlog.Err(err).
					Str("user_id", puppet.MXID.String()).
					Msg("Failed to ensure ghost is joined to portal")
			}
		}
	}
	if !metadata.IsGroupChat && !manyParticipants && portal.OtherUserID != firstParticipant.ID.ParticipantID {
		portal.zlog.Info().
			Str("old_other_user_id", portal.OtherUserID).
			Str("new_other_user_id", firstParticipant.ID.ParticipantID).
			Msg("Found other user ID in DM")
		portal.OtherUserID = firstParticipant.ID.ParticipantID
		changed = true
	}
	return userIDs, changed
}

func (portal *Portal) UpdateName(name string, updateInfo bool) bool {
	if portal.Name != name || (!portal.NameSet && len(portal.MXID) > 0 && portal.shouldSetDMRoomMetadata()) {
		portal.zlog.Debug().Str("old_name", portal.Name).Str("new_name", name).Msg("Updating name")
		portal.Name = name
		portal.NameSet = false
		if updateInfo {
			defer func() {
				err := portal.Update(context.TODO())
				if err != nil {
					portal.zlog.Err(err).Msg("Failed to save portal after updating name")
				}
			}()
		}

		if len(portal.MXID) > 0 && !portal.shouldSetDMRoomMetadata() {
			portal.UpdateBridgeInfo()
		} else if len(portal.MXID) > 0 {
			intent := portal.MainIntent()
			_, err := intent.SetRoomName(portal.MXID, name)
			if errors.Is(err, mautrix.MForbidden) && intent != portal.MainIntent() {
				_, err = portal.MainIntent().SetRoomName(portal.MXID, name)
			}
			if err == nil {
				portal.NameSet = true
				if updateInfo {
					portal.UpdateBridgeInfo()
				}
				return true
			} else {
				portal.zlog.Warn().Err(err).Msg("Failed to set room name")
			}
		}
	}
	return false
}

func (portal *Portal) UpdateMetadata(user *User, info *gmproto.Conversation) []id.UserID {
	participants, update := portal.SyncParticipants(user, info)
	if portal.OutgoingID != info.DefaultOutgoingID {
		portal.OutgoingID = info.DefaultOutgoingID
		update = true
	}
	if portal.MXID != "" {
		update = portal.addToPersonalSpace(user, false) || update
	}
	if portal.shouldSetDMRoomMetadata() {
		update = portal.UpdateName(info.Name, false) || update
	}
	if portal.MXID != "" {
		pls, err := portal.MainIntent().PowerLevels(portal.MXID)
		if err != nil {
			portal.zlog.Warn().Err(err).Msg("Failed to get power levels")
		} else if portal.updatePowerLevels(info, pls) {
			resp, err := portal.MainIntent().SetPowerLevels(portal.MXID, pls)
			if err != nil {
				portal.zlog.Warn().Err(err).Msg("Failed to update power levels")
			} else {
				portal.zlog.Debug().Str("event_id", resp.EventID.String()).Msg("Updated power levels")
			}
		}
	}

	// TODO avatar
	if update {
		err := portal.Update(context.TODO())
		if err != nil {
			portal.zlog.Err(err).Msg("Failed to save portal after updating metadata")
		}
		if portal.MXID != "" {
			portal.UpdateBridgeInfo()
		}
	}
	return participants
}

func (portal *Portal) ensureUserInvited(user *User) bool {
	return user.ensureInvited(portal.MainIntent(), portal.MXID, portal.IsPrivateChat())
}

func (portal *Portal) UpdateMatrixRoom(user *User, groupInfo *gmproto.Conversation) bool {
	if len(portal.MXID) == 0 {
		return false
	}

	portal.ensureUserInvited(user)
	portal.UpdateMetadata(user, groupInfo)
	return true
}

func (portal *Portal) GetBasePowerLevels() *event.PowerLevelsEventContent {
	anyone := 0
	nope := 99
	return &event.PowerLevelsEventContent{
		UsersDefault:    anyone,
		EventsDefault:   anyone,
		RedactPtr:       &anyone,
		StateDefaultPtr: &nope,
		BanPtr:          &nope,
		KickPtr:         &nope,
		InvitePtr:       &nope,
		Users: map[id.UserID]int{
			portal.MainIntent().UserID: 100,
			portal.bridge.Bot.UserID:   100,
		},
		Events: map[string]int{
			event.StateRoomName.Type:   anyone,
			event.StateRoomAvatar.Type: anyone,
			event.EventReaction.Type:   anyone,
			event.EventRedaction.Type:  anyone,
		},
	}
}

func (portal *Portal) updatePowerLevels(conv *gmproto.Conversation, pl *event.PowerLevelsEventContent) bool {
	expectedEventsDefault := 0
	if conv.GetReadOnly() {
		expectedEventsDefault = 99
	}

	changed := false
	if pl.EventsDefault != expectedEventsDefault {
		pl.EventsDefault = expectedEventsDefault
		changed = true
	}
	changed = pl.EnsureEventLevel(event.EventReaction, expectedEventsDefault) || changed
	// Explicitly set m.room.redaction level to 0 so redactions work even if sending is disabled
	changed = pl.EnsureEventLevel(event.EventRedaction, 0) || changed
	changed = pl.EnsureUserLevel(portal.bridge.Bot.UserID, 100) || changed
	return changed
}

func (portal *Portal) getBridgeInfoStateKey() string {
	return fmt.Sprintf("fi.mau.gmessages://gmessages/%s", portal.ID)
}

func (portal *Portal) getBridgeInfo() (string, event.BridgeEventContent) {
	return portal.getBridgeInfoStateKey(), event.BridgeEventContent{
		BridgeBot: portal.bridge.Bot.UserID,
		Creator:   portal.MainIntent().UserID,
		Protocol: event.BridgeInfoSection{
			ID:          "gmessages",
			DisplayName: "Google Messages",
			AvatarURL:   portal.bridge.Config.AppService.Bot.ParsedAvatar.CUString(),
			ExternalURL: "https://messages.google.com/",
		},
		Channel: event.BridgeInfoSection{
			ID:          portal.ID,
			DisplayName: portal.Name,
			AvatarURL:   portal.AvatarMXC.CUString(),
		},
	}
}

func (portal *Portal) UpdateBridgeInfo() {
	if len(portal.MXID) == 0 {
		portal.zlog.Debug().Msg("Not updating bridge info: no Matrix room created")
		return
	}
	portal.zlog.Debug().Msg("Updating bridge info...")
	stateKey, content := portal.getBridgeInfo()
	_, err := portal.MainIntent().SendStateEvent(portal.MXID, event.StateBridge, stateKey, content)
	if err != nil {
		portal.zlog.Warn().Err(err).Msg("Failed to update m.bridge")
	}
	// TODO remove this once https://github.com/matrix-org/matrix-doc/pull/2346 is in spec
	_, err = portal.MainIntent().SendStateEvent(portal.MXID, event.StateHalfShotBridge, stateKey, content)
	if err != nil {
		portal.zlog.Warn().Err(err).Msg("Failed to update uk.half-shot.bridge")
	}
}

func (portal *Portal) shouldSetDMRoomMetadata() bool {
	return !portal.IsPrivateChat() ||
		portal.bridge.Config.Bridge.PrivateChatPortalMeta == "always" ||
		((portal.IsEncrypted() || (portal.MXID == "" && portal.bridge.Config.Bridge.Encryption.Default)) &&
			portal.bridge.Config.Bridge.PrivateChatPortalMeta != "never")
}

func (portal *Portal) GetEncryptionEventContent() (evt *event.EncryptionEventContent) {
	evt = &event.EncryptionEventContent{Algorithm: id.AlgorithmMegolmV1}
	if rot := portal.bridge.Config.Bridge.Encryption.Rotation; rot.EnableCustom {
		evt.RotationPeriodMillis = rot.Milliseconds
		evt.RotationPeriodMessages = rot.Messages
	}
	return
}

func (portal *Portal) CreateMatrixRoom(user *User, conv *gmproto.Conversation) error {
	portal.roomCreateLock.Lock()
	defer portal.roomCreateLock.Unlock()
	if len(portal.MXID) > 0 {
		return nil
	}

	members := portal.UpdateMetadata(user, conv)

	if portal.IsPrivateChat() && portal.GetDMPuppet() == nil {
		portal.zlog.Error().Msg("Didn't find ghost of other user in DM :(")
		return fmt.Errorf("ghost not found")
	}

	intent := portal.MainIntent()
	if err := intent.EnsureRegistered(); err != nil {
		return err
	}

	portal.zlog.Info().Msg("Creating Matrix room")

	bridgeInfoStateKey, bridgeInfo := portal.getBridgeInfo()

	pl := portal.GetBasePowerLevels()
	portal.updatePowerLevels(conv, pl)

	initialState := []*event.Event{{
		Type:    event.StatePowerLevels,
		Content: event.Content{Parsed: pl},
	}, {
		Type:     event.StateBridge,
		Content:  event.Content{Parsed: bridgeInfo},
		StateKey: &bridgeInfoStateKey,
	}, {
		// TODO remove this once https://github.com/matrix-org/matrix-doc/pull/2346 is in spec
		Type:     event.StateHalfShotBridge,
		Content:  event.Content{Parsed: bridgeInfo},
		StateKey: &bridgeInfoStateKey,
	}}
	var invite []id.UserID
	if portal.bridge.Config.Bridge.Encryption.Default {
		initialState = append(initialState, &event.Event{
			Type: event.StateEncryption,
			Content: event.Content{
				Parsed: portal.GetEncryptionEventContent(),
			},
		})
		portal.Encrypted = true
		if portal.IsPrivateChat() {
			invite = append(invite, portal.bridge.Bot.UserID)
		}
	}
	if !portal.AvatarMXC.IsEmpty() && portal.shouldSetDMRoomMetadata() {
		initialState = append(initialState, &event.Event{
			Type: event.StateRoomAvatar,
			Content: event.Content{
				Parsed: event.RoomAvatarEventContent{URL: portal.AvatarMXC},
			},
		})
		portal.AvatarSet = true
	} else {
		portal.AvatarSet = false
	}

	creationContent := make(map[string]interface{})
	if !portal.bridge.Config.Bridge.FederateRooms {
		creationContent["m.federate"] = false
	}
	autoJoinInvites := portal.bridge.SpecVersions.Supports(mautrix.BeeperFeatureAutojoinInvites)
	if autoJoinInvites {
		portal.zlog.Debug().Msg("Hungryserv mode: adding all group members in create request")
		invite = append(invite, members...)
		invite = append(invite, user.MXID)
	}
	req := &mautrix.ReqCreateRoom{
		Visibility:      "private",
		Name:            portal.Name,
		Invite:          invite,
		Preset:          "private_chat",
		IsDirect:        portal.IsPrivateChat(),
		InitialState:    initialState,
		CreationContent: creationContent,

		BeeperAutoJoinInvites: autoJoinInvites,
	}
	if !portal.shouldSetDMRoomMetadata() {
		req.Name = ""
	}
	resp, err := intent.CreateRoom(req)
	if err != nil {
		return err
	}
	portal.zlog.Info().Str("room_id", resp.RoomID.String()).Msg("Matrix room created")
	portal.InSpace = false
	portal.NameSet = len(req.Name) > 0
	portal.forwardBackfillLock.Lock()
	portal.MXID = resp.RoomID
	portal.bridge.portalsLock.Lock()
	portal.bridge.portalsByMXID[portal.MXID] = portal
	portal.bridge.portalsLock.Unlock()
	err = portal.Update(context.TODO())
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to save portal after creating room")
	}

	// We set the memberships beforehand to make sure the encryption key exchange in initial backfill knows the users are here.
	inviteMembership := event.MembershipInvite
	if autoJoinInvites {
		inviteMembership = event.MembershipJoin
	}
	for _, userID := range invite {
		portal.bridge.StateStore.SetMembership(portal.MXID, userID, inviteMembership)
	}

	if !autoJoinInvites {
		if !portal.IsPrivateChat() {
			portal.SyncParticipants(user, conv)
		} else {
			if portal.bridge.Config.Bridge.Encryption.Default {
				err = portal.bridge.Bot.EnsureJoined(portal.MXID)
				if err != nil {
					portal.zlog.Err(err).Msg("Failed to join created portal with bridge bot for e2be")
				}
			}

			user.UpdateDirectChats(map[id.UserID][]id.RoomID{portal.GetDMPuppet().MXID: {portal.MXID}})
		}
		portal.ensureUserInvited(user)
	}
	user.syncChatDoublePuppetDetails(portal, conv, true)
	go portal.initialForwardBackfill(user, !conv.GetUnread())
	go portal.addToPersonalSpace(user, true)
	return nil
}

func (portal *Portal) addToPersonalSpace(user *User, updateInfo bool) bool {
	spaceID := user.GetSpaceRoom()
	if len(spaceID) == 0 || portal.InSpace {
		return false
	}
	_, err := portal.bridge.Bot.SendStateEvent(spaceID, event.StateSpaceChild, portal.MXID.String(), &event.SpaceChildEventContent{
		Via: []string{portal.bridge.Config.Homeserver.Domain},
	})
	if err != nil {
		portal.zlog.Err(err).Str("space_id", spaceID.String()).Msg("Failed to add room to user's personal filtering space")
		portal.InSpace = false
	} else {
		portal.zlog.Debug().Str("space_id", spaceID.String()).Msg("Added room to user's personal filtering space")
		portal.InSpace = true
	}
	if updateInfo {
		err = portal.Update(context.TODO())
		if err != nil {
			portal.zlog.Err(err).Msg("Failed to update portal after adding to personal space")
		}
	}
	return true
}

func (portal *Portal) IsPrivateChat() bool {
	return portal.OtherUserID != ""
}

func (portal *Portal) GetDMPuppet() *Puppet {
	if portal.IsPrivateChat() {
		puppet := portal.bridge.GetPuppetByKey(database.Key{Receiver: portal.Receiver, ID: portal.OtherUserID}, "")
		return puppet
	}
	return nil
}

func (portal *Portal) MainIntent() *appservice.IntentAPI {
	if puppet := portal.GetDMPuppet(); puppet != nil {
		return puppet.DefaultIntent()
	}
	return portal.bridge.Bot
}

func (portal *Portal) sendMainIntentMessage(content *event.MessageEventContent) (*mautrix.RespSendEvent, error) {
	return portal.sendMessage(portal.MainIntent(), event.EventMessage, content, nil, 0)
}

func (portal *Portal) encrypt(intent *appservice.IntentAPI, content *event.Content, eventType event.Type) (event.Type, error) {
	if !portal.Encrypted || portal.bridge.Crypto == nil {
		return eventType, nil
	}
	intent.AddDoublePuppetValue(content)
	// TODO maybe the locking should be inside mautrix-go?
	portal.encryptLock.Lock()
	defer portal.encryptLock.Unlock()
	err := portal.bridge.Crypto.Encrypt(portal.MXID, eventType, content)
	if err != nil {
		return eventType, fmt.Errorf("failed to encrypt event: %w", err)
	}
	return event.EventEncrypted, nil
}

func (portal *Portal) sendMessage(intent *appservice.IntentAPI, eventType event.Type, content *event.MessageEventContent, extraContent map[string]interface{}, timestamp int64) (*mautrix.RespSendEvent, error) {
	wrappedContent := event.Content{Parsed: content, Raw: extraContent}
	var err error
	eventType, err = portal.encrypt(intent, &wrappedContent, eventType)
	if err != nil {
		return nil, err
	}

	_, _ = intent.UserTyping(portal.MXID, false, 0)
	if timestamp == 0 {
		return intent.SendMessageEvent(portal.MXID, eventType, &wrappedContent)
	} else {
		return intent.SendMassagedMessageEvent(portal.MXID, eventType, &wrappedContent, timestamp)
	}
}

func (portal *Portal) encryptFileInPlace(data []byte, mimeType string) (string, *event.EncryptedFileInfo) {
	if !portal.Encrypted {
		return mimeType, nil
	}

	file := &event.EncryptedFileInfo{
		EncryptedFile: *attachment.NewEncryptedFile(),
		URL:           "",
	}
	file.EncryptInPlace(data)
	return "application/octet-stream", file
}

func (portal *Portal) uploadMedia(intent *appservice.IntentAPI, data []byte, content *event.MessageEventContent) error {
	uploadMimeType, file := portal.encryptFileInPlace(data, content.Info.MimeType)

	req := mautrix.ReqUploadMedia{
		ContentBytes: data,
		ContentType:  uploadMimeType,
	}
	var mxc id.ContentURI
	if portal.bridge.Config.Homeserver.AsyncMedia {
		uploaded, err := intent.UploadAsync(req)
		if err != nil {
			return err
		}
		mxc = uploaded.ContentURI
	} else {
		uploaded, err := intent.UploadMedia(req)
		if err != nil {
			return err
		}
		mxc = uploaded.ContentURI
	}

	if file != nil {
		file.URL = mxc.CUString()
		content.File = file
	} else {
		content.URL = mxc.CUString()
	}

	content.Info.Size = len(data)
	if content.Info.Width == 0 && content.Info.Height == 0 && strings.HasPrefix(content.Info.MimeType, "image/") {
		cfg, _, _ := image.DecodeConfig(bytes.NewReader(data))
		content.Info.Width, content.Info.Height = cfg.Width, cfg.Height
	}

	// This is a hack for bad clients like Element iOS that require a thumbnail (https://github.com/vector-im/element-ios/issues/4004)
	if strings.HasPrefix(content.Info.MimeType, "image/") && content.Info.ThumbnailInfo == nil {
		infoCopy := *content.Info
		content.Info.ThumbnailInfo = &infoCopy
		if content.File != nil {
			content.Info.ThumbnailFile = file
		} else {
			content.Info.ThumbnailURL = content.URL
		}
	}
	return nil
}

func (portal *Portal) convertMatrixMessage(ctx context.Context, sender *User, content *event.MessageEventContent, txnID string) (*gmproto.SendMessageRequest, error) {
	log := zerolog.Ctx(ctx)
	req := &gmproto.SendMessageRequest{
		ConversationID: portal.ID,
		TmpID:          txnID,
		MessagePayload: &gmproto.MessagePayload{
			ConversationID: portal.ID,
			TmpID:          txnID,
			TmpID2:         txnID,
			ParticipantID:  portal.OutgoingID,
		},
	}

	replyToMXID := content.RelatesTo.GetReplyTo()
	if replyToMXID != "" {
		replyToMsg, err := portal.bridge.DB.Message.GetByMXID(ctx, replyToMXID)
		if err != nil {
			log.Err(err).Str("reply_to_mxid", replyToMXID.String()).Msg("Failed to get reply target message")
		} else if replyToMsg == nil {
			log.Warn().Str("reply_to_mxid", replyToMXID.String()).Msg("Reply target message not found")
		} else {
			req.IsReply = true
			req.Reply = &gmproto.ReplyPayload{MessageID: replyToMsg.ID}
		}
	}

	switch content.MsgType {
	case event.MsgText, event.MsgEmote, event.MsgNotice:
		text := content.Body
		if content.MsgType == event.MsgEmote {
			text = "/me " + text
		}
		req.MessagePayload.MessageInfo = []*gmproto.MessageInfo{{
			Data: &gmproto.MessageInfo_MessageContent{MessageContent: &gmproto.MessageContent{
				Content: text,
			}},
		}}
	case event.MsgImage, event.MsgVideo, event.MsgAudio, event.MsgFile:
		var url id.ContentURI
		if content.File != nil {
			url = content.File.URL.ParseOrIgnore()
		} else {
			url = content.URL.ParseOrIgnore()
		}
		if url.IsEmpty() {
			return nil, errMissingMediaURL
		}
		data, err := portal.MainIntent().DownloadBytesContext(ctx, url)
		if err != nil {
			return nil, mutil.NewDualError(errMediaDownloadFailed, err)
		}
		if content.File != nil {
			err = content.File.DecryptInPlace(data)
			if err != nil {
				return nil, mutil.NewDualError(errMediaDecryptFailed, err)
			}
		}
		if content.Info.MimeType == "" {
			content.Info.MimeType = mimetype.Detect(data).String()
		}
		fileName := content.Body
		if content.FileName != "" {
			fileName = content.FileName
		}
		resp, err := sender.Client.UploadMedia(data, fileName, content.Info.MimeType)
		if err != nil {
			return nil, mutil.NewDualError(errMediaReuploadFailed, err)
		}
		req.MessagePayload.MessageInfo = []*gmproto.MessageInfo{{
			Data: &gmproto.MessageInfo_MediaContent{MediaContent: resp},
		}}
	default:
		return nil, fmt.Errorf("%w %s", errUnknownMsgType, content.MsgType)
	}
	return req, nil
}

func (portal *Portal) HandleMatrixMessage(sender *User, evt *event.Event, timings messageTimings) {
	ms := metricSender{portal: portal, timings: &timings}

	log := portal.zlog.With().
		Str("event_id", evt.ID.String()).
		Str("action", "handle matrix message").
		Logger()
	ctx := log.WithContext(context.TODO())
	log.Debug().Dur("age", timings.totalReceive).Msg("Handling Matrix message")

	content, ok := evt.Content.Parsed.(*event.MessageEventContent)
	if !ok {
		return
	}

	txnID := util.GenerateTmpID()
	portal.outgoingMessagesLock.Lock()
	portal.outgoingMessages[txnID] = &outgoingMessage{Event: evt}
	portal.outgoingMessagesLock.Unlock()
	if evt.Type == event.EventSticker {
		content.MsgType = event.MsgImage
	}

	start := time.Now()
	req, err := portal.convertMatrixMessage(ctx, sender, content, txnID)
	timings.convert = time.Since(start)
	if err != nil {
		go ms.sendMessageMetrics(evt, err, "Error converting", true)
		return
	}
	start = time.Now()
	_, err = sender.Client.SendMessage(req)
	timings.send = time.Since(start)
	if err != nil {
		go ms.sendMessageMetrics(evt, err, "Error sending", true)
	} else {
		go ms.sendMessageMetrics(evt, nil, "", true)
	}
}

func (portal *Portal) HandleMatrixReadReceipt(brUser bridge.User, eventID id.EventID, receipt event.ReadReceipt) {
	user := brUser.(*User)
	log := portal.zlog.With().
		Str("event_id", eventID.String()).
		Time("receipt_ts", receipt.Timestamp).
		Str("action", "handle matrix read receipt").
		Logger()
	ctx := log.WithContext(context.TODO())
	log.Debug().Msg("Handling Matrix read receipt")
	targetMessage, err := portal.bridge.DB.Message.GetByMXID(ctx, eventID)
	if err != nil {
		log.Err(err).Msg("Failed to get target message to handle read receipt")
		return
	} else if targetMessage == nil {
		lastMessage, err := portal.bridge.DB.Message.GetLastInChat(ctx, portal.Key)
		if err != nil {
			log.Err(err).Msg("Failed to get last message to handle read receipt")
			return
		} else if receipt.Timestamp.Before(lastMessage.Timestamp) {
			log.Debug().Msg("Ignoring read receipt for unknown message with timestamp before last message")
			return
		} else {
			log.Debug().Msg("Marking last message in chat as read for receipt targeting unknown message")
		}
		targetMessage = lastMessage
	}
	log = log.With().Str("message_id", targetMessage.ID).Logger()
	err = user.Client.MarkRead(portal.ID, targetMessage.ID)
	if err != nil {
		log.Err(err).Msg("Failed to mark message as read")
	} else {
		log.Debug().Msg("Marked message as read after Matrix read receipt")
	}
}

func (portal *Portal) HandleMatrixReaction(sender *User, evt *event.Event) {
	err := portal.handleMatrixReaction(sender, evt)
	go portal.sendMessageMetrics(evt, err, "Error sending", nil)
}

func (portal *Portal) handleMatrixReaction(sender *User, evt *event.Event) error {
	content, ok := evt.Content.Parsed.(*event.ReactionEventContent)
	if !ok {
		return fmt.Errorf("unexpected parsed content type %T", evt.Content.Parsed)
	}
	log := portal.zlog.With().
		Str("event_id", evt.ID.String()).
		Str("target_event_id", content.RelatesTo.EventID.String()).
		Str("action", "handle matrix reaction").
		Logger()
	ctx := log.WithContext(context.TODO())
	log.Debug().Msg("Handling Matrix reaction")

	msg, err := portal.bridge.DB.Message.GetByMXID(ctx, content.RelatesTo.EventID)
	if err != nil {
		log.Err(err).Msg("Failed to get reaction target event")
		return fmt.Errorf("failed to get event from database")
	} else if msg == nil {
		return errTargetNotFound
	}

	existingReaction, err := portal.bridge.DB.Reaction.GetByID(ctx, portal.Key, msg.ID, portal.OutgoingID)
	if err != nil {
		log.Err(err).Msg("Failed to get existing reaction")
		return fmt.Errorf("failed to get existing reaction from database")
	}

	emoji := variationselector.Remove(content.RelatesTo.Key)
	action := gmproto.SendReactionRequest_ADD
	if existingReaction != nil {
		action = gmproto.SendReactionRequest_SWITCH
	}
	resp, err := sender.Client.SendReaction(&gmproto.SendReactionRequest{
		MessageID:    msg.ID,
		ReactionData: gmproto.MakeReactionData(emoji),
		Action:       action,
	})
	if err != nil {
		return fmt.Errorf("failed to send reaction: %w", err)
	} else if !resp.Success {
		return fmt.Errorf("got non-success response")
	}
	if existingReaction == nil {
		existingReaction = portal.bridge.DB.Reaction.New()
		existingReaction.Chat = portal.Key
		existingReaction.MessageID = msg.ID
		existingReaction.Sender = portal.OutgoingID
	} else if sender.DoublePuppetIntent != nil {
		_, err = sender.DoublePuppetIntent.RedactEvent(portal.MXID, existingReaction.MXID)
		if err != nil {
			log.Err(err).Msg("Failed to redact old reaction with double puppet after new Matrix reaction")
		}
	} else {
		_, err = portal.MainIntent().RedactEvent(portal.MXID, existingReaction.MXID)
		if err != nil {
			log.Err(err).Msg("Failed to redact old reaction with main intent after new Matrix reaction")
		}
	}
	existingReaction.Reaction = emoji
	existingReaction.MXID = evt.ID
	err = existingReaction.Insert(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to save reaction from Matrix to database")
	}
	return nil
}

func (portal *Portal) HandleMatrixRedaction(sender *User, evt *event.Event) {
	err := portal.handleMatrixRedaction(sender, evt)
	go portal.sendMessageMetrics(evt, err, "Error sending", nil)
}

func (portal *Portal) handleMatrixMessageRedaction(ctx context.Context, sender *User, redacts id.EventID) error {
	log := zerolog.Ctx(ctx)
	msg, err := portal.bridge.DB.Message.GetByMXID(ctx, redacts)
	if err != nil {
		log.Err(err).Msg("Failed to get redaction target message")
		return fmt.Errorf("failed to get event from database")
	} else if msg == nil {
		return errTargetNotFound
	}
	resp, err := sender.Client.DeleteMessage(msg.ID)
	if err != nil {
		return fmt.Errorf("failed to send message removal: %w", err)
	} else if !resp.Success {
		return fmt.Errorf("got non-success response")
	}
	err = msg.Delete(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to delete message from database after Matrix redaction")
	}
	return nil
}

func (portal *Portal) handleMatrixReactionRedaction(ctx context.Context, sender *User, redacts id.EventID) error {
	log := zerolog.Ctx(ctx)
	existingReaction, err := portal.bridge.DB.Reaction.GetByMXID(ctx, redacts)
	if err != nil {
		log.Err(err).Msg("Failed to get redaction target reaction")
		return fmt.Errorf("failed to get event from database")
	} else if existingReaction == nil {
		return errTargetNotFound
	}

	resp, err := sender.Client.SendReaction(&gmproto.SendReactionRequest{
		MessageID:    existingReaction.MessageID,
		ReactionData: gmproto.MakeReactionData(existingReaction.Reaction),
		Action:       gmproto.SendReactionRequest_REMOVE,
	})
	if err != nil {
		return fmt.Errorf("failed to send reaction removal: %w", err)
	} else if !resp.Success {
		return fmt.Errorf("got non-success response")
	}
	err = existingReaction.Delete(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to remove reaction from database after Matrix redaction")
	}
	return nil
}

func (portal *Portal) handleMatrixRedaction(sender *User, evt *event.Event) error {
	log := portal.zlog.With().
		Str("event_id", evt.ID.String()).
		Str("target_event_id", evt.Redacts.String()).
		Str("action", "handle matrix redaction").
		Logger()
	ctx := log.WithContext(context.TODO())
	log.Debug().Msg("Handling Matrix redaction")

	err := portal.handleMatrixMessageRedaction(ctx, sender, evt.Redacts)
	if err == errTargetNotFound {
		err = portal.handleMatrixReactionRedaction(ctx, sender, evt.Redacts)
	}
	return err
}

func (portal *Portal) Delete() {
	err := portal.Portal.Delete(context.TODO())
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to delete portal from database")
	}
	portal.bridge.portalsLock.Lock()
	delete(portal.bridge.portalsByKey, portal.Key)
	if len(portal.MXID) > 0 {
		delete(portal.bridge.portalsByMXID, portal.MXID)
	}
	portal.bridge.portalsLock.Unlock()
}

func (portal *Portal) GetMatrixUsers() ([]id.UserID, error) {
	members, err := portal.MainIntent().JoinedMembers(portal.MXID)
	if err != nil {
		return nil, fmt.Errorf("failed to get member list: %w", err)
	}
	var users []id.UserID
	for userID := range members.Joined {
		_, isPuppet := portal.bridge.ParsePuppetMXID(userID)
		if !isPuppet && userID != portal.bridge.Bot.UserID {
			users = append(users, userID)
		}
	}
	return users, nil
}

func (portal *Portal) CleanupIfEmpty() {
	users, err := portal.GetMatrixUsers()
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to get Matrix user list to determine if portal needs to be cleaned up")
		return
	}

	if len(users) == 0 {
		portal.zlog.Info().Msg("Room seems to be empty, cleaning up...")
		portal.Delete()
		portal.Cleanup(false)
	}
}

func (portal *Portal) Cleanup(puppetsOnly bool) {
	if len(portal.MXID) == 0 {
		return
	}
	intent := portal.MainIntent()
	if portal.bridge.SpecVersions.Supports(mautrix.BeeperFeatureRoomYeeting) {
		err := intent.BeeperDeleteRoom(portal.MXID)
		if err != nil && !errors.Is(err, mautrix.MNotFound) {
			portal.zlog.Err(err).Msg("Failed to delete room using hungryserv yeet endpoint")
		}
		return
	}
	members, err := intent.JoinedMembers(portal.MXID)
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to get portal members for cleanup")
		return
	}
	for member := range members.Joined {
		if member == intent.UserID {
			continue
		}
		puppet := portal.bridge.GetPuppetByMXID(member)
		if puppet != nil {
			_, err = puppet.DefaultIntent().LeaveRoom(portal.MXID)
			if err != nil {
				portal.zlog.Err(err).Msg("Failed to leave as puppet while cleaning up portal")
			}
		} else if !puppetsOnly {
			_, err = intent.KickUser(portal.MXID, &mautrix.ReqKickUser{UserID: member, Reason: "Deleting portal"})
			if err != nil {
				portal.zlog.Err(err).Msg("Failed to kick user while cleaning up portal")
			}
		}
	}
	_, err = intent.LeaveRoom(portal.MXID)
	if err != nil {
		portal.zlog.Err(err).Msg("Failed to leave with main intent while cleaning up portal")
	}
}
