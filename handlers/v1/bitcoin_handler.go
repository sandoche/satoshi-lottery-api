package bitcoin_handler

import (
	"github.com/gin-gonic/gin"

	bitcoin_service "satoshi-lottery/services"
)


func GetRandomBitCoinPrivateKey(c *gin.Context) {
  privateKey :=  bitcoin_service.GenerateRandomBitCoinPrivateKey()

  c.JSON(200, gin.H{
    "private-key": privateKey,
  })
}