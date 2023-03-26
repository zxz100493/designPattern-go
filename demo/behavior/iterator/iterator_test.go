package iterator

func ExampleIterator() {
	var agreegate Agreegate
	agreegate = NewNumbers(1, 10)

	IteratorPrint(agreegate.Iterator())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
