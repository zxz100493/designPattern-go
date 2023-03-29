package visitor

func ExampleRequestVisitor() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A"))
	c.Add(NewEnterpriseCustomer("B"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// serving individual customer bob
}
