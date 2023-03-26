package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBMouduleAPI(),
	}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AMoudleAPI
	b BMoudleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AMoudleAPI interface {
	TestA() string
}

type aModuleImpl struct {
}

func NewAModuleAPI() AMoudleAPI {
	return &aModuleImpl{}
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func NewBMouduleAPI() BMoudleAPI {
	return &bModuleImpl{}
}

type BMoudleAPI interface {
	TestB() string
}

type bModuleImpl struct {
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
