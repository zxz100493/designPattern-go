package chain

import "fmt"

type Manager interface {
	HaveRight(money int) bool
	HandleFeeRequest(name string, money int) bool
}

type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int) bool {
	if r.Manager.HaveRight(money) {
		return r.Manager.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int) bool {
	return true
}

type ProjectManager struct{}

func NewProjectManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &ProjectManager{},
	}
}

func (r *ProjectManager) HandleFeeRequest(name string, money int) bool {
	if name == "bob" {
		fmt.Printf("project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("project manager do not permit %s %d fee request\n", name, money)
	return false
}

func (r *ProjectManager) HaveRight(money int) bool {
	return money < 500
}

type DepManager struct{}

func NewDepManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &DepManager{},
	}
}

func (r *DepManager) HandleFeeRequest(name string, money int) bool {
	if name == "tom" {
		fmt.Printf("project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("project manager do not permit %s %d fee request\n", name, money)
	return false
}

func (r *DepManager) HaveRight(money int) bool {
	return money < 5000
}

type GeneralManager struct{}

func NewGeneralManagerChain() *RequestChain {
	return &RequestChain{
		Manager: &GeneralManager{},
	}
}

func (r *GeneralManager) HandleFeeRequest(name string, money int) bool {
	if name == "ada" {
		fmt.Printf("project manager permit %s %d fee request\n", name, money)
		return true
	}
	fmt.Printf("project manager do not permit %s %d fee request\n", name, money)
	return false
}

func (r *GeneralManager) HaveRight(money int) bool {
	return true
}
