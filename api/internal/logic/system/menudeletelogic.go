package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuDeleteLogic {
	return MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req types.MenuDeleteReq) (resp *types.MenuDeleteResp, err error) {

	for _, id := range req.Ids {
		err = l.svcCtx.MenuModel.Delete(l.ctx, id)

		if err != nil {
			logx.WithContext(l.ctx).Errorf("根据menuId: %d,删除菜单异常:%s", id, err.Error())
			return nil, errorx.NewDefaultError("删除菜单异常")
		}
	}

	return &types.MenuDeleteResp{
		Code: 200,
		Msg:  "",
	}, nil
}
