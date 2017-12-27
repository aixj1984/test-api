package models

type CustomerCourse struct {
	Id         int
	CustomerId int  `xorm:"int 'customer_id'"`
	CourseId   int  `xorm:"int 'course_id'"`
	IsDisplay  int8 `xorm:"int 'is_display'"`
}

type CustomerCourseDetail struct {
	Course    `xorm:"extends"`
	IsDisplay int8 `xorm:"int 'is_display'"`
}

func (m *CustomerCourse) TableName() string {
	return "customer_course"
}

func (u *CustomerCourse) TableUnique() [][]string {
	return [][]string{
		[]string{"CustomerId", "CourseId"},
	}
}
