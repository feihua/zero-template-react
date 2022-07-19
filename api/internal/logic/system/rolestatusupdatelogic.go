package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleStatusUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleStatusUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleStatusUpdateLogic {
	return &RoleStatusUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleStatusUpdateLogic) RoleStatusUpdate(req *types.RoleStatusUpdateReq) (resp *types.RoleStatusUpdateResp, err error) {
	err = l.svcCtx.RoleModel.UpdateRoleStatus(l.ctx, req.Id, req.StatusID)

	if err != nil {
		logx.WithContext(l.ctx).Errorf(err.Error())
		return nil, errorx.NewDefaultError("更新角色状态失败")
	}

	return &types.RoleStatusUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
