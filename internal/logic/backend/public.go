package backend

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt"
	"go-service/api/backendRoute"
	"go-service/internal/dao/jh_site/dao"
	"go-service/internal/service/backend"
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

	// true

	username := req.Username
	password := req.Password
	admin, err := dao.GetAdmin(username, password)
	if err != nil {
		return "", err
	}
	fmt.Println("+++++++++")
	fmt.Println(admin)
	fmt.Println("+++++++++")

	//ok := helpers.CompareBcrypt(admin, "123456")
	//fmt.Println(ok)
	//helpers.Bcrypt(password)
	//// 这里应该做你的账号密码校验
	//if req.Username != "admin" || req.Password != "123456" {
	//	return "", gerror.New("账号或密码错误")
	//}

	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	//后台token生成结构
	claims := jwt.MapClaims{
		"admin_id":       1,
		"admin_name":     username,
		"admin_password": password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
