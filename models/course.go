package models

type Course struct {
	Id     int
	Name   string `xorm:"varchar(300) 'name'"`
	Price  int    `xorm:"int 'price'"`
	Status int8   `xorm:"int  'status'"`
}

func (m *Course) TableName() string {
	return "course"
}
