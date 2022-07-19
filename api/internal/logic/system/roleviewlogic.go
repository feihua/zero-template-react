package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleViewLogic {
	return &RoleViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleViewLogic) RoleView(req *types.RoleViewReq) (resp *types.RoleViewResp, err error) {
	sysRole, err := l.svcCtx.RoleModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, errorx.NewDefaultError("角色不存在")
	}

	return &types.RoleViewResp{
		Code: 200,
		Msg:  "",
		Data: types.RoleViewData{
			Id:       sysRole.Id,
			StatusID: sysRole.StatusId,
			Sort:     sysRole.Sort,
			RoleName: sysRole.RoleName,
			Remark:   sysRole.Remark,
		},
	}, nil
}
