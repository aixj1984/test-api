package providers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

	//	"errors"
)

type ITestResultProvider interface {
	GetOne(interface{}, ...string) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array interface{}, customer_id int, start, length int) (int64, error)
	Count(customer_id int) (int64, error)
}

type TestResultProvider struct {
}

func (p *TestResultProvider) GetOne(m interface{}, cols ...string) error {
	o := orm.NewOrm()

	err := o.Read(m, cols...)

	return err
}

func (p *TestResultProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *TestResultProvider) GetMore(array interface{}, customer_id int, start, length int) (int64, error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = " customer_id  = ? "

	qb.Select("test_result.id", "name", "question_num", "right_num", "add_time", "test_sec").
		From("test_result").
		InnerJoin("course").On("test_result.course_id = course.id ").
		Where(condition).
		OrderBy("test_result.id desc").
		Limit(length).Offset(start)

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()

	effact, err := o.Raw(sql, customer_id).QueryRows(array)
	return effact, err

}

func (p *TestResultProvider) Count(customer_id int) (int64, error) {
	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = " customer_id  = ? "

	qb.Select("count(*) as num ").
		From("test_result").
		InnerJoin("course").On("test_result.course_id = course.id ").
		Where(condition)

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql, customer_id).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps[0]["num"])
		count, _ := strconv.Atoi(maps[0]["num"].(string))
		return int64(count), nil
	} else {
		return 0, err
	}
}
