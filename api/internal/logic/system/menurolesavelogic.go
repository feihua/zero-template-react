package system

import (
	"context"
	"time"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/model"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRoleSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuRoleSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuRoleSaveLogic {
	return MenuRoleSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuRoleSaveLogic) MenuRoleSave(req types.MenuRoleSaveReq) (resp *types.MenuRoleSaveResp, err error) {

	if req.RoleId == 1 {
		return nil, errorx.NewDefaultError("不能修改超级管理员的权限")
	}

	_ = l.svcCtx.MenuRoleModel.DeleteByRoleId(l.ctx, req.RoleId)

	ids := req.MenuIds
	for _, id := range ids {
		_, _ = l.svcCtx.MenuRoleModel.Insert(l.ctx, &model.SysMenuRole{
			GmtCreate:   time.Now(),
			GmtModified: time.Now(),
			StatusId:    1,
			Sort:        1,
			MenuId:      id,
			RoleId:      req.RoleId,
		})
	}

	return &types.MenuRoleSaveResp{
		Code: 200,
		Msg:  "",
	}, nil
}
