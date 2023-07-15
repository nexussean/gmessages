package libgm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.mau.fi/mautrix-gmessages/libgm/events"
	"go.mau.fi/mautrix-gmessages/libgm/pblite"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

type RPC struct {
	client       *Client
	http         *http.Client
	conn         io.ReadCloser
	rpcSessionId string
	listenID     int

	skipCount int

	recentUpdates    [8][32]byte
	recentUpdatesPtr int
}

func (r *RPC) ListenReceiveMessages(payload []byte) {
	r.listenID++
	listenID := r.listenID
	errored := true
	for r.listenID == listenID {
		if r.client.authData.DevicePair != nil && r.client.authData.AuthenticatedAt.Add(20*time.Hour).Before(time.Now()) {
			r.client.Logger.Debug().Msg("Refreshing auth token before starting new long-polling request")
			err := r.client.refreshAuthToken()
			if err != nil {
				r.client.Logger.Err(err).Msg("Error refreshing auth token")
				r.client.triggerEvent(&events.ListenFatalError{Error: fmt.Errorf("failed to refresh auth token: %w", err)})
				return
			}
		}
		r.client.Logger.Debug().Msg("Starting new long-polling request")
		req, err := http.NewRequest("POST", util.RECEIVE_MESSAGES, bytes.NewReader(payload))
		if err != nil {
			panic(fmt.Errorf("Error creating request: %v", err))
		}
		util.BuildRelayHeaders(req, "application/json+protobuf", "*/*")
		resp, reqErr := r.http.Do(req)
		//r.client.Logger.Info().Any("bodyLength", len(payload)).Any("url", util.RECEIVE_MESSAGES).Any("headers", resp.Request.Header).Msg("RPC Request Headers")
		if reqErr != nil {
			r.client.triggerEvent(&events.ListenTemporaryError{Error: reqErr})
			errored = true
			r.client.Logger.Err(err).Msg("Error making listen request, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
			continue
		}
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			r.client.Logger.Error().Int("status_code", resp.StatusCode).Msg("Error making listen request")
			r.client.triggerEvent(&events.ListenFatalError{Error: events.HTTPError{Action: "polling", Resp: resp}})
			return
		} else if resp.StatusCode >= 500 {
			r.client.triggerEvent(&events.ListenTemporaryError{Error: events.HTTPError{Action: "polling", Resp: resp}})
			errored = true
			r.client.Logger.Debug().Int("statusCode", resp.StatusCode).Msg("5xx error in long polling, retrying in 5 seconds")
			time.Sleep(5 * time.Second)
			continue
		}
		if errored {
			errored = false
			r.client.triggerEvent(&events.ListenRecovered{})
		}
		r.client.Logger.Debug().Int("statusCode", resp.StatusCode).Msg("Long polling opened")
		r.conn = resp.Body
		if r.client.authData.DevicePair != nil {
			go func() {
				err := r.client.Session.NotifyDittoActivity()
				if err != nil {
					r.client.Logger.Err(err).Msg("Error notifying ditto activity")
				}
			}()
		}
		r.startReadingData(resp.Body)
	}
}

/*
	The start of a message always begins with byte 44 (",")
	If the message is parsable (after , has been removed) as an array of interfaces:
	func (r *RPC) tryUnmarshalJSON(jsonData []byte, msgArr *[]interface{}) error {
		err := json.Unmarshal(jsonData, &msgArr)
		return err
	}
	then the message is complete and it should continue to the HandleRPCMsg function and it should also reset the buffer so that the next message can be received properly.

	if it's not parsable, it should just append the received data to the buf and attempt to parse it until it's parsable. Because that would indicate that the full msg has been received
*/

func (r *RPC) startReadingData(rc io.ReadCloser) {
	defer rc.Close()
	reader := bufio.NewReader(rc)
	buf := make([]byte, 2621440)
	var accumulatedData []byte
	n, err := reader.Read(buf[:2])
	if err != nil {
		r.client.Logger.Err(err).Msg("Error reading opening bytes")
		return
	} else if n != 2 || string(buf[:2]) != "[[" {
		r.client.Logger.Err(err).Msg("Opening is not [[")
		return
	}
	for {
		n, err = reader.Read(buf)
		if err != nil {
			if errors.Is(err, os.ErrClosed) {
				r.client.Logger.Err(err).Msg("Closed body from server")
				r.conn = nil
				return
			}
			r.client.Logger.Err(err).Msg("Stopped reading data from server")
			return
		}
		chunk := buf[:n]
		if len(accumulatedData) == 0 {
			if len(chunk) == 2 && string(chunk) == "]]" {
				r.client.Logger.Debug().Msg("Got stream end marker")
			}
			chunk = bytes.TrimPrefix(chunk, []byte{','})
		}
		accumulatedData = append(accumulatedData, chunk...)
		if !json.Valid(accumulatedData) {
			r.client.Logger.Debug().Str("data", string(chunk)).Msg("Invalid JSON")
			continue
		}
		var msgArr []any
		err = json.Unmarshal(accumulatedData, &msgArr)
		if err != nil {
			r.client.Logger.Err(err).Msg("Error unmarshalling json")
			continue
		}
		accumulatedData = accumulatedData[:0]
		msg := &binary.InternalMessage{}
		err = pblite.Deserialize(msgArr, msg.ProtoReflect())
		if err != nil {
			r.client.Logger.Err(err).Msg("Error deserializing pblite message")
			continue
		}
		switch {
		case msg.GetData() != nil:
			r.HandleRPCMsg(msg)
		case msg.GetAck() != nil:
			r.client.Logger.Debug().Int32("count", msg.GetAck().GetCount()).Msg("Got startup ack count message")
			r.skipCount = int(msg.GetAck().GetCount())
		case msg.GetStartRead() != nil:
			r.client.Logger.Trace().Msg("Got startRead message")
		case msg.GetHeartbeat() != nil:
			r.client.Logger.Trace().Msg("Got heartbeat message")
		default:
			r.client.Logger.Warn().Interface("data", msgArr).Msg("Got unknown message")
		}
	}
}

func (r *RPC) CloseConnection() {
	if r.conn != nil {
		r.client.Logger.Debug().Msg("Attempting to connection...")
		r.conn.Close()
		r.conn = nil
	}
}

func (r *RPC) sendMessageRequest(url string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	util.BuildRelayHeaders(req, "application/json+protobuf", "*/*")
	resp, reqErr := r.client.http.Do(req)
	//r.client.Logger.Info().Any("bodyLength", len(payload)).Any("url", url).Any("headers", resp.Request.Header).Msg("RPC Request Headers")
	if reqErr != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return resp, reqErr
}
