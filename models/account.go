package models

import (
	//	"database/sql"
	"time"
)

//Account account model
type Account struct {
	Id          int    `orm:"pk"`
	ImToken     string `orm:"column(imtoken);size(128)"`
	CountryCode string `orm:"column(countrycode);size(20)",json:"countrycode"`
	PhoneNo     string `orm:"column(phoneno);size(20)"`
	OpenID      string `orm:"column(openid);size(64);unique"`
	//0：inner; 微信: 1; 微博:2 ; QQ :3
	AccountSrc   int8      `orm:"column(accountsrc)"`
	Nickname     string    `orm:"column(nickname);size(50)",json:"nickname"`
	Email        string    `orm:"column(email);size(30)"`
	PendingEmail string    `orm:"column(pendingemail);size(30)"` //新邮箱
	Password     string    `orm:"column(password);size(50)"`
	CreateAt     time.Time `orm:"column(create_at);auto_now_add;type(datetime)"`
	UpdateAt     time.Time `orm:"column(update_at);auto_now;type(datetime)"`
	DeletedAt    time.Time `orm:"column(deleted_at);type(datetime)"`
}

//TableName table name
func (m *Account) TableName() string {
	return "account"
}

type WxUserinfo struct {
	OpenId   string `orm:"pk;column(openid);size(64)",json:"openid"`  // 用户的唯一标识
	Nickname string `orm:"column(nickname);size(50)",json:"nickname"` // 用户昵称
	Sex      int    `orm:"column(sex)",json:"sex"`                    // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City     string `orm:"column(city);size(50)",json:"city"`         // 普通用户个人资料填写的城市
	Province string `orm:"column(province);size(20)",json:"province"` // 用户个人资料填写的省份
	Country  string `orm:"column(country);size(20)",json:"country"`   // 国家, 如中国为CN

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	HeadImageURL string `orm:"column(headimgurl);size(50)",json:"headimgurl,omitempty"`

	//	Privilege []string `xorm:" varchar(50) null 'privilege'",json:"privilege,omitempty"` // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionId string `orm:"column(unionid);size(50)",json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

func (w *WxUserinfo) TableName() string {
	return "wx_userinfo"
}
