package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/shopspring/decimal"
)

/****************配置******************/
type BookSellerConf struct {
	Port         string        `json:"port"`
	ListenWay    string        `json:"listen_way"`
	Pprof        *Pprof        `json:"pprof"`
	Redis        *Redis        `json:"redis"`
	Mysql        *Mysql        `json:"mysql"`
	LoggerConfig *LoggerConfig `json:"logger_config"`
	AppId        string        `json:"app_id"`
	Secret       string        `json:"secret"`
}

type Pprof struct {
	Addr string `json:"addr"`
}

// redis 配置项
type Redis struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	Timeout     int    `json:"timeout"`
	DB          int    `json:"db"`
	MaxIdle     int    `json:"max_idle"`
	IdleTimeout int    `json:"idle_timeout"`
	MaxActive   int    `json:"max_active"`
}

//  数据库配置项
type Mysql struct {
	Driver            string `json:"driver"`
	Addr              string `json:"addr"` // "root:123@tcp(127.0.0.1:3306)/dbname?charset=utf8"
	Password          string `json:"password"`
	MaxOpenConn       int    `json:"max_open_conn"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	ConnMaxLifetime   int    `json:"conn_max_lifetime"` // 单位毫秒
}

// 日志配置
type LoggerConfig struct {
	Level string `json:"level"`
	Path  string `json:"path"`
	Mode  string `json:"mode"`
}

// jwt 载荷
type Claims struct {
	AccountId uint64 `json:"account_id"`
	OpenId    string `json:"open_id"`
	jwt.StandardClaims
}

/****************请求结构体***********/
type LoginReq struct {
	Code          string `json:"code" binding:"required"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	EncryptedData string `json:"encrypted_data"`
	IV            string `json:"iv"`
}

type LoginResp struct {
	ResCode int    `json:"res_code"`
	Token   string `json:"token"`
	Phone   string `json:"phone"`
}

type RegisterReq struct {
	NickName    string          `json:"nick_name" binding:"required"`
	Email       string          `json:"email" binding:"required,email"`
	OpenId      string          `json:"open_id" binding:"required"`
	Mobile      string          `json:"mobile" binding:"required"`
	Money       decimal.Decimal `json:"money"`
	Avatar      string          `json:"avatar" binding:"required"` // 头像
	Gender      string          `json:"gender" binding:"required"`
	School      string          `json:"school" binding:"required"`
	Major       string          `json:"major" binding:"required"`
	IsAuth      uint8           `json:"is_auth"`
	Description string          `json:"description"`
	CreateTime  int64           `json:"create_time"`
	UpdateTime  int64           `json:"update_time"`
}

type RegisterResp struct {
	ResCode int `json:"res_code"`
}

type QueryAccountResp struct {
}

// 修改账户信息
type ModifyAccountInfoReq struct {
	Id       uint64 `json:"id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"omitempty"`
	NickName string `json:"nick_name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type ModifyAccountInfoResp struct {
}
