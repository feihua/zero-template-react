package system

import (
	"context"
	"database/sql"
	"time"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/model"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserUpdateLogic {
	return UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {

	err = l.svcCtx.UserModel.Update(l.ctx, &model.SysUser{
		Id:          req.Id,
		GmtModified: time.Now(),
		StatusId:    req.StatusID,
		Sort:        req.Sort,
		Mobile:      req.Mobile,
		RealName:    req.RealName,
		Remark:      sql.NullString{String: req.Remark, Valid: true},
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新用户异常")
	}

	return &types.UserUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
