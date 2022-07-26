package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {

	for _, id := range req.Ids {
		err = l.svcCtx.UserModel.Delete(l.ctx, id)

		if err != nil {
			logx.WithContext(l.ctx).Errorf("根据userId: %d,删除用户异常:%s", id, err.Error())
			return nil, errorx.NewDefaultError("删除用户异常")
		}
	}

	return &types.UserDeleteResp{
		Code: 200,
		Msg:  "",
	}, nil

}
