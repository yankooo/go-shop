/**
 * @Author: michael.liu@shopee.com
 * @Date: 2020/3/5
 * @Description:
 */
package cache

import (
	"context"
	"errors"
	"strconv"
	strings "strings"
	"time"
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/logger"
	"github.com/yankooo/school-eco/be/model"

	"github.com/gomodule/redigo/redis"
)

// redis 单次查询
func (re *redisEngine) Query(ctx context.Context, key string, ops func(conn redis.Conn, key string) (*model.Account, error)) (*model.Account, error) {
	rdCli := re.GetRedisClient()
	if rdCli == nil {
		return nil, errors.New("cant get redis conn")
	}
	defer rdCli.Close()

	t := time.Now()
	logger.Debugf("Query start with %d", t.UnixNano())

	account, err := ops(rdCli, key)

	logger.Debugf("Query end with %d", time.Now().Sub(t))

	return account, err
}

// 通过管道查询 TODO
func (re *redisEngine) QueryByPipeline(ctx context.Context, keySuffix []string, ops func(conn redis.Conn, keySuffix []string) ([]*model.Account, error)) ([]*model.Account, error) {
	rdCli := re.GetRedisClient()
	if rdCli == nil {
		return nil, errors.New("cant get redis conn")
	}
	defer rdCli.Close()
	return ops(rdCli, keySuffix)
}

// 执行redis操作
func (re *redisEngine) Do(ctx context.Context, commandName string, args ...interface{}) (err error) {
	rdCli := re.GetRedisClient()
	if rdCli == nil {
		return errors.New("cant get redis conn")
	}
	defer rdCli.Close()

	_, err = rdCli.Do(commandName, args...)
	return
}

// redis string object key for account info
func AccountInfoKey(accountId uint64) string {
	return constant.ACCOUNT_INFO_KEY_PREFIX + strconv.FormatUint(accountId, 10)
}

func ParseFollowingList(str string) []string {
	if len(str) <= 1 || str[1:len(str)-1] == "" {
		return []string{}
	}
	return strings.Split(str[1:len(str)-1], " ")
}
