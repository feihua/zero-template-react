package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuViewLogic {
	return &MenuViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuViewLogic) MenuView(req *types.MenuViewReq) (resp *types.MenuViewResp, err error) {
	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, errorx.NewDefaultError("菜单不存在")
	}

	return &types.MenuViewResp{
		Code: 200,
		Msg:  "",
		Data: types.MenuViewData{
			Id:       menu.Id,
			StatusID: menu.StatusId,
			Sort:     menu.Sort,
			ParentID: menu.ParentId,
			MenuName: menu.MenuName,
			Label:    menu.MenuName,
			MenuURL:  menu.MenuUrl,
			ApiUrl:   menu.ApiUrl,
			Remark:   menu.Remark.String,
			MenuType: menu.MenuType,
		},
	}, nil
}
