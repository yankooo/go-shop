/**
 * @Author: michael.liu@shopee.com
 * @Date: 2020/3/4
 * @Description:
 */
package cache

import (
	"context"
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/logger"
	"github.com/yankooo/school-eco/be/model"
	"github.com/yankooo/school-eco/be/utils"
	"time"

	"github.com/gomodule/redigo/redis"
)

func (re *redisEngine) QuerySingleAccountInfo(ctx context.Context, accountId uint64, opt ...interface{}) (*model.Account, error) {
	return re.Query(ctx, AccountInfoKey(accountId), mapperAccount)
}

func mapperAccount(rdCli redis.Conn, key string) (account *model.Account, err error) {
	t := time.Now()
	logger.Debugf("mapperAccount start with %d", t.UnixNano())

	reply, e := redis.StringMap(rdCli.Do(constant.HASH_HGETALL, key))
	if e != nil || len(reply) == 0 {
		return nil, e
	}
	// 反序列化，映射回结构体
	logger.Debugf("redis query account: %+v", reply)
	account = &model.Account{
		Id:         utils.StringConvertUint64(reply[constant.ACCOUNT_FIELD_ID]),
		NickName:   reply[constant.ACCOUNT_FIELD_NICKNAME],
		Email:      reply[constant.ACCOUNT_FIELD_EMAIL],
		CreateTime: utils.StringConvertInt64(reply[constant.ACCOUNT_FIELD_CREATE_TIME]),
	}

	logger.Debugf("mapperAccount end with %d", time.Now().Sub(t))
	return
}

// 插入账户信息
func (re *redisEngine) InsertAccountInfo(ctx context.Context, account *model.Account) (err error) {
	// 账号信息
	if err  = re.Do(ctx, constant.HASH_MSET, AccountInfoKey(account.Id),
		constant.ACCOUNT_FIELD_ID, account.Id,
		constant.ACCOUNT_FILED_OPENID, account.OpenId,
		constant.ACCOUNT_FIELD_NICKNAME, account.NickName,
		constant.ACCOUNT_FIELD_EMAIL, account.Email,
		constant.ACCOUNT_FILED_MOBILE, account.Mobile,
		constant.ACCOUNT_FILED_MONEY, account.Money,
		constant.ACCOUNT_FILED_AVATAR, account.Avatar,
		constant.ACCOUNT_FILED_GENDER, account.Gender,
		constant.ACCOUNT_FILED_SCHOOL, account.School,
		constant.ACCOUNT_FILED_MAJOR, account.Major,
		constant.ACCOUNT_FILED_DESCR, account.Description,
		constant.ACCOUNT_FILED_UPDATE_TIME, account.UpdateTime,
		constant.ACCOUNT_FIELD_CREATE_TIME, account.CreateTime,
	); err != nil {
		return
	}
	return re.Do(ctx, constant.COMMON_EXPIRE, AccountInfoKey(account.Id), 60*60)//设置过期时间 1hour
}

// 更新账户信息 覆盖除了关注者字段的所有信息 TODO 不合理
func (re *redisEngine) UpdateAccountInfoObject(ctx context.Context, account *model.Account) error {
	return re.Do(ctx, constant.HASH_MSET, AccountInfoKey(account.Id),
		constant.ACCOUNT_FIELD_ID, account.Id,
		constant.ACCOUNT_FIELD_NICKNAME, account.NickName,
		constant.ACCOUNT_FIELD_EMAIL, account.Email,
		constant.ACCOUNT_FIELD_CREATE_TIME, account.CreateTime)
}

// 更新账户指定字段
func (re *redisEngine) UpdateAccountInfoField(ctx context.Context, args ...interface{}) error {
	return re.Do(ctx, constant.HASH_MSET, args...)
}

// 删除账户信息
func (re *redisEngine) RemoveAccountInfo(ctx context.Context, account *model.Account) error {
	return re.Do(ctx, constant.COMMON_DEL, AccountInfoKey(account.Id))
}

func (re *redisEngine) RemoveAccountFollow(ctx context.Context, args ...interface{}) error {
	return re.Do(ctx, constant.HASH_DEL, args...)
}
