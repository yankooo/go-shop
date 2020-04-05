/**
 * @Author: michael.liu@shopee.com
 * @Date: 2020/3/4
 * @Description:
 */
package repository

import (
	"context"
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/model"
)

// 插入用户
func (e *dbEngine) InsertAccount(ctx context.Context, account *model.Account) error {
	return e.Create(ctx, constant.TABLE_ACCOUNT, account)
}

// 根据id查询批量用户
func (e *dbEngine) QueryMultiAccountById(ctx context.Context, accountIds []uint64) (accounts []*model.Account, err error) {
	if d := e.db.Table(constant.TABLE_ACCOUNT).Where("id IN (?)", accountIds).Find(&accounts); d.Error != nil {
		err = d.Error
	}
	return
}

// 根据用户id查询用户
func (e *dbEngine) QueryAccountById(ctx context.Context, account *model.Account) (err error) {
	d := e.db.Table(constant.TABLE_ACCOUNT).Where("id = ?", account.Id).First(account)
	if d.Error != nil && d.RecordNotFound() {
		return nil
	}
	return err
}

// 根据openId查询用户
func (e *dbEngine) QueryAccountByOpenId(ctx context.Context, openId string) (isExisted bool, err error) {
	var count int
	d := e.db.Table(constant.TABLE_ACCOUNT).Where("open_id = ?", openId).Count(&count)
	if d.Error != nil {
		return false, d.Error
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

// 修改用户
func (e *dbEngine) UpdateAccount(ctx context.Context, account *model.Account) (err error) {
	return e.Update(ctx, constant.TABLE_ACCOUNT, account)
}
