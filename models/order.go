package models

import (
	"time"
)

type Order struct {
	OrderNo string `orm:"column(order_no);pk;size(32)"`
	//	BusinessType  string    `orm:"column(business_type);size(16)`
	OpenId        string    `orm:"column(openid);size(100)"`
	TotalFee      int       `orm:"column(total_fee)"`
	FeeType       string    `orm:"column(fee_type);size(16)"`
	ClientIp      string    `orm:"column(client_ip);size(32)"`
	PrepayId      string    `orm:"column(prepay_id);size(100);unique"`
	Bank          string    `orm:"column(bank);size(16);null"`
	TransactionId string    `orm:"column(transaction_id);size(100);null"`
	Status        int       `orm:"column(status)"`
	Name          string    `orm:"column(name);size(128);null"`
	PayErrMsg     string    `orm:"column(pay_err_msg);size(256);null"`
	TradeType     string    `orm:"column(trade_type);size(16)"`
	WxMsg         string    `orm:"column(wx_msg);null"`
	CreateTime    time.Time `orm:"column(create_time);auto_now_add;type(datetime)"`
	PayTime       time.Time `orm:"column(pay_tiem);type(datetime);null"`
	UpdateTime    time.Time `orm:"column(update_time);auto_now;type(datetime)"`
}

func (m *Order) TableName() string {
	return "order"
}

type OrderCourse struct {
	Id       int    `orm:"pk;auto"`
	OrderNo  string `orm:"column(order_no);size(32)"`
	CourseId int    `orm:"column(course_id)"`
}

func (m *OrderCourse) TableName() string {
	return "order_course"
}
