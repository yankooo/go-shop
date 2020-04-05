package repository

import (
	"context"
	"github.com/pkg/errors"
)

func (e *dbEngine) Create(ctx context.Context, tabName string, row interface{}) (err error) {
	err = e.db.Table(tabName).Create(row).Error
	if err != nil {
		return errors.WithMessage(err, "Create")
	}
	return nil
}

func (e *dbEngine) Update(ctx context.Context, tabName string, row interface{}) (err error) {
	err = e.db.Table(tabName).Model(row).Updates(row).Error
	if err != nil {
		return errors.WithMessage(err, "Update")
	}
	return nil
}
