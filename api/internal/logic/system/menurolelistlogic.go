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

	findAll, _ := l.svcCtx.MenuModel.FindAll(l.ctx, "")

	var menuRoleList []types.MenuRoleList
	var menuIds []int64

	for _, menu := range *findAll {
		menuRoleList = append(menuRoleList, types.MenuRoleList{
			Id:       menu.Id,
			ParentID: menu.ParentId,
			Title:    menu.MenuName,
			Key:      strconv.FormatInt(menu.Id, 10),
		})

		menuIds = append(menuIds, menu.Id)
	}

	if req.RoleId != 1 {
		menuIds, _ = l.svcCtx.MenuRoleModel.FindAllByRoleId(l.ctx, req.RoleId)
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
