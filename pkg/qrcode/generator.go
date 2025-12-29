package qrcode

import (
	"encoding/base64"
	"fmt"

	"github.com/skip2/go-qrcode"
)

// GeneratePNG returns the PNG bytes of the QR code for the given content.
func GeneratePNG(content string, size int) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}
	return png, nil
}

// GenerateBase64 returns the Base64 encoded string of the PNG QR code.
func GenerateBase64(content string, size int) (string, error) {
	png, err := GeneratePNG(content, size)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(png), nil
}
