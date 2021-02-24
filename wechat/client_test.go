package wechat

import (
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClientAccessToken(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/token", func(resp http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		if query.Get("grant_type") != "client_credential" {
			resp.WriteHeader(400)
			return
		}
		if query.Get("appid") != "appid" {
			resp.WriteHeader(400)
			return
		}
		if query.Get("secret") != "secret" {
			resp.WriteHeader(400)
			return
		}

		token := accessTokenResp{
			AccessToken: "token",
			ExpiresIn:   7200,
		}
		data, _ := json.Marshal(token)
		if _, err := resp.Write(data); err != nil {
			panic(err)
		}
	})
	server := http.Server{}
	server.Handler = handler
	server.Addr = "127.0.0.1:8080"

	go server.ListenAndServe()
	defer server.Close()

	Convey("check get access token from wechat", t, func() {
		config := ClientConfig{
			AppID:  "appid",
			Secret: "secret",
		}
		client := NewClient(&config)
		client.hostname = "http://localhost:8080"

		Convey("get access token", func() {
			token, _ := client.AccessToken()
			So(token, ShouldEqual, "token")
		})
	})
}
