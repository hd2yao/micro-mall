package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/hd2yao/micro-mall/service/order/model"
	"github.com/hd2yao/micro-mall/service/order/rpc/internal/config"
	"github.com/hd2yao/micro-mall/service/product/rpc/productclient"
	"github.com/hd2yao/micro-mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
