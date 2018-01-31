package providers

import (
	"test-api/models"

	"fmt"

	"github.com/astaxie/beego/orm"

	//"errors"
)

//IAccountProvider account provider interface
type IAccountProvider interface {
	GetOne(interface{}) error
	InsertOne(interface{}) (int64, error)
}

//AccountProvider account provider
type AccountProvider struct {
}

//Get 获取account
func (p *AccountProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()
	/*
		err := o.QueryTable("account").Filter("id", 1).One(m)
		if err == orm.ErrMultiRows {
			// 多条的时候报错
			return errors.New("test")
		}
		if err == orm.ErrNoRows {
			// 没有找到记录
			return errors.New("test111")
		}*/
	err := o.Read(m)

	return err
}

func (p *AccountProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *AccountProvider) UpdateOne(m *models.Account) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("update  account set title = ?, content = ? ,source = ? ,abstract = ? where id = ? ").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}
	return 0, err
}
