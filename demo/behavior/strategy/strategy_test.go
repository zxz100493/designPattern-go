package strategy

func ExamplePayByCash() {
	payment := NewPayment("ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to ada vy cash
}

func ExamplePayByBank() {
	payment := NewPayment("ada", "", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $123 to ada by back account
}
