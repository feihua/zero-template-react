package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleDeleteLogic {
	return RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req types.RoleDeleteReq) (*types.RoleDeleteResp, error) {

	err := l.svcCtx.RoleModel.Delete(l.ctx, req.Id)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据roleId: %d,删除角色异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除角色异常")
	}

	return &types.RoleDeleteResp{
		Code: 200,
		Msg:  "",
	}, nil
}
