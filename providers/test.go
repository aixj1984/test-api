package providers

import (
	"github.com/astaxie/beego/orm"
)

//IAccountProvider account provider interface
type ITestProvider interface {
	GetOne(interface{}) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array, object interface{}, course_id, test_type int, start, length int) (int64, error)
	UpdateOne(m interface{}, cols ...string) (int64, error)
	Count(object interface{}, course_id, test_type int) (int64, error)
}

//AccountProvider account provider
type TestProvider struct {
}

func (p *TestProvider) GetOne(m interface{}) error {
	o := orm.NewOrm()

	err := o.Read(m)

	return err
}

func (p *TestProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *TestProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *TestProvider) GetMore(array interface{}, object interface{}, course_id, test_type int, start, length int) (int64, error) {
	o := orm.NewOrm()
	// LIMIT 10 OFFSET 20 注意跟 SQL 反过来的

	effact, err := o.QueryTable(object).Filter("status", 1).Filter("course_id", course_id).Filter("test_type", test_type).OrderBy("-id").Limit(length, start).All(array)

	return effact, err
}

func (p *TestProvider) Count(object interface{}, course_id, test_type int) (int64, error) {
	o := orm.NewOrm()

	count, err := o.QueryTable(object).Filter("status", 1).Filter("course_id", course_id).Filter("test_type", test_type).Count()

	return count, err
}
