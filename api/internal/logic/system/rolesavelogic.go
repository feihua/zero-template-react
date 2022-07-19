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

type RoleSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleSaveLogic {
	return RoleSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleSaveLogic) RoleSave(req types.RoleSaveReq) (resp *types.RoleSaveResp, err error) {
	_, err = l.svcCtx.RoleModel.Insert(l.ctx, &model.SysRole{
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
		StatusId:    1,
		Sort:        1,
		RoleName:    req.RoleName,
		Remark:      req.Remark,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加角色失败")
	}

	return &types.RoleSaveResp{
		Code: 200,
		Msg:  "",
	}, nil
}
