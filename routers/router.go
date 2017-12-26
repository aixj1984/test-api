package routers

import (
	"log"
	"test-api/controllers"

	"test-api/controllers/test"

	"test-api/controllers/course"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println(c.GetHeader("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func init() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Simple group: api
	api := router.Group("/api")
	{
		api.POST("/course/list", controllers.Test)

		api.GET("/course/list", course.ListCourse)
		api.GET("/courses/list", course.ListAllCourse)
		api.PUT("/course/setting", course.CustomerDefaultCourse)

		api.GET("/coursetest/list", test.ListCourseTest)

		api.GET("/test/question/list", test.ListTestQuestion)

		api.GET("/collect/question/list", test.ListCollectQuestion)
		api.POST("/collect/question/del", test.DelCollectQuestion)
		api.POST("/collect/question/add", test.SaveCollectQuestion)

		api.PUT("/testresult/add", test.SaveTestResult)
		api.GET("/testresult/list", test.ListTestResult)

	}

	router.Static("/static", "./static")
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.StaticFile("/", "./views/index.html")

	router.Run(":80")

}
