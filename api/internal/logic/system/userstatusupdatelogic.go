package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserStatusUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserStatusUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserStatusUpdateLogic {
	return &UserStatusUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserStatusUpdateLogic) UserStatusUpdate(req *types.UserStatusUpdateReq) (resp *types.UserStatusUpdateResp, err error) {
	err = l.svcCtx.UserModel.UpdateUserStatus(l.ctx, req.Id, req.StatusID)

	if err != nil {
		return nil, errorx.NewDefaultError("更新用户状态异常")
	}

	return &types.UserStatusUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
