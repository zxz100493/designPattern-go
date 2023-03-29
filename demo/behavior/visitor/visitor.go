package visitor

import "fmt"

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (e *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (i *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(i)
}

type ServiceRequestVisitor struct {
}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("EnterpriseCustomer %s request service\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("individualCustomer %s request service\n", c.name)
	}
}

type AnalysisVisitor struct {
}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("EnterpriseCustomer %s request service\n", c.name)
	}
}
