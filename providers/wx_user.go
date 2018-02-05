package providers

import (
	"test-api/models"

	"github.com/astaxie/beego/orm"

	"fmt"

	"errors"
)

//IAccountProvider account provider interface
type IWxProvider interface {
	GetOne(interface{}) error
	CheckOpenID(string) bool
	InsertOne(interface{}) (int64, error)
}

type IWxpayProvider interface {
	GetOne(interface{}) error
	UpdateOne(m interface{}, cols ...string) (int64, error)
	InsertOne(interface{}) (int64, error)
	GetOneByCondition(interface{}, string, interface{}) error
	GetAll(interface{}, string) (int64, error)
}

type WxPayProvider struct {
}

//Get 获取wx_order
func (p *WxPayProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()
	err := o.Read(m)

	return err
}

//Get 获取account
func (p *WxPayProvider) GetOneByCondition(m interface{}, fileld string, value interface{}) error {
	o := orm.NewOrm()

	err := o.QueryTable("order").Filter(fileld, value).One(m)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return errors.New("find more not one")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return errors.New("not exist")
	}

	return err
}

func (p *WxPayProvider) GetAll(m interface{}, order_no string) (int64, error) {
	var num int64
	var err error
	o := orm.NewOrm()

	if num, err = o.QueryTable("order_course").Filter("order_no", order_no).All(m); num > 0 && err != nil {
		return num, nil
	}
	return 0, err
}

func (p *WxPayProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *WxPayProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
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
