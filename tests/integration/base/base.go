package base

type TestSuite interface {
	BeforeSuite() error
	RunSuite()
	AfterSuite()
}
