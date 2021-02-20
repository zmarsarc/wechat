package wechat

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
