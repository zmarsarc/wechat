package wechat

import (
	"github.com/gin-gonic/gin"
)

// Access will let wechat server access your server
func Access(c *gin.Context) {
	c.String(200, c.Query("echostr"))
}

// BasicMessage is wechat message template
type BasicMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`

	// if common text message
	Content string `xml:"Content"`

	// if image message
	PicURL  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`

	MsgID string `xml:"MsgId"`
}

// MessageSaver define those who can save wechat message should implem
type MessageSaver interface {
	Save(msg BasicMessage) error
}

// MessageHandler handle message which send from wechat server
func MessageHandler(saver MessageSaver) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg BasicMessage
		if err := c.ShouldBindXML(&msg); err != nil {
			panic(err)
		}
		if err := saver.Save(msg); err != nil {
			panic(err)
		}
		c.String(200, "success")
	}
}
