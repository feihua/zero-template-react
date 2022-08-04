package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserUpdatePasswordLogic {
	return UserUpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdatePasswordLogic) UserUpdatePassword(req types.UserUpdatePasswordReq) (resp *types.UserUpdatePasswordResp, err error) {
	err = l.svcCtx.UserModel.UpdatePassword(l.ctx, req.Id, req.MobilePsw)

	if err != nil {
		return nil, errorx.NewDefaultError("重置用户密码异常")
	}

	logx.WithContext(l.ctx).Error("更新密码成功")
	return &types.UserUpdatePasswordResp{
		Code: 200,
		Msg:  "",
	}, nil
}
