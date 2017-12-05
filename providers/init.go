package providers

var (
	//Account Account Provider
	Account  IAccountProvider
	Question IQuestionProvider
	Test     ITestProvider

	TestQuestion ITestQuestionProvider

	CustomerCourse ICustomerCourseProvider
)

//Init 初始化服务
func init() {
	Account = &AccountProvider{}

	Question = &QuestionProvider{}

	Test = &TestProvider{}

	TestQuestion = &TestQuestionProvider{}

	CustomerCourse = &CustomerCourseProvider{}

	println("初始化providers")
}
