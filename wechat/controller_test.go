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

func TestMessage(t *testing.T) {
	router := gin.New()
	router.POST("/", Message)

	Convey("check handle message from wechat", t, func() {
		Convey("common message", func() {
			const text = "<xml>" +
				"<ToUserName><![CDATA[toUser]]></ToUserName>" +
				"<FromUserName><![CDATA[fromUser]]></FromUserName>" +
				"<CreateTime>1348831860</CreateTime>" +
				"<MsgType><![CDATA[text]]></MsgType>" +
				"<Content><![CDATA[this is a test]]></Content>" +
				"<MsgId>1234567890123456</MsgId>" +
				"</xml>"

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/", strings.NewReader(text))

			router.ServeHTTP(w, req)

			body, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				panic(err)
			}

			So(string(body), ShouldEqual, "success")
		})
	})
}
