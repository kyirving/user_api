// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/kyirving/common_proto/user/userclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	res, err := l.svcCtx.UserRpcClient.Login(l.ctx, &userclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return
	}

	resp = &types.LoginResp{}
	resp.Code = 200
	resp.Msg = "success"
	resp.Data = &types.LoginData{
		Token:         res.Token,
		Expire:        res.Expire,
		RefreshToken:  res.RefreshToken,
		RefreshExpire: res.RefreshExpire,
	}
	return resp, nil
}
