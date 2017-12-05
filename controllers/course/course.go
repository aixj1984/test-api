package course

import (
	"test-api/models"

	"test-api/providers"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"test-api/payloads"
)

func ListCourse(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	customer_id, _ := strconv.Atoi(c.DefaultQuery("customerid", "0"))

	if customer_id == 0 {
		c.JSON(200, gin.H{
			"code":  100,
			"msg":   "客户ID没有给",
			"count": 0,
			"data":  nil,
		})
		return
	}

	log.Println(start, length)

	var courses []*models.CustomerCourseDetail

	_, err := providers.CustomerCourse.GetMore(&courses, customer_id, (start-1)*length, length)
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

	count, err := providers.CustomerCourse.Count(customer_id)

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
		"data":  courses,
	})

}

func ListAllCourse(c *gin.Context) {

	customer_id, _ := strconv.Atoi(c.DefaultQuery("customerid", "0"))

	if customer_id == 0 {
		c.JSON(200, gin.H{
			"code":  100,
			"msg":   "客户ID没有给",
			"count": 0,
			"data":  nil,
		})
		return
	}

	var courses []*models.CustomerCourseDetail

	_, err := providers.CustomerCourse.GetAll(&courses, customer_id)
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
		"code": 0,
		"msg":  "",
		"data": courses,
	})

}

func CustomerDefaultCourse(c *gin.Context) {

	var payload payloads.SaveCustomerCoursesSetting
	if c.ShouldBind(&payload) == nil {

		_, err := providers.CustomerCourse.UpdataDefault(payload.CustomerId, payload.DefalutCourses)
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
