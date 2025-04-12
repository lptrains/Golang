package user

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func RouterRegistry(e *gin.Engine) {
	r := e.Group("/user")
	r.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "XXX",
		})
	})
	r.GET("/login", func(c *gin.Context) {

		c.HTML(http.StatusOK, "user/login.html",gin.H{
			
		})
	})
	//login post
	r.POST("/doLogin", func(c *gin.Context) {
		//获取form表单数据
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Printf("name: %v, pwd: %v", username, password)
		c.HTML(http.StatusOK, "/user/info.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
}
