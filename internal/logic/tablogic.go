package logic

import (
	"context"

	"e_tab_be/internal/svc"
	"e_tab_be/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TabLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTabLogic(ctx context.Context, svcCtx *svc.ServiceContext) TabLogic {
	return TabLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TabLogic) SaveTabs(req types.Request) (*types.Response, error) {

	err := l.svcCtx.SaveTabs(l.ctx, req.Tabs)
	if err != nil {
		return nil, err
	}

	return &types.Response{}, nil
}
