package course

import (
	"test-api/models"

	"test-api/providers"

	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"test-api/payloads"
)

func ListCourse(c *gin.Context) {

	start, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	length, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	uid, err := c.Request.Cookie("UID")

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	customer_id, err := strconv.Atoi(uid.Value)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	//customer_id, _ := strconv.Atoi(c.DefaultQuery("customerid", "0"))

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

	_, err = providers.CustomerCourse.GetMore(&courses, customer_id, (start-1)*length, length)
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

	uid, err := c.Request.Cookie("UID")

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	iUid, err := strconv.Atoi(uid.Value)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	var payload payloads.SaveCustomerCoursesSetting
	if c.ShouldBind(&payload) == nil {
		_, err := providers.CustomerCourse.UpdataDefault(iUid, payload.DefalutCourses)
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

func CustomerPurchaseCourses(c *gin.Context) {

	var payload payloads.PurchaseCourses
	if c.ShouldBind(&payload) == nil {
		course_array := strings.Split(payload.PurchaseCourses, ",")
		var course_id_array []int
		for i := 0; i < len(course_array); i++ {
			course_id_s := strings.TrimSpace(course_array[i])
			course_id, err := strconv.Atoi(course_id_s)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 101,
					"msg":  err.Error(),
				})
				return
			}
			if course_id <= 0 || course_id > 11 {
				c.JSON(200, gin.H{
					"code": 101,
					"msg":  "参数错误1",
				})
				return
			}
			course_id_array = append(course_id_array, course_id)
		}
		if len(course_id_array) == 0 {
			c.JSON(200, gin.H{
				"code": 101,
				"msg":  "参数错误2",
			})
			return
		}

		total_mount := len(course_id_array) * 120 * (100 - (len(course_id_array)-1)*5) / 100

		log.Println(total_mount)

		for i := 0; i < len(course_id_array); i++ {
			var customer_course = models.CustomerCourse{}
			customer_course.CustomerId = 1
			customer_course.CourseId = course_id_array[i]
			customer_course.IsDisplay = 1

			providers.CustomerCourse.InsertOne(&customer_course)

		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "",
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 101,
		"msg":  "参数异常3",
	})

}
