package providers

var (
	//Account Account Provider
	Account  IAccountProvider
	Question IQuestionProvider
	Test     ITestProvider

	TestQuestion ITestQuestionProvider

	CustomerCourse ICustomerCourseProvider

	TestResult ITestResultProvider

	CollectQuestion ICollectQuestionProvider
)

//Init 初始化服务
func init() {
	Account = &AccountProvider{}

	Question = &QuestionProvider{}

	Test = &TestProvider{}

	TestQuestion = &TestQuestionProvider{}

	CustomerCourse = &CustomerCourseProvider{}

	TestResult = &TestResultProvider{}

	CollectQuestion = &CollectQuestionProvider{}

	println("初始化providers")
}
