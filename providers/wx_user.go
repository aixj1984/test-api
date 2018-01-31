package providers

import (
	"test-api/models"

	"github.com/astaxie/beego/orm"

	"fmt"

	//"errors"
)

//IAccountProvider account provider interface
type IWxProvider interface {
	GetOne(interface{}) error
	CheckOpenID(string) bool
	InsertOne(interface{}) (int64, error)
}

//WxProvider wx_userinfo provider
type WxProvider struct {
}

//Get 获取wx_userinfo
func (p *WxProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()
	err := o.Read(m)

	return err
}

func (p *WxProvider) CheckOpenID(openID string) bool {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.WxUserinfo)).Filter("openid", openID).Count()
	if err == nil && num > 0 {
		return true
	}
	return false

}

func (p *WxProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *WxProvider) UpdateOne(m *models.WxUserinfo) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("update  wx_userinfo set title = ?, content = ? ,source = ? ,abstract = ? where id = ? ").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}
