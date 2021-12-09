package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"webtest/conf"
	"webtest/router"
	"webtest/spider"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	r := router.InitRouter()
	log.Info(conf.Config.Port)
	log.Info(conf.Config.Url)
	spider.GenSpider()
	err := r.Run(":" + conf.Config.Port)
	if err != nil {
		panic(err)
	}
}
