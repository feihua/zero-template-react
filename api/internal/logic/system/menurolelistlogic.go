package system

import (
	"context"
	"strconv"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuRoleListLogic {
	return MenuRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuRoleListLogic) MenuRoleList(req types.MenuRoleListReq) (resp *types.MenuRoleListResp, err error) {

	var menuIds []int64
	if req.RoleId == 1 {
		menus, _ := l.svcCtx.MenuModel.FindAll(l.ctx, "")
		for _, menu := range *menus {
			menuIds = append(menuIds, menu.Id)
		}
	} else {
		menuIds, _ = l.svcCtx.MenuRoleModel.FindAllByRoleId(l.ctx, req.RoleId)
	}

	findAll, _ := l.svcCtx.MenuModel.FindAll(l.ctx, "")

	var menuRoleList []types.MenuRoleList

	for _, menu := range *findAll {
		menuRoleList = append(menuRoleList, types.MenuRoleList{
			Id:       menu.Id,
			ParentID: menu.ParentId,
			Title:    menu.MenuName,
			Key:      strconv.FormatInt(menu.Id, 10),
		})
	}

	return &types.MenuRoleListResp{
		Code: 200,
		Msg:  "",
		Data: types.MenuRoleListData{
			RoleMenus:    menuIds,
			MenuRoleList: menuRoleList,
		},
	}, nil
}
