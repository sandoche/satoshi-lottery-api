package bitcoin_handler

import (
	"github.com/gin-gonic/gin"
)


func GenerateRandomBitCoinPrivateKey(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}