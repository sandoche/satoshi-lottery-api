package main

import (
	"github.com/gin-gonic/gin"

	"satoshi-lottery/handlers/v1"
)

func main() {
  r := gin.Default()

  r.GET("/v1/bitcoin-key", bitcoin_handler.GenerateRandomBitCoinPrivateKey)
  
  r.Run()
}
