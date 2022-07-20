package system

import (
	"context"
	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserListLogic {
	return UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req types.UserListReq) (resp *types.UserListResp, err error) {

	userList, _ := l.svcCtx.UserModel.FindAll(l.ctx, req.Current, req.PageSize, req.Mobile)
	count, _ := l.svcCtx.UserModel.Count(l.ctx, req.Mobile)

	var list []types.UserList

	for _, user := range *userList {
		list = append(list, types.UserList{
			Id:       user.Id,
			StatusID: user.StatusId,
			Sort:     user.Sort,
			UserNo:   user.UserNo,
			Mobile:   user.Mobile,
			RealName: user.RealName,
			Remark:   user.Remark.String,
		})
	}

	return &types.UserListResp{
		Code:     200,
		Msg:      "",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Success:  true,
		Total:    count,
	}, nil
}
