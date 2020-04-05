package repository

import (
	"context"
	"github.com/yankooo/school-eco/be/model"
	"testing"
)

func TestDbEngine_InsertAccount(t *testing.T) {
	if err := InitDbEngine(&model.Mysql{
		Driver:            "mysql",
		Addr:              "root:123456@tcp(127.0.0.1:3306)/ums_db?charset=utf8mb4",
		Password:          "123456",
		MaxOpenConn:       1000,
		MaxIdleConnection: 5000,
		ConnMaxLifetime:   5000,
	}); err != nil {
		t.Logf("initDbEngine err : %+v", err)
	}

	account := model.Account{
		Id:         0,
		NickName:   "nick",
		Email:      "123@qq.com",
		CreateTime: 124356,
	}
	if err := GormDb().InsertAccount(context.TODO(), &account); err != nil {
		t.Logf("insert db err: %+v", err)
	}
	t.Logf("success insert %+v", account)
}


func TestDbEngine_QueryAccountById(t *testing.T) {
	if err := InitDbEngine(&model.Mysql{
		Driver:            "mysql",
		Addr:              "root:123456@tcp(127.0.0.1:3306)/ums_db?charset=utf8mb4",
		Password:          "123456",
		MaxOpenConn:       1000,
		MaxIdleConnection: 5000,
		ConnMaxLifetime:   5000,
	}); err != nil {
		t.Logf("initDbEngine err : %+v", err)
	}

	/*accoutList, err := GormDb().QueryAccountById(context.TODO(), []uint64{1, 3, 20})
	if err != nil {
		t.Logf("insert db err: %+v", err)
	}
	for _, accout := range accoutList {
		t.Logf("success query %+v\n", accout)
	}*/
}

func TestDbEngine_UpdateAccount(t *testing.T) {
	if err := InitDbEngine(&model.Mysql{
		Driver:            "mysql",
		Addr:              "root:123456@tcp(127.0.0.1:3306)/ums_db?charset=utf8mb4",
		Password:          "123456",
		MaxOpenConn:       1000,
		MaxIdleConnection: 5000,
		ConnMaxLifetime:   5000,
	}); err != nil {
		t.Logf("initDbEngine err : %+v", err)
	}

	account := model.Account{
		Id:       9999973,
		//NickName:   "nick",
		//Email:      "123@qq.com",
		CreateTime: 999999,
	}
	err := GormDb().UpdateAccount(context.TODO(), &account)
	if err != nil {
		t.Logf("insert db err: %+v", err)
	}

	t.Logf("success update %+v\n", account)
}
