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
	hostname   string
	config     ClientConfig
}

// ClientConfig use to config wechat access token
type ClientConfig struct {
	AppID  string
	Secret string
}

// NewClient is constructor of Client
func NewClient(conf *ClientConfig) *Client {
	return &Client{
		httpClient: &http.Client{},
		hostname:   "https://api.weixin.qq.com",
		config:     *conf,
	}
}

type accessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AccessToken get wechat access token
func (c *Client) AccessToken() (string, error) {
	query := url.Values{}
	query.Add("grant_type", "client_credential")
	query.Add("appid", c.config.AppID)
	query.Add("secret", c.config.Secret)

	resp, err := c.httpClient.Get(c.hostname + "/token?" + query.Encode())
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
