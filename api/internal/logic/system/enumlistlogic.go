package system

import (
	"context"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnumListLogic {
	return &EnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnumListLogic) EnumList(req *types.EnumReq) (resp *types.EnumResp, err error) {

	var list []types.EnumData
	list = append(list, types.EnumData{
		Code:  0,
		Color: "red",
		Name:  "NO",
		Desc:  "禁用",
	})

	list = append(list, types.EnumData{
		Code:  1,
		Color: "",
		Name:  "YES",
		Desc:  "",
	})
	return &types.EnumResp{
		Code: 200,
		Msg:  "正常",
		Data: list,
	}, nil
}
