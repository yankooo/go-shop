package model

import "github.com/shopspring/decimal"

// 数据库model
type Account struct {
	Id          uint64          `json:"id"`
	OpenId      string          `json:"open_id omitempty"`
	Mobile      string          `json:"mobile"`
	Money       decimal.Decimal `json:"money"`
	NickName    string          `json:"nick_name"`
	Email       string          `json:"email"`
	Avatar      string          `json:"avatar"` // 头像
	Gender      string          `json:"gender"`
	School      string          `json:"school"`
	Major       string          `json:"major"`
	Description string          `json:"description"`
	CreateTime  int64           `json:"create_time"`
	UpdateTime  int64           `json:"update_time"`
	SessionKey  string          `json:"session_key"`
}

type Book struct {
	Id         uint32 `json:"id"`
	Title      string `json:"title"`      // 书名
	Pic        string `json:"pic"`        //图片
	Author     string `json:"author"`     // 作者
	Summary    string `json:"summary"`    // 摘要
	Publisher  string `json:"publisher"`  //出版社
	Pubdate    string `json:"pubdate"`    // 出版时间
	Page       string `json:"page"`       // 页数
	Price      string `json:"price"`      // 价格
	Isbn       string `json:"isbn"`       // ISBN
	Keyword    string `json:"keyword"`    // 主题词
	Edition    string `json:"edition"`    // 版次
	Impression string `json:"impression"` // 印次
}
