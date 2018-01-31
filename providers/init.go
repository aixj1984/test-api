package providers

var (
	//Account Account Provider
	Account  IAccountProvider
	Wx       IWxProvider
	Question IQuestionProvider
	Test     ITestProvider

	TestQuestion ITestQuestionProvider

	CustomerCourse ICustomerCourseProvider

	TestResult ITestResultProvider

	Article IArticleProvider

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

	Article = &ArticleProvider{}

	CollectQuestion = &CollectQuestionProvider{}

	Wx = &WxProvider{}

	println("初始化providers")
}
