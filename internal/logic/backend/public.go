package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/dao"
	"go-service/internal/service/backend"
	"time"
)

type sPublic struct {
}

func init() {
	backend.RegisterPublic(&sPublic{})
}

//LLogin
/**
 * @desc：登录操作
 * @param ctx
 * @param req
 * @return string
 * @return error
 * @author : Carson
 */
func (s *sPublic) LLogin(ctx context.Context, req *backendRoute.PublicReq) (string, error) {
	//传参
	username := req.Username
	password := req.Password
	//验证
	admin, err := dao.GetAdmin(username, password)
	if err != nil {
		return "", err
	}

	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	//后台生成token
	claims := jwt.MapClaims{
		"app":      "jh",
		"id":       admin.Id,
		"site_id":  admin.SiteId,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
