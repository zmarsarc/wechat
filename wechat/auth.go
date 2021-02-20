package wechat

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth return a middleware which will check if a request is from wechat server
func Auth(token string) gin.HandlerFunc {
	if token == "" {
		panic(errors.New("wechat token must be specified"))
	}

	return func(c *gin.Context) {
		signature := c.Query("signature")
		timestamp := c.Query("timestamp")
		nonce := c.Query("nonce")

		arr := []string{nonce, timestamp, token}
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})

		checkSum := sha1.Sum([]byte(strings.Join(arr, "")))
		if strings.ToLower(hex.EncodeToString(checkSum[:])) == strings.ToLower(signature) {
			c.Next()
			return
		}

		c.AbortWithStatus(400)
	}
}
