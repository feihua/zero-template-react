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

type UserSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserSaveLogic {
	return UserSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSaveLogic) UserSave(req types.UserSaveReq) (resp *types.UserSaveResp, err error) {

	mobile, _ := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)

	if mobile != nil {
		return nil, errorx.NewDefaultError("手机号码已注册")
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.SysUser{
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
		StatusId:    1,
		Sort:        1,
		UserNo:      11,
		Mobile:      req.Mobile,
		RealName:    req.RealName,
		Remark: sql.NullString{
			String: req.Remark,
			Valid:  true,
		},
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加用户异常")
	}

	return &types.UserSaveResp{
		Code: 200,
		Msg:  "",
	}, nil
}
