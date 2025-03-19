package middleware

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/DarwinWDEWEN/ranxingo/config"
	"github.com/DarwinWDEWEN/ranxingo/types/constant"
	"github.com/DarwinWDEWEN/ranxingo/utils"
	"github.com/DarwinWDEWEN/ranxingo/utils/resp"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
)

type NeedLoginFunc func(c *gin.Context) bool

// 用户授权验证
func AuthorizeMiddleware(client *redis.Client, needLogin NeedLoginFunc, authConfig config.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString = c.GetHeader(constant.UserAuthHeader)
		if tokenString == "" {
			if needLogin(c) {
				resp.ERROR(c, "You should put Authorization in request headers")
				c.Abort()
				return
			} else { // 直接放行
				c.Next()
				return
			}
		}
		if !strings.HasPrefix(tokenString, "Bearer") {
			resp.ERROR(c, "You should put Bearer in Authorization")
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok && needLogin(c) {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(authConfig.SecretKey), nil

		})

		if err != nil && needLogin(c) {
			resp.NotAuth(c, fmt.Sprintf("Error with parse auth token: %v", err))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid && needLogin(c) {
			resp.NotAuth(c, "Token is invalid")
			c.Abort()
			return
		}

		expr := utils.IntValue(utils.InterfaceToString(claims["expired"]), 0)
		if expr > 0 && int64(expr) < time.Now().Unix() && needLogin(c) {
			resp.NotAuth(c, "Token is expired")
			c.Abort()
			return
		}

		key := fmt.Sprintf("%s/%v", authConfig.SessionName, claims["user_id"])
		if _, err := client.Get(context.Background(), key).Result(); err != nil && needLogin(c) {
			resp.NotAuth(c, "Token is not found in redis")
			c.Abort()
			return
		}
		c.Set(constant.LoginUserID, claims["user_id"])
	}
}
