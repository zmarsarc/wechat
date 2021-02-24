package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client provided wechat access
type Client struct {
	httpClient *http.Client
}

// ClientConfig use to config wechat access token
type ClientConfig struct{}

type accessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AccessToken get wechat access token
func (c *Client) AccessToken() (string, error) {
	query := url.Values{}
	query.Add("grant_type", "client_credential")
	query.Add("appid", "")
	query.Add("secret", "")

	resp, err := c.httpClient.Get("https://api.weixin.qq.com/cgi-bin/token?" + query.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var token accessTokenResp
	if err := json.Unmarshal(data, &token); err != nil {
		panic(err)
	}

	return token.AccessToken, nil
}
