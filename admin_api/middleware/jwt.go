package middleware

import (
	"backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
/**
 * Author: lingjianrui
 * Email:  toiklaun@gmail.com
 */
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		codeMap := map[int]string{
			50001: "无效token",
			50002: "token超出有效时间",
		}
		code = 20000
		token := c.GetHeader("token")
		if token == "" {
			code = 1
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = 50001
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 50002
			}
		}

		if code != 20000 {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  codeMap[code],
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
