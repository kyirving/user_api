## 用户服务

### 初始化项目
```bash
mkdir user_api && cd user_api

goctl api go -api ./DSL/user.api -dir .

# 获取公共协议
go get github.com/kyirving/common_proto@v1.0.0

# 清理依赖
go mod tidy
```

### 配置微服务RPC客户端
* `./etc/user.yaml`
```bash
# 新增用户RPC服务配置项
UserRpc:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: user.rpc # RPC 服务名称 务必和 user_rpc/etc/user.yaml 中的 Name 一致
  Timeout: 2000 # 2s
  KeepAliveTime: 20s
```
* `./internal/config/config.go` 配置结构体
```bash
type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf 
}
```

* `internal/svc/servicecontext.go` 服务上下文结构体
```go
package svc

import (
    "user-api/internal/config"
    	"github.com/kyirving/common_proto/user/userclient" // 引入 userclient 包 
    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {  
    Config      config.Config
    UserRpcClient userclient.User // RPC 客户端实例
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)), // 创建客户端
    }
}
```

* `internal/logic/loginlogic.go` 登录逻辑实现

```go
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	res, err := l.svcCtx.UserRpcClient.Login(l.ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}

	resp = &types.LoginResp{
		Token: res.Token,
	}

	return
}
````


