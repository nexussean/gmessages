package libgm

import (
	"fmt"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
)

type Messages struct {
	client *Client
}

func (m *Messages) React(payload *binary.SendReactionPayload) (*binary.SendReactionResponse, error) {
	actionType := binary.ActionType_SEND_REACTION

	sentRequestId, sendErr := m.client.sessionHandler.completeSendMessage(actionType, true, payload)
	if sendErr != nil {
		return nil, sendErr
	}

	response, err := m.client.sessionHandler.WaitForResponse(sentRequestId, actionType)
	if err != nil {
		return nil, err
	}

	res, ok := response.Data.Decrypted.(*binary.SendReactionResponse)
	if !ok {
		return nil, fmt.Errorf("failed to assert response into SendReactionResponse")
	}

	m.client.Logger.Debug().Any("res", res).Msg("sent reaction!")
	return res, nil
}

func (m *Messages) Delete(messageId string) (*binary.DeleteMessageResponse, error) {
	payload := &binary.DeleteMessagePayload{MessageID: messageId}
	actionType := binary.ActionType_DELETE_MESSAGE

	sentRequestId, sendErr := m.client.sessionHandler.completeSendMessage(actionType, true, payload)
	if sendErr != nil {
		return nil, sendErr
	}

	response, err := m.client.sessionHandler.WaitForResponse(sentRequestId, actionType)
	if err != nil {
		return nil, err
	}

	res, ok := response.Data.Decrypted.(*binary.DeleteMessageResponse)
	if !ok {
		return nil, fmt.Errorf("failed to assert response into DeleteMessageResponse")
	}

	m.client.Logger.Debug().Any("res", res).Msg("deleted message!")
	return res, nil
}
