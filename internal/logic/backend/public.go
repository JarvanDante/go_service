package backend

import (
	"context"
	"fmt"
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

	// true

	username := req.Username
	password := req.Password
	admin, err := dao.GetAdmin(username, password)
	if err != nil {
		return "", err
	}

	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	//后台token生成结构
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
	fmt.Println(tokenString)
	return tokenString, nil

	//eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vaS5tZF9zZXJ2aWNlLmNvbS9hcGkvYmFja2VuZC9sb2dpbiIsImlhdCI6MTc2MjQ0MDIxOCwiZXhwIjoxNzYzMDQ1MDE4LCJuYmYiOjE3NjI0NDAyMTgsImp0aSI6Ilc3ZktURWY1R203ZXp3THIiLCJzdWIiOjI3LCJwcnYiOiJkZjg4M2RiOTdiZDA1ZWY4ZmY4NTA4MmQ2ODZjNDVlODMyZTU5M2E5In0.LEHivbtYdAmWjl3kNmlu-qdJqv-DI8nNrb1fJKemdjs
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZWxldGVfYXQiOjAsInBhc3N3b3JkIjoicGcxMjM0NTYiLCJzaXRlX2lkIjoxLCJzdGF0dXMiOjEsInVzZXJuYW1lIjoibWljaGFlbCJ9.uaY-u7DstYiR9490Cs1D2sh147EhK1dE8LRJIIb-r-s
}
