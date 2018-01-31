package models

import (
	"test-api/comm/beelog"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Init  初始化model
func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", "root:rootroot@tcp(111.230.221.146:3310)/neihe_test?charset=utf8")

	orm.RegisterModel(new(Account), new(WxUserinfo),
		new(Question1), new(Question2), new(Question3),
		new(Question4), new(Question5), new(Question6),
		new(Question7), new(Question8), new(Question9),
		new(Question10), new(Question11),
		new(CollectQuestion1), new(CollectQuestion2), new(CollectQuestion3),
		new(CollectQuestion4), new(CollectQuestion5), new(CollectQuestion6),
		new(CollectQuestion7), new(CollectQuestion8), new(CollectQuestion9),
		new(CollectQuestion10), new(CollectQuestion11),
		new(CourseTest), new(TestQuestion), new(Article),
		new(Course), new(CustomerCourse), new(TestResult))

	orm.SetMaxIdleConns("default", 5)

	orm.SetMaxOpenConns("default", 10)

	//create table
	orm.RunSyncdb("default", false, true)

	beelog.Info("init db info  ..... ")

}
