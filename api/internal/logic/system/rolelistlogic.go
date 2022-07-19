package system

import (
	"context"
	"zero-template-react/api/internal/common/errorx"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleListLogic {
	return RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req types.RoleListReq) (resp *types.RoleListResp, err error) {
	count, _ := l.svcCtx.RoleModel.Count(l.ctx)
	roleList, err := l.svcCtx.RoleModel.FindAll(l.ctx, req.PageCurrent, req.PageSize, req.RoleName)

	if err != nil {
		return nil, errorx.NewDefaultError("查询角色列表异常")
	}
	var list []types.RoleList

	for _, role := range *roleList {
		list = append(list, types.RoleList{
			Id:       role.Id,
			StatusID: role.StatusId,
			Sort:     role.Sort,
			RoleName: role.RoleName,
			Remark:   role.Remark,
		})
	}

	return &types.RoleListResp{
		Code: 200,
		Msg:  "",
		Data: types.RoleListData{
			List:        list,
			TotalCount:  count,
			TotalPage:   0,
			PageCurrent: req.PageCurrent,
			PageSize:    req.PageSize,
		},
	}, nil
}
