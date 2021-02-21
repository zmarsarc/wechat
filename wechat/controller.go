package wechat

import "github.com/gin-gonic/gin"

// Access will let wechat server access your server
func Access(c *gin.Context) {
	c.String(200, c.Query("echostr"))
}

// Message handle message which send from wechat server
func Message(c *gin.Context) {
	c.String(200, "success")
}
