package http

import (
	"github.com/gin-gonic/gin"
)

//二手书交易平台接口
type BookSellerApi interface {
	RegisterAccount(c *gin.Context)
	Login(c *gin.Context)
	QuerySingleAccountInfo(c *gin.Context)
	ModifyAccountInfo(c *gin.Context)
}
