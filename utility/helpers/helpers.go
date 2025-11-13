package helpers

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt"
	"github.com/thinkeridea/go-extend/exnet"
	"go-service/internal/dao/jh_site/model/entity"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

//Bcrypt
/**
 * @desc：Bcrypt 加密密码（类似 PHP 的 bcrypt），cost 默认 10，可根据需要调整，值越大越安全也越慢
 * @param value
 * @param cost
 * @return string
 * @return error
 * @author : Carson
 */
func Bcrypt(value string, cost ...int) (string, error) {
	c := bcrypt.DefaultCost
	if len(cost) > 0 {
		c = cost[0]
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(value), c)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//CompareBcrypt
/**
 * @desc：CompareBcrypt 验证密码（类似 Hash::check）
 * @param hashed
 * @param plain
 * @return bool
 * @author : Carson
 */
func CompareBcrypt(hashed, plain string) bool {
	//bcrypt, err := helpers.Bcrypt("123123")
	//	if err != nil {
	//		return "", err
	//	}
	//	fmt.Println(bcrypt)
	//	ok := helpers.CompareBcrypt(bcrypt, "123456")
	//	fmt.Println(ok)

	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

//GetEnv
/**
 * @desc：获取配置变量的值
 * @param key
 * @param defaultValue
 * @return string
 * @author : Carson
 */
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

//Print
/**
 * @desc：打印
 * @param str
 * @author : Carson
 */
func Print(str interface{}) {
	fmt.Println("*** 如下 ****")
	fmt.Println(str)
	fmt.Println("*** 如上 ****")
}

//SetToken
/**
 * @desc：生成token
 * @param ctx
 * @param admin
 * @return tokenString
 * @return err
 * @author : Carson
 */
func SetToken(ctx context.Context, admin *entity.Admin) (tokenString string, err error) {
	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()
	//后台生成token
	claims := jwt.MapClaims{
		"app":      "jh",
		"id":       admin.Id,
		"site_id":  admin.SiteId,
		"username": admin.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return
}

//ParseToken
/**
 * @desc：解析token
 * @param ctx
 * @param tokenString
 * @return jwt.MapClaims
 * @return error
 * @author : Carson
 */
func ParseToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	secret := g.Cfg().MustGet(ctx, "jwt.secret").String()

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方式是否一致
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, gerror.New("无效的签名方法")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// 获取 Claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, gerror.New("无效token")
}

//GetAdminInfoFromToken
/**
 * @desc：获取admin信息
 * @param r
 * @return map[string]interface{}
 * @return error
 * @author : Carson
 */
func GetAdminInfoFromToken(r *ghttp.Request) (map[string]interface{}, error) {
	tokenString := r.Header.Get("Authorization")
	// 如果是 "Bearer xxx" 格式，截取一下
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := ParseToken(r.Context(), tokenString)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":       claims["id"],
		"site_id":  claims["site_id"],
		"username": claims["username"],
	}, nil
}

//RemoteIp
/**
 * @desc：获取远程ip
 * @param r
 * @return res
 * @author : Carson
 */
func RemoteIp(r *http.Request) (res interface{}) {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}
