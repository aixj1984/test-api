package models

/*
1	轮机基础
2	轮机管理
3	船舶辅机与电气
4	船舶动力装置
5	避碰与信号
6	船舶管理
7	航道与引航
8	船舶操纵
9	机舱管理
10	主推进动力装置
11	船舶驾驶与管理
*/

var CourseCollectMap = map[string]string{
	"1":  "collect_question_1",
	"2":  "collect_question_2",
	"3":  "collect_question_3",
	"4":  "collect_question_4",
	"5":  "collect_question_5",
	"6":  "collect_question_6",
	"7":  "collect_question_7",
	"8":  "collect_question_8",
	"9":  "collect_question_9",
	"10": "collect_question_10",
	"11": "collect_question_11",
}

//Account account model
type CollectQuestion struct {
	Id         int
	CustomerId int    `xorm:"int 'customer_id'"`
	QuestionId int    `xorm:"int 'question_id'"`
	AddTime    string `xorm:"varchar(60)  'add_time'"`
}

type CollectQuestion1 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion1) TableName() string {
	return "collect_question_1"
}

type CollectQuestion2 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion2) TableName() string {
	return "collect_question_2"
}

type CollectQuestion3 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion3) TableName() string {
	return "collect_question_3"
}

type CollectQuestion4 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion4) TableName() string {
	return "collect_question_4"
}

type CollectQuestion5 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion5) TableName() string {
	return "collect_question_5"
}

type CollectQuestion6 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion6) TableName() string {
	return "collect_question_6"
}

type CollectQuestion7 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion7) TableName() string {
	return "collect_question_7"
}

type CollectQuestion8 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion8) TableName() string {
	return "collect_question_8"
}

type CollectQuestion9 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion9) TableName() string {
	return "collect_question_9"
}

type CollectQuestion10 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion10) TableName() string {
	return "collect_question_10"
}

type CollectQuestion11 struct {
	CollectQuestion `xorm:"extends"`
}

func (m *CollectQuestion11) TableName() string {
	return "collect_question_11"
}
