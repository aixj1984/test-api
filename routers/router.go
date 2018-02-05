package routers

import (
	"log"
	"strings"

	"test-api/comm/beelog"
	"test-api/controllers"
	"test-api/controllers/article"
	"test-api/controllers/course"
	"test-api/controllers/test"
	"test-api/controllers/wechat"

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

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "api") && c.Request.RequestURI != "/api/wechat/login" && c.Request.RequestURI != "/api/wechat/callback" {
			//			beelog.Debug("test")
			if uid, err := c.Request.Cookie("UID"); err != nil {
				beelog.Debug(err)
				c.AbortWithStatus(401)
			} else {
				log.Println(uid)
			}
		}

	}
}

func init() {
	router := gin.Default()

	router.Use(CORSMiddleware())
	//router.Use(Check())

	// Simple group: api
	api := router.Group("/api")
	{
		api.POST("/course/list", controllers.Test)

		api.GET("/course/list", course.ListCourse)
		api.GET("/courses/list", course.ListAllCourse)
		api.PUT("/course/setting", course.CustomerDefaultCourse)
		api.PUT("/course/purchase", course.CustomerPurchaseCourses)

		api.GET("/coursetest/list", test.ListCourseTest)

		api.GET("/test/question/list", test.ListTestQuestion)

		api.GET("/collect/question/list", test.ListCollectQuestion)
		api.POST("/collect/question/del", test.DelCollectQuestion)
		api.POST("/collect/question/add", test.SaveCollectQuestion)

		api.PUT("/testresult/add", test.SaveTestResult)
		api.GET("/testresult/list", test.ListTestResult)

		api.GET("/article/list", article.ListArticle)

		api.GET("article/detail", article.GetArticleDetail)

		api.GET("/wechat/login", wechat.Login)
		api.GET("/wechat/callback", wechat.CallBack)

	}

	router.Static("/static", "./static")
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.StaticFile("/", "./views/index.html")

	router.Run(":8000")

}
