package system

import (
	"context"
	"encoding/json"
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

// UserUpdatePassword 更新用户密码
func (l *UserUpdatePasswordLogic) UserUpdatePassword(req types.UserUpdatePasswordReq) (resp *types.UserUpdatePasswordResp, err error) {

	//密码是否相等
	if req.MobilePsw != req.RePwd {
		return nil, errorx.NewDefaultError("密码不一致")
	}

	//如果不传userId,就拿token中的
	if req.Id == 0 {
		userId, _ := l.ctx.Value("userId").(json.Number).Int64()
		req.Id = userId
	}

	//判断用户是否存在
	_, err = l.svcCtx.UserModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

	//更新密码
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
