package wechat

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

// fake wechat request query
const testQuery = "signature=7c7f567747cd0ed5138720ea35b4e69f910c727a&echostr=5706686390652010495&timestamp=1612061747&nonce=1274437272"
const testToken = "zmarsarc"

func TestAuth(t *testing.T) {
	router := gin.New()
	router.Use(Auth(testToken))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
		})
	})

	Convey("check server request check", t, func() {
		Convey("should not pass if no signature", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)

			router.ServeHTTP(w, req)

			So(w.Result().StatusCode, ShouldNotEqual, 200)
		})

		Convey("should pass if have signature", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/?"+testQuery, nil)

			router.ServeHTTP(w, req)
			So(w.Result().StatusCode, ShouldEqual, 200)
		})
	})
}
