package jwt

import (
	"Campusforum/app/user/models"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

type JwtUser struct {
	Uid      int    `json:"uid"`      // ID
	Nickname string `json:"nickname"` // 昵称
}

const (
	//ContextKeyUserObj 存登录的用户
	ContextKeyUserObj = "authedUserId"
	//LoginCode 验证码
	LoginCode        = "login_code:"
	DefaultAvatarURL = "resource/favicon.ico"
)

type userStdClaims struct {
	JwtUser
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间设置
const TokenExpireDuration = time.Hour * 336 //两周

// Secret token秘钥
var Secret = []byte("dijiexiaV3")
var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(user *models.User) (string, error) {
	var jwtUser = JwtUser{
		Uid:      user.Id,
		Nickname: user.Username,
	}
	c := userStdClaims{
		jwtUser, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "backstage",                                // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*JwtUser, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtUser, err
}

// GetUserId 返回id
func GetUserId(c *gin.Context) (int, error) {
	u, exist := c.Get(ContextKeyUserObj)
	if !exist {
		return -1, errors.New("can't get user id")
	}
	user, ok := u.(*JwtUser)
	if ok {
		return user.Uid, nil
	}
	return -1, errors.New("can't convert to id struct")
}

// GetUserNickname 返回用户名；
func GetUserNickname(c *gin.Context) (string, error) {
	// 从上下文中获取JWT用户对象
	u, exist := c.Get(ContextKeyUserObj)
	if !exist {
		return "", errors.New("无法获取用户信息")
	}

	// 将获取到的用户对象转换为 *models.JwtUser 类型
	user, ok := u.(*JwtUser)
	if !ok {
		return "", errors.New("无法转换为用户对象")
	}

	// 返回JWT中的昵称
	return user.Nickname, nil
}
