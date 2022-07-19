package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserViewLogic {
	return &UserViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserViewLogic) UserView(req *types.UserViewReq) (resp *types.UserViewResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

	return &types.UserViewResp{
		Code: 200,
		Msg:  "",
		Data: types.UserViewData{
			Id:       user.Id,
			StatusID: user.StatusId,
			Sort:     user.Sort,
			UserNo:   user.UserNo,
			Mobile:   user.Mobile,
			RealName: user.RealName,
			Remark:   user.Remark.String,
		},
	}, nil
}
