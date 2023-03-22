package bitoin_service

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil"
)

func GenerateRandomBitCoinPrivateKey() string {
	privateKey := make([]byte, 32)
	rand.Read(privateKey)
	return hex.EncodeToString(privateKey)
}

func GetAddressFromPrivateKey(privateKeyString string) string {
	decodedPrivateKey, _ := hex.DecodeString(privateKeyString)
	privateKey, _ := btcec.PrivKeyFromBytes(decodedPrivateKey)
	publicKey := privateKey.PubKey().SerializeCompressed()
	address, _ := btcutil.NewAddressPubKey(publicKey, &chaincfg.MainNetParams)
	return address.EncodeAddress()
}