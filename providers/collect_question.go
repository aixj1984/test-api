package providers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

	"test-api/models"

	//	"errors"
)

type ICollectQuestionProvider interface {
	GetOne(interface{}, ...string) error
	InsertOne(m *models.CollectQuestion, tablename string) (int64, error)
	GetMore(array interface{}, question_tablename, collection_tablename string, customer_id int, start, length int) (int64, error)
	Count(question_tablename, collection_tablename string, customer_id int) (int64, error)
	DeleteOne(question_id, customer_id int, tablename string) (int64, error)
}

type CollectQuestionProvider struct {
}

func (p *CollectQuestionProvider) GetOne(m interface{}, cols ...string) error {
	o := orm.NewOrm()

	err := o.Read(m, cols...)

	return err
}

func (p *CollectQuestionProvider) InsertOne(m *models.CollectQuestion, tablename string) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("insert into "+tablename+" (customer_id, question_id,add_time) values(?,?,?)", m.CustomerId, m.QuestionId, m.AddTime).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}

	return 0, err

}

func (p *CollectQuestionProvider) UpdateOne(m interface{}, cols ...string) (int64, error) {
	o := orm.NewOrm()

	effact, err := o.Update(m, cols...)

	return effact, err
}

func (p *CollectQuestionProvider) DeleteOne(question_id, customer_id int, tablename string) (int64, error) {
	o := orm.NewOrm()

	res, err := o.Raw("delete from "+tablename+" where  question_id = ? and customer_id=? ", question_id, customer_id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		return num, nil
	}

	return 0, err
}

func (p *CollectQuestionProvider) GetMore(array interface{}, question_tablename, collection_tablename string, customer_id int, start, length int) (int64, error) {

	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "status = 1"

	qb.Select("title",
		question_tablename+".id",
		"question_type",
		"options",
		"answer",
		"note",
		"status").
		From(question_tablename).
		InnerJoin("(select distinct(question_id) from " + collection_tablename + " where customer_id = ?) as temp_a ").On(question_tablename + ".id = temp_a.question_id").
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

func (p *CollectQuestionProvider) Count(question_tablename, collection_tablename string, customer_id int) (int64, error) {
	qb, _ := orm.NewQueryBuilder("mysql")

	var condition = "status = 1 "

	qb.Select("count(*) as num").
		From(question_tablename).
		InnerJoin("(select distinct(question_id) from " + collection_tablename + " where customer_id = ?) as temp_a ").On(question_tablename + ".id = temp_a.question_id").
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
