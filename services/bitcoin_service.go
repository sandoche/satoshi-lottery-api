package bitoin_service

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomBitCoinPrivateKey() string {
	privateKey := make([]byte, 32)
	rand.Read(privateKey)
	return hex.EncodeToString(privateKey)
}