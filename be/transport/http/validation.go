package http

import (
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/transport/http/m_token"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constant.DEAL_SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = constant.INVALID_PARAMS
		} else {
			_, err := m_token.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constant.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = constant.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != constant.DEAL_SUCCESS {
			SendResponse(c, http.StatusUnauthorized, code, constant.GetMsg(code), data)
			c.Abort()
			return
		}

		c.Next()
	}
}
