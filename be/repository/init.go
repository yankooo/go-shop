package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yankooo/school-eco/be/model"
	"sync"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _G_gormDb *dbEngine

type dbEngine struct {
	m  sync.Mutex
	db *gorm.DB
}

// 初始化数据库连接池
func InitDbEngine(mysqlConfig *model.Mysql) (err error) {
	_G_gormDb = &dbEngine{}
	_G_gormDb.db, err = gorm.Open(mysqlConfig.Driver, mysqlConfig.Addr)
	if err != nil {
		return err
	}

	//设置数据库连接池参数
	//_G_gormDb.db.LogMode(false)                                      // 关闭错误日志
	_G_gormDb.db.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConn)       //设置数据库连接池最大连接数
	_G_gormDb.db.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConnection) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	_G_gormDb.db.DB().SetConnMaxLifetime(time.Duration(mysqlConfig.ConnMaxLifetime) * time.Second)
	return
}

func GormDb() *dbEngine {
	return _G_gormDb
}
