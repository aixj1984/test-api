package test

import (
	"test-api/models"

	"test-api/providers"

	"test-api/payloads"

	"log"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

func ListTestResult(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var test_results []*models.TestResultDetail

	var customer_id = 1

	_, err := providers.TestResult.GetMore(&test_results, customer_id, (start-1)*length, length)
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

	count, err := providers.TestResult.Count(customer_id)

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
		"data":  test_results,
	})

}

func SaveTestResult(c *gin.Context) {
	var payload payloads.SaveTestResult
	if c.ShouldBind(&payload) == nil {

		var test_reslt = models.TestResult{}
		test_reslt.CustomerId = 1
		test_reslt.AddTime = time.Now().Format("2006-01-02 15:04:05")
		test_reslt.TestSec = payload.TestSec
		test_reslt.CourseId = payload.CourseId
		test_reslt.TestId = payload.TestId
		test_reslt.QuestionNum = payload.QuestionNum
		test_reslt.RightNum = payload.RightNum

		_, err := providers.TestResult.InsertOne(&test_reslt)
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
