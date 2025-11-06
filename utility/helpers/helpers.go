package helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
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
