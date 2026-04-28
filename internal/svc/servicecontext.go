// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"user_api/internal/config"

	"github.com/kyirving/common_proto/user/userclient" // 引入 userclient 包
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userclient.User // 引入 userclient 包中的 User 类型
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)), // 创建客户端
	}
}
