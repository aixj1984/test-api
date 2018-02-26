package providers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

	//	"errors"
)

type ICustomerCourseProvider interface {
	GetOne(interface{}, ...string) error
	InsertOne(m interface{}) (int64, error)
	GetMore(array interface{}, customer_id int, start, length int) (int64, error)
	UpdateOne(m interface{}, cols ...string) (int64, error)
	GetAll(array interface{}, customer_id int) (int64, error)
	Count(customer_id int) (int64, error)
	DeleteOne(m interface{}) (int64, error)
	UpdataDefault(customer_id int, default_courses string) (int64, error)
}

type CustomerCourseProvider struct {
}

func (p *CustomerCourseProvider) GetOne(m interface{}, cols ...string) error {
	o := orm.NewOrm()

	err := o.Read(m, cols...)

	return err
}

func (p *CustomerCourseProvider) InsertOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Insert(m)

	return effact, err
}

func (p *CustomerCourseProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *CustomerCourseProvider) DeleteOne(m interface{}) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Delete(m)

	return effact, err
}

func (p *CustomerCourseProvider) UpdataDefault(customer_id int, default_courses string) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("update customer_course set is_display = 1 where customer_id = ? and course_id in ("+default_courses+")", customer_id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		o.Raw("update customer_course set is_display = 0 where customer_id = ? and course_id not in ("+default_courses+")", customer_id).Exec()
		return num, nil
	}

	return 0, err
}

func (p *CustomerCourseProvider) GetMore(array interface{}, customer_id int, start, length int) (int64, error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "course.status = 1"

	qb.Select("course.id", "name", "price", "is_display",
		"IFNULL(temp_a.course_id, 0) as status").
		From("course").
		LeftJoin("(select * from customer_course where customer_id = ? ) as temp_a").On("course.id = temp_a.course_id").
		Where(condition).
		Limit(length).Offset(start)

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()

	effact, err := o.Raw(sql, customer_id).QueryRows(array)
	return effact, err

}

func (p *CustomerCourseProvider) GetAll(array interface{}, customer_id int) (int64, error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "course.status = 1"

	qb.Select("course.id", "name", "price",
		"IFNULL(temp_a.course_id, 0) as status").
		From("course").
		LeftJoin("(select id, course_id from customer_course where customer_id = ? ) as temp_a").On("course.id = temp_a.course_id").
		Where(condition)

	// 导出 SQL 语句
	sql := qb.String()

	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()

	effact, err := o.Raw(sql, customer_id).QueryRows(array)
	return effact, err

}

func (p *CustomerCourseProvider) Count(customer_id int) (int64, error) {
	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "course.status = 1 "

	qb.Select("count(*) as num ").
		From("course").
		LeftJoin("(select * from customer_course where customer_id = ? ) as temp_a").On("course.id = temp_a.course_id").
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
