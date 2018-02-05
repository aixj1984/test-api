package payloads

type SaveCustomerCoursesSetting struct {
	DefalutCourses string `validate:"required"`
	//CustomerId     int    `validate:"required"`
}

type PurchaseCourses struct {
	PurchaseCourses string `validate:"required"`
}
