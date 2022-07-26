package system

import (
	"context"
	"strings"
	"time"
	"zero-template-react/api/internal/common/errorx"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"zero-template-react/api/internal/svc"
	"zero-template-react/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	if len(strings.TrimSpace(req.Mobile)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)

	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", req.Mobile, err.Error())
		return nil, errorx.NewDefaultError("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", req.Mobile, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	if userInfo.Password != req.Password {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", req.Password)
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret

	jwtToken, _ := l.getJwtToken(accessSecret, now, accessExpire, userInfo.Id)

	return &types.UserLoginResp{
		Code:   200,
		Msg:    "",
		Status: "ok",
		Data: types.UserLoginData{
			UserNo: userInfo.UserNo,
			Mobile: userInfo.Mobile,
			Token:  jwtToken,
		},
	}, nil
}

func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
