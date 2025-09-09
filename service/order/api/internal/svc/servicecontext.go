package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/hd2yao/micro-mall/service/order/api/internal/config"
	"github.com/hd2yao/micro-mall/service/order/rpc/orderclient"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
