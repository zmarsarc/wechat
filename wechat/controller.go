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

	// if message
	MsgID int64 `xml:"MsgId"`

	// if common text message
	Content string `xml:"Content"`

	// if image message
	PicURL string `xml:"PicUrl"`

	// if voice message
	Format string `xml:"Format"`

	// if voice recognization enable
	Recognition string `xml:"Recognition"`

	// if any mulit media message
	MediaID string `xml:"MediaId"`

	// if video message
	ThumbMediaID string `xml:"ThumbMediaId"`

	// if link message
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`

	// if event
	Event string `xml:"Event"`

	// if scan QR code event
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`

	// if report location event
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
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
