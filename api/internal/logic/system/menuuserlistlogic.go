package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"strings"
	"zero-template-react/api/internal/common/errorx"
	"zero-template-react/api/model"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUserListLogic {
	return &MenuUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUserListLogic) MenuUserList() (resp *types.MenuUserListResp, err error) {

	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Infof("用户不存在userId: %s", userId)
		return nil, errorx.NewDefaultError(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(userId, 10)))
	default:
		return nil, err
	}

	var menuList *[]model.SysMenu

	if userId == 1 {
		menuList, _ = l.svcCtx.MenuModel.FindAll(l.ctx, "")
	} else {
		menuList, err = l.svcCtx.MenuModel.FindAllByUserId(l.ctx, userId)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("用户: %s,还没有权限,请求联系管理员", strconv.FormatInt(userId, 10))
			return nil, errorx.NewDefaultError("还没有权限,请求联系管理员")
		}
	}

	var list []types.MenuUserList

	var apiUrls []string

	for _, menu := range *menuList {
		if menu.MenuType != 3 {
			list = append(list, types.MenuUserList{
				Id:       menu.Id,
				ParentId: menu.ParentId,
				Name:     menu.MenuName,
				Path:     menu.MenuUrl,
				ApiUrl:   menu.ApiUrl,
				MenuType: menu.MenuType,
			})
		}

		if len(strings.TrimSpace(menu.ApiUrl)) != 0 {
			apiUrls = append(apiUrls, menu.ApiUrl)
		}
	}

	//把能访问的url存在在redis，在middleware中检验
	userIdStr := fmt.Sprintf("react-vue:%s", strconv.FormatInt(userId, 10))
	err = l.svcCtx.Redis.Set(userIdStr, strings.Join(apiUrls, ","))

	if err != nil {
		logx.Errorf("设置用户：%s,权限到redis异常: %+v", userInfo.RealName, err)
	}

	return &types.MenuUserListResp{
		Code: 200,
		Msg:  "",
		Data: types.MenuUserData{
			SysMenu: list,
		},
	}, nil
}
