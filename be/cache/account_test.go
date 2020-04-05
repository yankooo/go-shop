/**
 * @Author: michael.liu@shopee.com
 * @Date: 2020/3/5
 * @Description:
 */
package cache

import (
	"context"
	"github.com/yankooo/school-eco/be/model"
	"testing"
	"time"
)

func TestRedisEngine_InsertAccountInfo(t *testing.T) {
	if err := InitRedisPool(model.Redis{
		Host:        "127.0.0.1",
		Port:        6379,
		Password:    "redis",
		Timeout:     10,
		DB:          0,
		MaxIdle:     1000,
		IdleTimeout: 10,
		MaxActive:   1000,
	}); err != nil {
		t.Logf("gInit redis err : %+v", err)
	}

	account := model.Account{
		Id:         3,
		NickName:   "niccc",
		Email:      "ccc@qq.com",
		CreateTime: 123,
	}
	if err := RedisEngine().InsertAccountInfo(context.TODO(), &account); err != nil {
		t.Logf("insert into cache err :%+v", err)
	} else {
		t.Logf("insert success")
	}
}

func TestRedisEngine_RemoveAccountInfo(t *testing.T) {
	if err := InitRedisPool(model.Redis{
		Host:        "127.0.0.1",
		Port:        6379,
		Password:    "redis",
		Timeout:     10,
		DB:          0,
		MaxIdle:     1000,
		IdleTimeout: 10,
		MaxActive:   1000,
	}); err != nil {
		t.Logf("gInit redis err : %+v", err)
	}

	account := model.Account{
		Id:         3,
		NickName:   "niccc",
		Email:      "ccc@qq.com",
		CreateTime: 123,
	}
	if err := RedisEngine().RemoveAccountInfo(context.TODO(), &account); err != nil {
		t.Logf("remove from cache err :%+v", err)
	} else {
		t.Logf("remove success")
	}
}

func TestRedisEngine_QuerySingleAccountInfo(t *testing.T) {
	if err := InitRedisPool(model.Redis{
		Host:        "127.0.0.1",
		Port:        6379,
		Password:    "redis",
		Timeout:     10,
		DB:          0,
		MaxIdle:     1000,
		IdleTimeout: 10,
		MaxActive:   1000,
	}); err != nil {
		t.Logf("gInit redis err : %+v", err)
	}

	times := time.Now()
	if account, err := RedisEngine().QuerySingleAccountInfo(context.TODO(), 0); err != nil {
		t.Logf("query from cache err :%+v", err)
	} else {
		t.Logf("query success %+v", account)
	}
	t.Logf("end %d", time.Now().Sub(times))
}