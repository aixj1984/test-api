package test

import (
	"test-api/models"

	"test-api/providers"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	//	"test-api/payloads"
)

func ListCourseTest(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	test_type, _ := strconv.Atoi(c.DefaultQuery("testtype", "0"))
	course_id, _ := strconv.Atoi(c.DefaultQuery("courseid", "0"))

	if course_id > 11 || course_id == 0 {
		c.JSON(200, gin.H{
			"code":  100,
			"msg":   "课程ID不对",
			"count": 0,
			"data":  nil,
		})
		return
	}

	var coursetest models.CourseTest

	var coursetests []*models.CourseTest

	_, err := providers.Test.GetMore(&coursetests, coursetest, course_id, test_type, (start-1)*length, length)
	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"code":  101,
			"msg":   err.Error(),
			"count": 0,
			"data":  nil,
		})
		return
	}

	count, err := providers.Test.Count(&coursetest, course_id, test_type)

	if err != nil {
		log.Println(err.Error())
		c.JSON(200, gin.H{
			"code":  101,
			"msg":   err.Error(),
			"count": 0,
			"data":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "",
		"count": count,
		"data":  coursetests,
	})

}
