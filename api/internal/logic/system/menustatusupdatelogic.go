package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuStatusUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuStatusUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuStatusUpdateLogic {
	return &MenuStatusUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuStatusUpdateLogic) MenuStatusUpdate(req *types.MenuStatusUpdateReq) (resp *types.MenuStatusUpdateResp, err error) {
	err = l.svcCtx.MenuModel.UpdateMenuStatus(l.ctx, req.Id, req.StatusID)

	if err != nil {
		return nil, errorx.NewDefaultError("更新菜单状态失败")
	}

	return &types.MenuStatusUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
