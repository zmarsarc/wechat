package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/zmarsarc/wechat/wechat"
)

func main() {
	config, err := ini.Load("conf/app.conf")
	if err != nil {
		panic(fmt.Errorf("load config error: %s", err.Error()))
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	wechatRouter := r.Group("wechat")
	wechatRouter.Use(wechat.Auth(config.Section("wechat").Key("token").String()))
	{
	}

	r.Run()
}
