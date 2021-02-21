package wechat

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccess(t *testing.T) {
	router := gin.New()
	router.GET("/", Access)

	Convey("check access", t, func() {
		Convey("should return echostr if exists", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/?echostr=test", nil)

			router.ServeHTTP(w, req)

			res, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				panic(err)
			}

			So(string(res), ShouldEqual, "test")
		})
	})
}

type fakeSaver struct {
	msg BasicMessage
}

func (f *fakeSaver) Save(msg BasicMessage) error {
	f.msg = msg
	return nil
}

func TestMessage(t *testing.T) {
	Convey("check handle message from wechat", t, func() {
		var saver fakeSaver
		router := gin.New()
		router.POST("/", MessageHandler(&saver))

		Convey("common message", func() {

			w := httptest.NewRecorder()

			Convey("should return success if ok", func() {
				req, _ := http.NewRequest("POST", "/", strings.NewReader("<xml></xml>"))
				router.ServeHTTP(w, req)

				body, err := ioutil.ReadAll(w.Result().Body)
				if err != nil {
					panic(err)
				}

				So(string(body), ShouldEqual, "success")
			})

			Convey("should write message to storage", func() {
				const text = "<xml>" +
					"<ToUserName><![CDATA[toUser]]></ToUserName>" +
					"<FromUserName><![CDATA[fromUser]]></FromUserName>" +
					"<CreateTime>1348831860</CreateTime>" +
					"<MsgType><![CDATA[text]]></MsgType>" +
					"<Content><![CDATA[this is a test]]></Content>" +
					"<MsgId>1234567890123456</MsgId>" +
					"</xml>"

				req, _ := http.NewRequest("POST", "/", strings.NewReader(text))
				router.ServeHTTP(w, req)
				So(saver.msg.Content, ShouldEqual, "this is a test")
			})

			Convey("should write image message", func() {
				const text = "<xml>" +
					"<ToUserName><![CDATA[toUser]]></ToUserName>" +
					"<FromUserName><![CDATA[fromUser]]></FromUserName>" +
					"<CreateTime>1348831860</CreateTime>" +
					"<MsgType><![CDATA[image]]></MsgType>" +
					"<PicUrl><![CDATA[this is a url]]></PicUrl>" +
					"<MediaId><![CDATA[media_id]]></MediaId>" +
					"<MsgId>1234567890123456</MsgId>" +
					"</xml>"

				req, _ := http.NewRequest("POST", "/", strings.NewReader(text))
				router.ServeHTTP(w, req)
				Convey("type should be image", func() {
					So(saver.msg.MsgType, ShouldEqual, "image")
				})
				Convey("content should be media_id", func() {
					So(saver.msg.MediaID, ShouldEqual, "media_id")
				})
				Convey("pic url should set", func() {
					So(saver.msg.PicURL, ShouldEqual, "this is a url")
				})
			})

			Convey("should write voice message", func() {
				const text = "<xml>" +
					"<ToUserName><![CDATA[toUser]]></ToUserName>" +
					"<FromUserName><![CDATA[fromUser]]></FromUserName>" +
					"<CreateTime>1357290913</CreateTime>" +
					"<MsgType><![CDATA[voice]]></MsgType>" +
					"<MediaId><![CDATA[media_id]]></MediaId>" +
					"<Format><![CDATA[Format]]></Format>" +
					"<MsgId>1234567890123456</MsgId>" +
					"</xml>"

				req, _ := http.NewRequest("POST", "/", strings.NewReader(text))
				router.ServeHTTP(w, req)

				Convey("message type should be voice", func() {
					So(saver.msg.MsgType, ShouldEqual, "voice")
				})
				Convey("format should be format", func() {
					So(saver.msg.Format, ShouldEqual, "Format")
				})
			})

			Convey("should support video message", func() {
				const text = "<xml>" +
					"<ToUserName><![CDATA[toUser]]></ToUserName>" +
					"<FromUserName><![CDATA[fromUser]]></FromUserName>" +
					"<CreateTime>1357290913</CreateTime>" +
					"<MsgType><![CDATA[video]]></MsgType>" +
					"<MediaId><![CDATA[media_id]]></MediaId>" +
					"<ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>" +
					"<MsgId>1234567890123456</MsgId>" +
					"</xml>"

				req, _ := http.NewRequest("POST", "/", strings.NewReader(text))
				router.ServeHTTP(w, req)

				Convey("type should be video", func() {
					So(saver.msg.MsgType, ShouldEqual, "video")
				})
				Convey("thumb media id should right", func() {
					So(saver.msg.ThumbMediaID, ShouldEqual, "thumb_media_id")
				})
			})
		})
	})
}
