package system

import (
	"context"
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

	//所有角色
	allRoleList, _ := l.svcCtx.RoleModel.FindAll(l.ctx, 1, 1000, "", "")

	var roleList []types.RoleUserList
	var allUserRoleIds []int64

	for _, role := range *allRoleList {
		roleList = append(roleList, types.RoleUserList{
			Id:         role.Id,
			StatusID:   role.StatusId,
			Sort:       role.Sort,
			RoleName:   role.RoleName,
			Remark:     role.Remark,
			CreateTime: role.GmtCreate.Format("2006-01-02 15:04:05"),
			UpdateTime: role.GmtModified.Format("2006-01-02 15:04:05"),
		})
		allUserRoleIds = append(allUserRoleIds, role.Id)
	}

	if req.UserId != 1 {
		allUserRoleIds, _ = l.svcCtx.RoleUserModel.FindAllRoleIdsByUserId(l.ctx, req.UserId)
	}

	return &types.RoleUserListResp{
		Code: 200,
		Msg:  "",
		Data: types.RoleUserListData{
			AllRoles:       roleList,
			AllUserRoleIds: allUserRoleIds,
		},
	}, nil
}
