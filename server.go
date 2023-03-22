package main

import (
	"github.com/gin-gonic/gin"

	bitcoin_handler "satoshi-lottery/handlers/v1"
)

func main() {
  r := gin.Default()

  r.GET("/v1/bitcoin", bitcoin_handler.GetRandomBitCoinPrivateKey)

  r.GET("/v1/bitcoin/:private-key", bitcoin_handler.GetAddressAndBalanceFromPrivateKey)
  
  r.Run()
}
