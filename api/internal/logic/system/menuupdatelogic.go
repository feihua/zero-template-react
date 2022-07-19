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

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuUpdateLogic {
	return MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req types.MenuUpdateReq) (resp *types.MenuUpdateResp, err error) {
	err = l.svcCtx.MenuModel.Update(l.ctx, &model.SysMenu{
		Id:          req.Id,
		GmtModified: time.Now(),
		StatusId:    req.StatusId,
		Sort:        req.Sort,
		ParentId:    req.ParentId,
		MenuName:    req.MenuName,
		MenuUrl:     req.MenuUrl,
		ApiUrl:      req.ApiUrl,
		MenuIcon:    sql.NullString{},
		Remark: sql.NullString{
			String: req.Remark,
			Valid:  true,
		},
		MenuType: req.MenuType,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("更新菜单失败")
	}

	return &types.MenuUpdateResp{
		Code: 200,
		Msg:  "",
	}, nil
}
