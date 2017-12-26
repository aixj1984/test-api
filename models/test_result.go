package models

//	"time"

type TestResult struct {
	Id          int
	CustomerId  int    `xorm:"int 'customer_id'"`
	CourseId    int    `xorm:"int 'course_id'"`
	TestId      int    `xorm:"int   'test_id'"`
	QuestionNum int    `xorm:"int  'question_num'"`
	RightNum    int    `xorm:"int  'right_num'"`
	AddTime     string `xorm:"varchar(60)  'add_time'"`
	TestSec     int    `xorm:"int  'test_sec'"`
}

type TestResultDetail struct {
	TestResult `xorm:"extends"`
	Name       string `xorm:"varchar(60)  'name'"`
}

//TableName table name
func (m *TestResult) TableName() string {
	return "test_result"
}
