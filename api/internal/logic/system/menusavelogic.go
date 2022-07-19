package system

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/model"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuSaveLogic {
	return MenuSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuSaveLogic) MenuSave(req types.MenuSaveReq) (resp *types.MenuSaveResp, err error) {

	_, err = l.svcCtx.MenuModel.Insert(l.ctx, &model.SysMenu{
		GmtCreate:   time.Now(),
		GmtModified: time.Now(),
		StatusId:    1,
		Sort:        req.Sort,
		ParentId:    req.ParentId,
		MenuName:    req.MenuName,
		MenuUrl:     req.MenuUrl,
		ApiUrl:      req.MenuUrl,
		MenuIcon:    sql.NullString{},
		Remark: sql.NullString{
			String: req.Remark,
			Valid:  true,
		},
		MenuType: req.MenuType,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,添加菜单异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("添加菜单失败")
	}

	return &types.MenuSaveResp{
		Code: 200,
		Msg:  "",
	}, nil
}
