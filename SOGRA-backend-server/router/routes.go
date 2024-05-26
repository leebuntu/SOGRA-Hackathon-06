package router

import (
	"SOGRA/event_controller"
	"SOGRA/user_controller"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func initRouterHandler(r *gin.Engine) {
	r.GET("/ping", pingHandler)

	r.POST("/register", user_controller.GetRegisterManager)
	r.POST("/login", user_controller.GetLoginManager)

	r.GET("/events/", event_controller.GetTrendEvent)                  // get trend event
	r.GET("/events/:id/outline", event_controller.GetEventOutlineInfo) // get specific event outline
	r.GET("/events/:id", event_controller.GetEventFullInfo)            // get specific event
	r.GET("/events/search", event_controller.GetEventSearchResult)     // search event

	r.GET("/saves/:id", user_controller.GetSave)                // get my saves
	r.POST("/saves/:id/:EventID", user_controller.AddSave)      // add event to my saves
	r.DELETE("/saves/:id/:EventID", user_controller.DeleteSave) // delete

	r.POST("/courses/:id/create", event_controller.CreateCourse)                 // create course
	r.POST("/courses/:id/:CourseID/:EventID", event_controller.AddEventToCourse) // add event to course
	r.DELETE("/courses/:id/:CourseID", event_controller.DeleteCourse)            // delete course

	r.GET("/courses", event_controller.GetPublicCourse)                  // get public course
	r.GET("/courses/:id/outline", event_controller.GetCourseOutlineInfo) // get course detail
	r.GET("/courses/:id", event_controller.GetCourseFullInfo)            // get course detail

	r.GET("/users/:id/profile-image", user_controller.GetProfileImage)
	r.GET("/users/:id/nickname", user_controller.GetNickname)

}

func InitServer() {
	r := gin.Default()

	initRouterHandler(r)

	r.Run()
}
