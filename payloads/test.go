package payloads

type SaveTest struct {
	CourseId int `validate:"required"`
	TestId   int `validate:"required"`
	TestType int
	Abstract string `validate:"required"`
	Title    string `validate:"required"`
	Sources  string `validate:"required"`
}

type SaveTestStatus struct {
	CourseId int `validate:"required"`
	TestId   int `validate:"required"`
	Status   int `validate:"required"`
}

type InsertTest struct {
	CourseId int `validate:"required"`
	TestType int
	Abstract string `validate:"required"`
	Title    string `validate:"required"`
	Sources  string `validate:"required"`
}
type SaveTestQuestionStatus struct {
	TestId     int `validate:"required"`
	QuestionId int `validate:"required"`
	Status     int
}

type SaveTestResult struct {
	CourseId    int `validate:"required"`
	TestId      int `validate:"required"`
	QuestionNum int `validate:"required"`
	TestSec     int `validate:"required"`
	RightNum    int
}
