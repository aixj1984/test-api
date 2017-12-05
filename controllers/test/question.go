package test

import (
	"test-api/models"

	"test-api/providers"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	//	"test-api/payloads"
)

func ListTestQuestion(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	testid := c.DefaultQuery("testid", "0")
	courseid := c.DefaultQuery("courseid", "0")
	i_courseid, _ := strconv.Atoi(courseid)
	i_testid, _ := strconv.Atoi(testid)
	if i_testid == 0 || i_courseid == 0 || i_courseid > 11 {
		c.JSON(200, gin.H{
			"code":  100,
			"msg":   "参数不对",
			"count": 0,
			"data":  nil,
		})
		return
	}

	var questions []*models.Question

	_, err := providers.TestQuestion.GetMore(&questions, "", "1", models.CourseMap[courseid], testid, (start-1)*length, length)

	//count, err := providers.TestQuestion.Count(question, this.GetString("query_key"), this.GetString("question_status"), models.CourseMap[this.GetString("course_id")], this.GetString("test_id"))
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

	count, err := providers.TestQuestion.Count("", "1", models.CourseMap[courseid], testid)

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
		"data":  questions,
	})

}
