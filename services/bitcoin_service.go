package bitoin_service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/btcsuite/btcd/chaincfg"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil"
)

type blockcypherBalanceData struct {
    Address           string `json:"address"`
    TotalReceived     int    `json:"total_received"`
    TotalSent         int    `json:"total_sent"`
    Balance           int    `json:"balance"`
    UnconfirmedBalance int   `json:"unconfirmed_balance"`
    FinalBalance      int    `json:"final_balance"`
    NTx               int    `json:"n_tx"`
    UnconfirmedNTx    int    `json:"unconfirmed_n_tx"`
    FinalNTx          int    `json:"final_n_tx"`
}

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

func GetBalanceFromAddress(address string) (int, error) {
	endpoint := "https://api.blockcypher.com/v1/btc/main/addrs/" + address + "/balance"

	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data blockcypherBalanceData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.FinalBalance, nil
}