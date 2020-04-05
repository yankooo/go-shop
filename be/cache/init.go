package cache


import (
	"context"
	"fmt"
	"time"
	"github.com/yankooo/school-eco/be/logger"
	"github.com/yankooo/school-eco/be/model"

	"github.com/gomodule/redigo/redis"
)

var _G_RedisPool *redisEngine

type redisEngine struct {
	pool *redis.Pool
}

func NewRedisPool(redisCfg *model.Redis) (*redis.Pool, error) {
	pool := &redis.Pool{}

	redisUrl := fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port)
	pool = &redis.Pool{
		MaxIdle:     redisCfg.MaxIdle,
		MaxActive:   redisCfg.MaxActive,
		IdleTimeout: time.Duration(redisCfg.IdleTimeout) * time.Millisecond,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", redisUrl,
				redis.DialPassword(redisCfg.Password),
				redis.DialDatabase(redisCfg.DB),
				redis.DialConnectTimeout(time.Duration(redisCfg.Timeout)*time.Second),
				redis.DialReadTimeout(time.Duration(redisCfg.Timeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(redisCfg.Timeout)*time.Second))
			if err != nil {
				return nil, err
			}

			return con, nil
		},
	}

	return pool, nil
}

func (re *redisEngine) GetRedisClient() redis.Conn {
	if re.pool == nil {
		logger.Errorf("pool nil")
		return nil
	}

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2) // 2s 取不到连接就返回
	defer cancel()

	conn, err := re.pool.GetContext(ctx)
	if err != nil {
		logger.Debugf("can't get conn from redis pool # %s\n", err.Error())
		return nil
	}

	return conn
}

func RedisEngine() *redisEngine {
	return _G_RedisPool
}

func InitRedisPool(redisConfig *model.Redis) (err error) {
	_G_RedisPool = &redisEngine{}
	_G_RedisPool.pool, err = NewRedisPool(redisConfig)
	return
}
