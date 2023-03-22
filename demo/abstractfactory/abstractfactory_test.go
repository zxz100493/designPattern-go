package abstractfactory

func getMainAndDetailDao(factory DaoFactory) {
	factory.CreateOrderDetailDao().SaveOrderDetail()
	factory.CreateOrderMainDao().SaveOrderMain()
}

func ExampleRdbFactory() {
	var factory DaoFactory
	factory = &RDBDAOFactory{}
	getMainAndDetailDao(factory)
}

func ExampleXmlFactory() {
	var factory DaoFactory
	factory = &XMLDAOFactory{}
	getMainAndDetailDao(factory)
}
