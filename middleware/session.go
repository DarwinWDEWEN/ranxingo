package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var store = cookie.NewStore([]byte("secret"))

func SessionInit(c *gin.Engine) {
	c.Use(sessions.Sessions("rx-session", store))
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := session.Get("session_id")
		if sessionID == nil {
			session.Set("session_id", "1234567890")
			_ = session.Save()
		}
		c.Next()
	}
}
