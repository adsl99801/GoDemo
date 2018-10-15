package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

//TokenAuth o
func TokenAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Println("TokenAuth")
		token := c.Request.Header.Get("Token")
		if len(token) <= 0 {
			log.Println("len(token) <= 0")
			c.Abort()
			return
		}
		c.Set("userID", token)
		c.Next()

	}

}
