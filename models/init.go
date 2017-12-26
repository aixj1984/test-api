package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

//Init  初始化model
func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", "dev_richard:dev_richard@tcp(10.0.12.239:3310)/test?charset=utf8")

	orm.RegisterModel(new(Account),
		new(Question1), new(Question2), new(Question3),
		new(Question4), new(Question5), new(Question6),
		new(Question7), new(Question8), new(Question9),
		new(Question10), new(Question11),
		new(CollectQuestion1), new(CollectQuestion2), new(CollectQuestion3),
		new(CollectQuestion4), new(CollectQuestion5), new(CollectQuestion6),
		new(CollectQuestion7), new(CollectQuestion8), new(CollectQuestion9),
		new(CollectQuestion10), new(CollectQuestion11),
		new(CourseTest), new(TestQuestion),
		new(Course), new(CustomerCourse), new(TestResult))

	orm.SetMaxIdleConns("default", 5)

	orm.SetMaxOpenConns("default", 10)

	//create table
	orm.RunSyncdb("default", false, true)

	beego.Info("init db info  ..... ")

}
