package wechat

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"test-api/comm/beelog"
	"test-api/comm/utils"
	//	"test-api/comm/myredis"
	"test-api/comm/wechat"
	"test-api/models"
	"test-api/payloads"
	"test-api/providers"

	"github.com/chanxuehong/rand"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
)

func Test(c *gin.Context) {
	uid := c.Query("UID")
	state := c.Query("html_url")
	beelog.Debug(uid, state)
	cookie := &http.Cookie{
		Name:  "UID",
		Value: uid,
	}
	http.SetCookie(c.Writer, cookie)
	c.Redirect(http.StatusMovedPermanently, state)
}

func Pay(c *gin.Context) {
	var (
		uid string
		//		err     error
		payload payloads.PurchaseCourses
		//		money   int
		//		name    int
	)
	if cookie, err := c.Request.Cookie("UID"); err != nil {
		beelog.Debug(err)
		c.AbortWithStatus(401)
		return
	} else {
		uid = cookie.Value
	}
	id, _ := strconv.Atoi(uid)
	account := new(models.Account)
	if err := providers.Account.GetOneByCondition(account, "id", id); err != nil {
		c.JSON(200, gin.H{
			"code": 1001,
			"msg":  "user not exist",
		})
		return
	}
	if c.ShouldBind(&payload) == nil {
		beelog.Debug(payload)
		course_array := strings.Split(payload.PurchaseCourses, ",")
		var course_id_array []int
		for i := 0; i < len(course_array); i++ {

			course_id_s := strings.TrimSpace(course_array[i])
			beelog.Debug(course_id_s)
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
		clientIP := c.ClientIP()
		//		if uid, err = c.Cookie("UID"); err != nil {
		//			c.JSON(401, gin.H{
		//				"code": 1001,
		//				"msg":  "user not login",
		//			})
		//			return
		//		}

		var wx_order models.Order
		wx_order.OrderNo = utils.GetOrderNo()
		for i := 0; i < len(course_id_array); i++ {
			var order_course models.OrderCourse
			order_course.OrderNo = wx_order.OrderNo
			order_course.CourseId = course_id_array[i]
			if _, err := providers.WxPay.InsertOne(&order_course); err != nil {
				beelog.Error(err)
			}
		}
		//		wx_order.Uid = uid
		wx_order.ClientIp = clientIP
		wx_order.FeeType = "CNY"
		wx_order.TotalFee = 100
		wx_order.OpenId = account.OpenID
		wx_order.Name = "test"
		wx_order.TradeType = "JSAPI"
		service := wechat.GetwxPayService()
		if ret, err := service.UnifiedOrder(total_mount, wx_order.Name, wx_order.ClientIp, wx_order.OpenId, wx_order.FeeType, wx_order.TradeType, service.NotifyUrl, wx_order.OrderNo); err != nil {
			c.JSON(200, gin.H{
				"code": 1002,
				"msg":  err.Error(),
			})
			return
		} else {
			packages := ret["package"]
			prepayId := strings.Split(packages, "=")
			wx_order.PrepayId = prepayId[1]
			if _, err := providers.WxPay.InsertOne(&wx_order); err != nil {
				beelog.Error(err)
			}
			c.JSON(200, ret)
			return
		}

	} else {
		c.JSON(200, gin.H{
			"code": 101,
			"msg":  "参数异常",
		})
		return
	}

}
func PayCallback(c *gin.Context) {
	var response map[string]interface{}
	c.Bind(&response)
	beelog.Debug(response)
	paycallback(response["prepayId"].(string))
	c.XML(200, gin.H{
		"return_code": "<![CDATA[SUCCESS]]>",
		"return_msg":  "<![CDATA[OK]]>",
	})
	return
}

func paycallback(prepayId string) {
	var wx_order models.Order
	var account models.Account

	if err := providers.WxPay.GetOneByCondition(&wx_order, "prepayId", prepayId); err != nil {
		if err := providers.Account.GetOneByCondition(&account, "openid", wx_order.OpenId); err != nil {
			beelog.Error(err)
			return
		}
		var order_course []models.OrderCourse
		if _, err1 := providers.WxPay.GetAll(&order_course, wx_order.OrderNo); err1 == nil {
			for i := 0; i < len(order_course); i++ {
				var customer_course models.CustomerCourse
				customer_course.CourseId = order_course[i].CourseId
				customer_course.CustomerId = account.Id
				if _, err := providers.CustomerCourse.InsertOne(&customer_course); err != nil {
					beelog.Error(err)
				}
			}
		} else {
			beelog.Error(err1)
		}
	}
}

func Login(c *gin.Context) {
	html_url := c.Query("html_url")
	if html_url == "" {
		html_url = "/"
	}
	service := wechat.GetwxService()
	redirectUrl := service.GetRedirectUrl(html_url)
	c.Redirect(http.StatusMovedPermanently, redirectUrl)
	return
}

func CallBack(c *gin.Context) {

	var wx_userinfo = models.WxUserinfo{}
	//	var cookie *http.Cookie
	redirectUrl := "http://testing.foxhelper.cn/api/test"

	service := wechat.GetwxService()
	code := c.Query("code")
	state := c.Query("state")
	redirectUrl = redirectUrl + "?html_url=" + state
	//	cache_service := myredis.GetCache()
	userinfo := service.GetUserInfo(code, state)
	beelog.Debug(userinfo)
	if userinfo != nil {
		openid := userinfo.(*oauth2.UserInfo).OpenId
		wx_userinfo.OpenId = openid
		beelog.Info(wx_userinfo)
		account := new(models.Account)
		if err := providers.Account.GetOneByCondition(account, "openId", openid); err != nil {
			beelog.Error(err)
			if isexist := providers.Wx.CheckOpenID(openid); !isexist {
				wx_userinfo.Nickname = userinfo.(*oauth2.UserInfo).Nickname
				wx_userinfo.City = userinfo.(*oauth2.UserInfo).City
				wx_userinfo.HeadImageURL = userinfo.(*oauth2.UserInfo).HeadImageURL
				wx_userinfo.Country = userinfo.(*oauth2.UserInfo).Country
				wx_userinfo.Province = userinfo.(*oauth2.UserInfo).Province
				wx_userinfo.Sex = userinfo.(*oauth2.UserInfo).Sex
				wx_userinfo.UnionId = userinfo.(*oauth2.UserInfo).UnionId
				if _, err := providers.Wx.InsertOne(&wx_userinfo); err == nil {
					account.OpenID = userinfo.(*oauth2.UserInfo).OpenId
					account.Nickname = userinfo.(*oauth2.UserInfo).Nickname
					account.CountryCode = userinfo.(*oauth2.UserInfo).Country
					account.AccountSrc = 1
					account.Password = string(rand.NewHex())
					account.DeletedAt = time.Now()
					if id, err1 := providers.Account.InsertOne(account); err1 != nil {
						beelog.Error(err1)
						c.JSON(200, gin.H{
							"code": 1000,
							"msg":  err1.Error(),
						})
						return
					} else {
						uid := strconv.FormatInt(id, 10)
						redirectUrl = redirectUrl + "&UID=" + uid
					}
				} else {
					beelog.Error(err)
				}

			}
		} else {
			beelog.Debug(account)
			redirectUrl = redirectUrl + "&UID=" + strconv.Itoa(account.Id)
		}

	} else {
		c.JSON(200, gin.H{
			"code": 101,
			"msg":  "wechat login error",
		})
	}

	c.Redirect(http.StatusFound, redirectUrl)
	return
}
