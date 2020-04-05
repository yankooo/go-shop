package main

import (
	"flag"
	"github.com/yankooo/school-eco/be/cache"
	"github.com/yankooo/school-eco/be/config"
	"github.com/yankooo/school-eco/be/logger"
	"github.com/yankooo/school-eco/be/repository"
	"github.com/yankooo/school-eco/be/transport/http"
	"log"
)

var filePath string

func init()  {
	flag.StringVar(&filePath, "-conf", "./book_seller_conf.json", "配置文件")
	flag.Parse()
}

func main() {
	var err error
	// 初始化配置
	if err = config.InitConfig(filePath); err != nil {
		log.Fatalf("init config fail with : %+v", err)
	}

	// 初始化组件
	if err = initComponent(); err != nil {
		log.Fatalf("init component err: %+v", err)
	}

	http.InitWebServer().Run()
}

func initComponent() (err error) {
	// 初始化日志
	if err = logger.InitLogger(); err != nil {
		return
	}

	// 初始化mysql
	if err = repository.InitDbEngine(config.GlobalConf().Mysql); err != nil {
		return
	}

	// init redis
	if err = cache.InitRedisPool(config.GlobalConf().Redis); err != nil {
		return
	}

	return
}