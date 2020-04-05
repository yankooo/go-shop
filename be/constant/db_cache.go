package constant

import (
	"strconv"
)

// 表名
const (
	TABLE_ACCOUNT        = "account_tab"   // account表
	_TABLE_FOLLOW_PREFIX = "account_book_" // follow表前缀
	_TABLE_FOLLOW_SUBFIX = "_tab"

	COMMON_DEL    = "DEL"
	COMMON_EXPIRE = "expire"
	HASH_SET      = "HSET"
	HASH_MSET     = "HMSET"
	HASH_MGET     = "HMGET"
	HASH_HGETALL  = "HGETALL"
	HASH_DEL      = "HDEL"

	ACCOUNT_INFO_KEY_PREFIX   = "user:"
	ACCOUNT_FIELD_ID          = "id"
	ACCOUNT_FIELD_NICKNAME    = "nick_name"
	ACCOUNT_FIELD_EMAIL       = "email"
	ACCOUNT_FILED_OPENID      = "open_id"
	ACCOUNT_FILED_MOBILE      = "mobile"
	ACCOUNT_FILED_MONEY       = "money"
	ACCOUNT_FILED_AVATAR      = "avatar"
	ACCOUNT_FILED_GENDER      = "gender"
	ACCOUNT_FILED_SCHOOL      = "school"
	ACCOUNT_FILED_MAJOR       = "major"
	ACCOUNT_FILED_DESCR       = "descr"
	ACCOUNT_FILED_UPDATE_TIME = "update_time"
	ACCOUNT_FIELD_CREATE_TIME = "create_time"

	// 一页最多显示数据
	MAX_PAGE_SIZE = 50
)

// 用户发布的售卖图书根据时间分表
func GetFollowTableName(id uint64) string {
	tableIdx := strconv.FormatUint(id%8, 10)
	return _TABLE_FOLLOW_PREFIX + tableIdx + _TABLE_FOLLOW_SUBFIX
}
