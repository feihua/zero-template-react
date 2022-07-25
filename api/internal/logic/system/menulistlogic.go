package system

import (
	"context"
	"encoding/json"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuListLogic {
	return MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req types.MenuListReq) (resp *types.MenuListResp, err error) {
	menuList, err := l.svcCtx.MenuModel.FindAll(l.ctx, req.MenuName)

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询菜单列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询菜单失败")
	}

	var list []types.MenuList
	for _, menu := range *menuList {
		list = append(list, types.MenuList{
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
			Icon:     menu.MenuIcon.String,
		})
	}

	return &types.MenuListResp{
		Code: 200,
		Msg:  "",
		Data: list,
	}, nil
}
