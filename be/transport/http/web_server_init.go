package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yankooo/school-eco/be/config"
	"log"
)

type runner struct {
	listenPort string
	listenWay  string
	e          *gin.Engine
	bs         *bookSeller
}

var webServe *runner

// 初始化web服务
func InitWebServer() *runner {
	//gin.SetMode(gin.ReleaseMode)
	webServe = &runner{
		listenPort: config.GlobalConf().Port,
		listenWay:  config.GlobalConf().ListenWay,
		e:          gin.New(),
		bs:         &bookSeller{},
	}
	// 1. 注册路由
	webServe.registerRouter()

	// 2. 添加中间件
	webServe.addMiddleWare()

	return webServe
}

// 挂起web服务
func (r *runner) Run() {
	if r.listenWay == "http" {
		if err := r.e.Run(":" + r.listenPort); err != nil {
			// 挂起失败直接退出
			log.Fatalf("can't start http server with err : %+v", err)
		}
	} else if r.listenWay == "https" {
		// TODO
	}
}
