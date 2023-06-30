package textgapi

import (
	"encoding/base64"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
	"go.mau.fi/mautrix-gmessages/libgm/util"
)

func (p *Pairer) GenerateQRCodeData() (string, error) {
	urlData := &binary.UrlData{
		PairingKey:      p.pairingKey,
		AES_CTR_KEY_256: p.client.cryptor.AES_CTR_KEY_256,
		SHA_256_KEY:     p.client.cryptor.SHA_256_KEY,
	}
	encodedUrlData, err := binary.EncodeProtoMessage(urlData)
	if err != nil {
		return "", err
	}
	cData := base64.StdEncoding.EncodeToString(encodedUrlData)
	return util.QR_CODE_URL + cData, nil
}
