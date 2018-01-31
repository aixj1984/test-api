package wechat

import (
	"net/http"

	"test-api/comm/beelog"
	"test-api/comm/wechat"
	"test-api/models"
	"test-api/providers"

	"github.com/chanxuehong/rand"
	"github.com/gin-gonic/gin"
	"gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
)

func Login(c *gin.Context) {
	html_url := c.Query("html_url")
	if html_url == "" {
		html_url = "/"
	}
	service := wechat.GetwxService()
	redirectUrl := service.GetRedirectUrl(html_url)
	c.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func CallBack(c *gin.Context) {
	var wx_userinfo = models.WxUserinfo{}

	service := wechat.GetwxService()
	code := c.Query("code")
	state := c.Query("state")

	userinfo := service.GetUserInfo(code, state)
	beelog.Debug(userinfo)
	if userinfo != nil {
		openid := userinfo.(*oauth2.UserInfo).OpenId
		wx_userinfo.OpenId = openid
		beelog.Info(wx_userinfo)
		if isexist := providers.Wx.CheckOpenID(openid); !isexist {
			wx_userinfo.Nickname = userinfo.(*oauth2.UserInfo).Nickname
			wx_userinfo.City = userinfo.(*oauth2.UserInfo).City
			wx_userinfo.HeadImageURL = userinfo.(*oauth2.UserInfo).HeadImageURL
			wx_userinfo.Country = userinfo.(*oauth2.UserInfo).Country
			wx_userinfo.Province = userinfo.(*oauth2.UserInfo).Province
			wx_userinfo.Sex = userinfo.(*oauth2.UserInfo).Sex
			wx_userinfo.UnionId = userinfo.(*oauth2.UserInfo).UnionId
			if num, err := providers.Wx.InsertOne(&wx_userinfo); num > 0 && err == nil {
				account := new(models.Account)
				account.OpenID = userinfo.(*oauth2.UserInfo).OpenId
				account.Nickname = userinfo.(*oauth2.UserInfo).Nickname
				account.CountryCode = userinfo.(*oauth2.UserInfo).Country
				account.AccountSrc = 1
				account.Password = string(rand.NewHex())
				if _, err1 := providers.Account.InsertOne(account); err1 != nil {
					beelog.Error(err1)
					c.JSON(200, gin.H{
						"code": 1000,
						"msg":  err1.Error(),
					})
					return
				}
			}

		}
	}

	c.Redirect(http.StatusMovedPermanently, state)
}
