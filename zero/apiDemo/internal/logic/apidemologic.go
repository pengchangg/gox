package logic

import (
	"context"

	"gox/zero/apiDemo/internal/svc"
	"gox/zero/apiDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiDemoLogic {
	return &ApiDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiDemoLogic) ApiDemo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
