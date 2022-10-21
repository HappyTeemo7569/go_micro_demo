package logic

import (
	"context"

	"go-zero-demo/game/internal/svc"
	"go-zero-demo/game/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GameLogic {
	return GameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameLogic) Game(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
