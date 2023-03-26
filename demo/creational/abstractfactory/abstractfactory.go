package abstractfactory

import "fmt"

//OrderMainDAO 为订单主记录
type OrderMainDao interface {
	SaveOrderMain()
}

//OrderDetailDAO 为订单详情纪录
type OrderDetailDao interface {
	SaveOrderDetail()
}

//DAOFactory DAO 抽象模式工厂接口
type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

//RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDao struct{}

func (*RDBMainDao) SaveOrderMain() {
	fmt.Print("Rdb main save \n")
}

//RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDao struct{}

func (*RDBDetailDao) SaveOrderDetail() {
	fmt.Print("Rdb detail save \n")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (*RDBDAOFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

type XMLMainDao struct{}
type XMLDetailDao struct{}

func (*XMLMainDao) SaveOrderMain() {
	fmt.Print("xml main save \n")
}

func (*XMLDetailDao) SaveOrderDetail() {
	fmt.Print("xml detail save \n")
}

type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (*XMLDAOFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}
