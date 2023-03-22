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

func GetAddressAndBalanceFromPrivateKey(c *gin.Context) {
  privateKey := c.Params.ByName("private-key")
  
  address, err := bitcoin_service.GetAddressFromPrivateKey(privateKey)
  if err != nil {
    c.JSON(500, gin.H{
      "error": err.Error(),
    })
    return
  }

  balance, err := bitcoin_service.GetBalanceFromAddress(address)

  if err != nil {
    c.JSON(500, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "private-key": privateKey,
    "address": address,
    "balance": balance,
  })
}