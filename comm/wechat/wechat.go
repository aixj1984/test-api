package wechat

import (
	"test-api/comm/beelog"

	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

type WxServer struct {
	WxAppId           string
	WxAppSerect       string
	Oauth2RedirectURI string
	Oauth2Scope       string
}

var wxServer *WxServer

func init() {
	init_service()
}

func init_service() {
	if wxServer == nil {
		wxServer = new(WxServer)
	}
	wxServer.WxAppId = "wxa2b17d163cc88e9d"
	wxServer.WxAppSerect = "30d51b4a18121cc0a7f27b5a8f03d588"
	wxServer.Oauth2RedirectURI = "http://testing.foxhelper.cn/api/wechat/callback"
	wxServer.Oauth2Scope = "snsapi_userinfo"
}

//获取微信服务
func GetwxService() *WxServer {
	if wxServer == nil {
		init_service()
	}

	return wxServer
}

func (serv *WxServer) GetRedirectUrl(state string) string {
	AuthCodeURL := mpoauth2.AuthCodeURL(serv.WxAppId, serv.Oauth2RedirectURI, serv.Oauth2Scope, state)
	beelog.Info("AuthCodeURL:", AuthCodeURL)
	return AuthCodeURL
}

func (serv *WxServer) GetUserInfo(code, state string) interface{} {
	oauth2Client := oauth2.Client{
		Endpoint: mpoauth2.NewEndpoint(serv.WxAppId, serv.WxAppSerect),
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		beelog.Error(err)
		return nil
	}
	beelog.Debug("token: %+v\r\n", token)

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		beelog.Error(err)
		return nil
	}
	return userinfo
}
