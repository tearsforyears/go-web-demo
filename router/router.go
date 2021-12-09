package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"webtest/db"
	"webtest/middleware"
	. "webtest/model"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &gin.H{"status": "200", "msg": "pong"})
	})

	main := r.Group("/main")
	{
		main.GET("/:id", middleware.Baned(), func(ctx *gin.Context) {
			demo := Demo{}
			if err := db.Mysql.Find(&demo, ctx.Param("id")).Error; err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusOK, demo)
		})
		main.GET("/", func(ctx *gin.Context) {
			demos := []Demo{}
			if err := db.Mysql.Raw("select * from springdemo").Select(db.Mysql.Rows()).Scan(&demos).Error; err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusOK, demos)
		})
	}

	test := r.Group("/t")
	{
		test.GET("/p", func(ctx *gin.Context) {
			type Param struct {
				Id   int    `json:"id" form:"id"`
				Name string `json:"name" form:"name"`
			}
			param := Param{}
			if err := ctx.ShouldBindQuery(&param); err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusOK, param)
		})

		test.PUT("/p", func(ctx *gin.Context) {
			type Param struct {
				Id   int    `json:"id" form:"id"`
				Name string `json:"name" form:"name"`
			}
			param := Param{}
			if err := ctx.ShouldBindJSON(&param); err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusOK, param)
		})

		test.POST("/p", func(ctx *gin.Context) {
			type Param struct {
				Id   int    `json:"id" form:"id"`
				Name string `json:"name" form:"name"`
			}
			param := Param{}
			if err := ctx.ShouldBindWith(&param, binding.Form); err != nil {
				panic(err)
			}
			ctx.JSON(http.StatusOK, param)
		})

		test.POST("/upload", func(c *gin.Context) {
			file, _ := c.FormFile("file")
			c.SaveUploadedFile(file, "./data/"+file.Filename)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})
	}
	return r
}
