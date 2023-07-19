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
	"fmt"

	"github.com/skip2/go-qrcode"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/bridge/commands"
	"maunium.net/go/mautrix/bridge/status"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

type WrappedCommandEvent struct {
	*commands.Event
	Bridge *GMBridge
	User   *User
	Portal *Portal
}

func (br *GMBridge) RegisterCommands() {
	proc := br.CommandProcessor.(*commands.Processor)
	proc.AddHandlers(
		cmdLogin,
		cmdDeleteSession,
		cmdLogout,
		cmdReconnect,
		cmdDisconnect,
		cmdPing,
		cmdDeletePortal,
		cmdDeleteAllPortals,
	)
}

func wrapCommand(handler func(*WrappedCommandEvent)) func(*commands.Event) {
	return func(ce *commands.Event) {
		user := ce.User.(*User)
		var portal *Portal
		if ce.Portal != nil {
			portal = ce.Portal.(*Portal)
		}
		br := ce.Bridge.Child.(*GMBridge)
		handler(&WrappedCommandEvent{ce, br, user, portal})
	}
}

var (
	HelpSectionConnectionManagement = commands.HelpSection{Name: "Connection management", Order: 11}
	HelpSectionPortalManagement     = commands.HelpSection{Name: "Portal management", Order: 20}
)

var cmdLogin = &commands.FullHandler{
	Func: wrapCommand(fnLogin),
	Name: "login",
	Help: commands.HelpMeta{
		Section:     commands.HelpSectionAuth,
		Description: "Link the bridge to Google Messages on your Android phone as a web client.",
	},
}

func fnLogin(ce *WrappedCommandEvent) {
	if ce.User.Session != nil {
		if ce.User.IsConnected() {
			ce.Reply("You're already logged in")
		} else {
			ce.Reply("You're already logged in. Perhaps you wanted to `reconnect`?")
		}
		return
	} else if ce.User.pairSuccessChan != nil {
		ce.Reply("You already have a login in progress")
		return
	}

	ch, err := ce.User.Login(6)
	if err != nil {
		ce.ZLog.Err(err).Msg("Failed to start login")
		ce.Reply("Failed to start login: %v", err)
		return
	}
	var prevEvent id.EventID
	for item := range ch {
		switch {
		case item.qr != "":
			ce.ZLog.Debug().Msg("Got code in QR channel")
			prevEvent = ce.User.sendQR(ce, item.qr, prevEvent)
		case item.err != nil:
			ce.ZLog.Err(err).Msg("Error in QR channel")
			prevEvent = ce.User.sendQREdit(ce, &event.MessageEventContent{
				MsgType: event.MsgNotice,
				Body:    fmt.Sprintf("Failed to log in: %v", err),
			}, prevEvent)
		case item.success:
			ce.ZLog.Debug().Msg("Got pair success in QR channel")
			prevEvent = ce.User.sendQREdit(ce, &event.MessageEventContent{
				MsgType: event.MsgNotice,
				Body:    "Successfully logged in",
			}, prevEvent)
		default:
			ce.ZLog.Error().Any("item_data", item).Msg("Unknown item in QR channel")
		}
	}
	ce.ZLog.Trace().Msg("Login command finished")
}

func (user *User) sendQREdit(ce *WrappedCommandEvent, content *event.MessageEventContent, prevEvent id.EventID) id.EventID {
	if len(prevEvent) != 0 {
		content.SetEdit(prevEvent)
	}
	resp, err := ce.Bot.SendMessageEvent(ce.RoomID, event.EventMessage, &content)
	if err != nil {
		ce.ZLog.Err(err).Msg("Failed to send edited QR code")
	} else if len(prevEvent) == 0 {
		prevEvent = resp.EventID
	}
	return prevEvent
}

func (user *User) sendQR(ce *WrappedCommandEvent, code string, prevEvent id.EventID) id.EventID {
	var content event.MessageEventContent
	url, err := user.uploadQR(code)
	if err != nil {
		ce.ZLog.Err(err).Msg("Failed to upload QR code")
		content = event.MessageEventContent{
			MsgType: event.MsgNotice,
			Body:    fmt.Sprintf("Failed to upload QR code: %v", err),
		}
	} else {
		content = event.MessageEventContent{
			MsgType: event.MsgImage,
			Body:    code,
			URL:     url.CUString(),
		}
	}
	return user.sendQREdit(ce, &content, prevEvent)
}

func (user *User) uploadQR(code string) (id.ContentURI, error) {
	qrCode, err := qrcode.Encode(code, qrcode.Low, 256)
	if err != nil {
		return id.ContentURI{}, err
	}
	resp, err := user.bridge.Bot.UploadBytes(qrCode, "image/png")
	if err != nil {
		return id.ContentURI{}, err
	}
	return resp.ContentURI, nil
}

var cmdLogout = &commands.FullHandler{
	Func: wrapCommand(fnLogout),
	Name: "logout",
	Help: commands.HelpMeta{
		Section:     commands.HelpSectionAuth,
		Description: "Unpair the bridge and delete session information.",
	},
}

func fnLogout(ce *WrappedCommandEvent) {
	if ce.User.Session == nil && ce.User.Client == nil {
		ce.Reply("You're not logged in")
		return
	}
	logoutOK := ce.User.Logout(status.BridgeState{StateEvent: status.StateLoggedOut}, true)
	if logoutOK {
		ce.Reply("Successfully logged out")
	} else {
		ce.Reply("Session information removed, but unpairing may not have succeeded")
	}
}

var cmdDeleteSession = &commands.FullHandler{
	Func: wrapCommand(fnDeleteSession),
	Name: "delete-session",
	Help: commands.HelpMeta{
		Section:     commands.HelpSectionAuth,
		Description: "Delete session information and disconnect from Google Messages without sending a logout request.",
	},
}

func fnDeleteSession(ce *WrappedCommandEvent) {
	if ce.User.Session == nil && ce.User.Client == nil {
		ce.Reply("Nothing to purge: no session information stored and no active connection.")
		return
	}
	ce.User.Logout(status.BridgeState{StateEvent: status.StateLoggedOut}, false)
	ce.Reply("Session information purged")
}

var cmdReconnect = &commands.FullHandler{
	Func: wrapCommand(fnReconnect),
	Name: "reconnect",
	Help: commands.HelpMeta{
		Section:     HelpSectionConnectionManagement,
		Description: "Reconnect to Google Messages.",
	},
}

func fnReconnect(ce *WrappedCommandEvent) {
	if ce.User.Client == nil {
		if ce.User.Session == nil {
			ce.Reply("You're not logged into Google Messages. Please log in first.")
		} else {
			ce.User.Connect()
			ce.Reply("Started connecting to Google Messages")
		}
	} else {
		ce.User.DeleteConnection()
		ce.User.BridgeState.Send(status.BridgeState{StateEvent: status.StateTransientDisconnect, Error: GMNotConnected})
		ce.User.Connect()
		ce.Reply("Restarted connection to Google Messages")
	}
}

var cmdDisconnect = &commands.FullHandler{
	Func: wrapCommand(fnDisconnect),
	Name: "disconnect",
	Help: commands.HelpMeta{
		Section:     HelpSectionConnectionManagement,
		Description: "Disconnect from Google Messages (without logging out).",
	},
}

func fnDisconnect(ce *WrappedCommandEvent) {
	if ce.User.Client == nil {
		ce.Reply("You don't have a Google Messages connection.")
		return
	}
	ce.User.DeleteConnection()
	ce.Reply("Successfully disconnected. Use the `reconnect` command to reconnect.")
	ce.User.BridgeState.Send(status.BridgeState{StateEvent: status.StateBadCredentials, Error: GMNotConnected})
}

var cmdPing = &commands.FullHandler{
	Func: wrapCommand(fnPing),
	Name: "ping",
	Help: commands.HelpMeta{
		Section:     HelpSectionConnectionManagement,
		Description: "Check your connection to Google Messages.",
	},
}

func fnPing(ce *WrappedCommandEvent) {
	if ce.User.Session == nil {
		if ce.User.Client != nil {
			ce.Reply("Connected to Google Messages, but not logged in.")
		} else {
			ce.Reply("You're not logged into Google Messages.")
		}
	} else if ce.User.Client == nil || !ce.User.Client.IsConnected() {
		ce.Reply("You're logged in as %s, but you don't have a Google Messages connection.", ce.User.PhoneID)
	} else {
		ce.Reply("Logged in as %s, connection to Google Messages may be OK", ce.User.PhoneID)
	}
}

func canDeletePortal(portal *Portal, userID id.UserID) bool {
	if len(portal.MXID) == 0 {
		return false
	}

	members, err := portal.MainIntent().JoinedMembers(portal.MXID)
	if err != nil {
		portal.log.Errorfln("Failed to get joined members to check if portal can be deleted by %s: %v", userID, err)
		return false
	}
	for otherUser := range members.Joined {
		_, isPuppet := portal.bridge.ParsePuppetMXID(otherUser)
		if isPuppet || otherUser == portal.bridge.Bot.UserID || otherUser == userID {
			continue
		}
		user := portal.bridge.GetUserByMXID(otherUser)
		if user != nil && user.Session != nil {
			return false
		}
	}
	return true
}

var cmdDeletePortal = &commands.FullHandler{
	Func: wrapCommand(fnDeletePortal),
	Name: "delete-portal",
	Help: commands.HelpMeta{
		Section:     HelpSectionPortalManagement,
		Description: "Delete the current portal. If the portal is used by other people, this is limited to bridge admins.",
	},
	RequiresPortal: true,
}

func fnDeletePortal(ce *WrappedCommandEvent) {
	if !ce.User.Admin && !canDeletePortal(ce.Portal, ce.User.MXID) {
		ce.Reply("Only bridge admins can delete portals with other Matrix users")
		return
	}

	ce.Portal.log.Infoln(ce.User.MXID, "requested deletion of portal.")
	ce.Portal.Delete()
	ce.Portal.Cleanup(false)
}

var cmdDeleteAllPortals = &commands.FullHandler{
	Func: wrapCommand(fnDeleteAllPortals),
	Name: "delete-all-portals",
	Help: commands.HelpMeta{
		Section:     HelpSectionPortalManagement,
		Description: "Delete all portals.",
	},
}

func fnDeleteAllPortals(ce *WrappedCommandEvent) {
	portals := ce.Bridge.GetAllPortalsForUser(ce.User.RowID)
	if len(portals) == 0 {
		ce.Reply("Didn't find any portals to delete")
		return
	}

	leave := func(portal *Portal) {
		if len(portal.MXID) > 0 {
			_, _ = portal.MainIntent().KickUser(portal.MXID, &mautrix.ReqKickUser{
				Reason: "Deleting portal",
				UserID: ce.User.MXID,
			})
		}
	}
	intent := ce.User.DoublePuppetIntent
	if intent != nil {
		leave = func(portal *Portal) {
			if len(portal.MXID) > 0 {
				_, _ = intent.LeaveRoom(portal.MXID)
				_, _ = intent.ForgetRoom(portal.MXID)
			}
		}
	}
	roomYeeting := ce.Bridge.SpecVersions.Supports(mautrix.BeeperFeatureRoomYeeting)
	if roomYeeting {
		leave = func(portal *Portal) {
			portal.Cleanup(false)
		}
	}
	ce.Reply("Found %d portals, deleting...", len(portals))
	for _, portal := range portals {
		portal.Delete()
		leave(portal)
	}
	if !roomYeeting {
		ce.Reply("Finished deleting portal info. Now cleaning up rooms in background.")
		go func() {
			for _, portal := range portals {
				portal.Cleanup(false)
			}
			ce.Reply("Finished background cleanup of deleted portal rooms.")
		}()
	} else {
		ce.Reply("Finished deleting portals.")
	}
}
