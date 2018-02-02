package test

import (
	"test-api/models"

	"test-api/providers"

	"log"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"

	"test-api/payloads"
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

	var questions []*models.CQuestion

	if count == 0 {
		length = 50
		count = 50
		_, err = providers.TestQuestion.GetRandom(&questions, models.CourseMap[courseid], models.CourseCollectMap[courseid], testid, length)
	} else {
		_, err = providers.TestQuestion.GetMore(&questions, "", "1", models.CourseMap[courseid], models.CourseCollectMap[courseid], testid, (start-1)*length, length)
	}

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

func ListCollectQuestion(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	courseid := c.DefaultQuery("courseid", "0")
	i_courseid, _ := strconv.Atoi(courseid)
	if i_courseid == 0 || i_courseid > 11 {
		c.JSON(200, gin.H{
			"code":  100,
			"msg":   "参数不对",
			"count": 0,
			"data":  nil,
		})
		return
	}

	var questions []*models.Question

	_, err := providers.CollectQuestion.GetMore(&questions, models.CourseMap[courseid], models.CourseCollectMap[courseid], 1, (start-1)*length, length)

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

	count, err := providers.CollectQuestion.Count(models.CourseMap[courseid], models.CourseCollectMap[courseid], 1)

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

func SaveCollectQuestion(c *gin.Context) {
	var payload payloads.SaveQuestionCollect
	if c.ShouldBind(&payload) == nil {
		var collect_question = models.CollectQuestion{}
		collect_question.CustomerId = 1
		collect_question.AddTime = time.Now().Format("2006-01-02 15:04:05")
		collect_question.QuestionId = payload.QuestionId

		_, err := providers.CollectQuestion.InsertOne(&collect_question, models.CourseCollectMap[strconv.Itoa(payload.CourseId)])
		if err != nil {

			c.JSON(200, gin.H{
				"code": 101,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 101,
		"msg":  "参数异常",
	})
}

func DelCollectQuestion(c *gin.Context) {
	var payload payloads.SaveQuestionCollect
	if c.ShouldBind(&payload) == nil {

		_, err := providers.CollectQuestion.DeleteOne(payload.QuestionId, 1, models.CourseCollectMap[strconv.Itoa(payload.CourseId)])
		if err != nil {

			c.JSON(200, gin.H{
				"code": 101,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 101,
		"msg":  "参数异常",
	})
}
