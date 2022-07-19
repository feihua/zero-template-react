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

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleUpdateLogic {
	return RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req types.RoleUpdateReq) (resp *types.RoleUpdateResp, err error) {
	err = l.svcCtx.RoleModel.Update(l.ctx, &model.SysRole{
		Id:          req.Id,
		GmtModified: time.Now(),
		StatusId:    req.StatusID,
		Sort:        req.Sort,
		RoleName:    req.RoleName,
		Remark:      req.Remark,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf(err.Error())
		return nil, errorx.NewDefaultError("更新角色失败")
	}

	return &types.RoleUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
