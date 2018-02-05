package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	ORDERTYPE = "xg"
)

func GetOrderNo() string {
	timenow := time.Now().Format("20060102150304")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(100)
	order_no := ORDERTYPE + timenow + strconv.Itoa(random)
	return order_no
}
