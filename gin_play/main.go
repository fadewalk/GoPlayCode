package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	// 无参数
	r.GET("/",func(c *gin.Context){
		c.String(200,"Hello World")
	})

	// 解析路径参数

	r.GET("/user/:name",func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
		},
	)


	// 获取Query参数 
	r.GET("/users",func(ctx *gin.Context) {
		name := ctx.Query("name")
		role := ctx.DefaultQuery("role","teacher")
		ctx.String(http.StatusOK,"Hello %s, you are %s", name, role)
	})

	// 获取POST参数
	r.POST("/form",func(ctx *gin.Context) { 
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password","123456")
		ctx.JSON(http.StatusOK,
			gin.H{
				"username":username,
				"password":password,
			},
		)
	})


	// Query和POST混合参数 
	r.POST("/posts",func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page","0")

		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password","123456")

		ctx.JSON(http.StatusOK,
			gin.H{
				"id":id,
				"page":page,
				"username":username,
				"password":password,
			},
		)
	})

	// Map参数(字典参数)

	r.POST("/post", func(ctx *gin.Context) {
		ids := ctx.QueryMap("ids")
		names := ctx.PostFormMap("names")

		ctx.JSON(http.StatusOK,
			gin.H{
				"ids":ids,
				"names":names})
	})





	r.Run(":9999")
}