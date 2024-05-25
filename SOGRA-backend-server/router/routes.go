package router

import (
	"SOGRA/account_controller"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func initRouterHandler(r *gin.Engine) {
	r.GET("/ping", pingHandler)

	r.POST("/register", account_controller.GetRegisterManager)
	r.POST("/login", account_controller.GetLoginManager)

	r.GET("/events/", _)
	r.GET("/events/{:id}/outline", _)
	r.GET("/events/{:id}", _)
	r.GET("/events/search", _)

	r.GET("/users/saves", _)
	r.POST("/users/saves", _)
	r.DELETE("/users/saves/{:id}", _)

	r.POST("/courses/create", _)
	r.POST("/courses/add")
	r.DELETE("/courses/{:id}")

	r.GET("/courses", _)
	r.GET("/courses/{:id}", _)

	r.GET("/users/profile-image", _)
	r.GET("/users/nickname", _)

}

func InitServer() {
	r := gin.Default()

	initRouterHandler(r)

	r.Run()
}
