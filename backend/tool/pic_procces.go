package tools

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func ParseBase64Image(data string) (*[]byte, error) {
	// Убираем префикс "data:image/png;base64,"
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid Base64 image format")
	}

	// Декодируем Base64
	imageData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64: %w", err)
	}

	return &imageData, nil
}
