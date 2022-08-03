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

type RoleUserSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUserSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleUserSaveLogic {
	return RoleUserSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUserSaveLogic) RoleUserSave(req types.RoleUserSaveReq) (resp *types.RoleUserSaveResp, err error) {

	if req.UserId == 1 {
		return nil, errorx.NewDefaultError("不能修改超级管理员的角色")
	}

	_ = l.svcCtx.RoleUserModel.DeleteByUserId(l.ctx, req.UserId)

	for _, id := range req.RoleIds {
		_, _ = l.svcCtx.RoleUserModel.Insert(l.ctx, &model.SysRoleUser{
			GmtCreate:   time.Now(),
			GmtModified: time.Now(),
			StatusId:    1,
			Sort:        1,
			RoleId:      id,
			UserId:      req.UserId,
		})
	}

	return &types.RoleUserSaveResp{
		Code: 200,
		Msg:  "",
	}, nil
}
