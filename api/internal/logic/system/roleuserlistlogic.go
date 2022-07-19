package system

import (
	"context"
	"zero-template-react/api/model"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleUserListLogic {
	return RoleUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUserListLogic) RoleUserList(req types.RoleUserListReq) (resp *types.RoleUserListResp, err error) {

	var roleList *[]model.SysRole
	if req.UserId == 1 {
		roleList, _ = l.svcCtx.RoleModel.FindAll(l.ctx, 1, 1000, "")
	} else {
		roleList, _ = l.svcCtx.RoleModel.FindAllByUserId(l.ctx, req.UserId)
	}

	var list []types.RoleUserList

	for _, role := range *roleList {
		list = append(list, types.RoleUserList{
			Id:       role.Id,
			StatusID: role.StatusId,
			Sort:     role.Sort,
			RoleName: role.RoleName,
			Remark:   role.Remark,
		})
	}

	return &types.RoleUserListResp{
		Code: 200,
		Msg:  "",
		Data: types.RoleUserListData{
			List: list,
		},
	}, nil
}
