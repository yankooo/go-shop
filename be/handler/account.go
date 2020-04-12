package handler

// cd35a6b862497e74fc3ffe414f559349
// wxfa3ab17a6f7a4b6a
import (
	"github.com/medivhzhan/weapp/v2"
	"context"
	"fmt"
	"github.com/yankooo/school-eco/be/cache"
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/logger"
	"github.com/yankooo/school-eco/be/model"
	repo "github.com/yankooo/school-eco/be/repository"
)

func RegisterAccount(ctx context.Context, req *model.Account) (resp *model.RegisterResp, err error) {
	// 1. 查重 //TODO 如果有非法调用，会有问题
	logger.Debugf("req %+v", *req)
	fmt.Printf("req %+v\n", *req)
	if isExistEd, e := repo.GormDb().QueryAccountByOpenId(ctx, req.OpenId); e != nil {
		return nil , e
	} else if isExistEd {
		return &model.RegisterResp{ResCode:constant.DuplicateRegister}, nil
	}

	// 2. 插入数据库
	if err = repo.GormDb().InsertAccount(ctx, req); err != nil {
		logger.Errorf("account insert fail with: %+v", err)
	}

	// 3. 写入缓存
	if err = cache.RedisEngine().InsertAccountInfo(ctx, req); err != nil {
		// 缓存写失败也无所谓，查询会去数据库查询 TODO 大量缓存不命中会
		logger.Errorf("cache insert fail with: %+v", err)
	}
	return &model.RegisterResp{ResCode:constant.AccountRegisterSuccess}, nil
}

func Login(ctx context.Context, req *model.LoginReq) (*model.Account, error) {

	res, err := weapp.Login("wxfa3ab17a6f7a4b6a", "cd35a6b862497e74fc3ffe414f559349", "code")
	if err != nil {
		// 处理一般错误信息

	}

	if err := res.GetResponseError(); err !=nil {
		// 处理微信返回错误信息

	}
	fmt.Printf("%+v", res)
	/*// 校验参数
	if req.UserName == "" {
		return &pb.QuerySingleAccountInfoResp{Res: GenInValidParameterRes()}, nil
	}

	// 缓存不命中查询数据库
	var (
		account     *pb.Account
		accountInfo = &model.Account{UserName: req.UserName}
		err         error
	)

	if accountInfo, err = cache.RedisEngine().QuerySingleAccountInfo(ctx, accountInfo.UserName); err != nil {
		return &pb.QuerySingleAccountInfoResp{Res: GenRedisDealRes(err.Error())}, nil
	}
	if accountInfo == nil {
		// 查询数据库
		accountInfo = &model.Account{UserName: req.UserName}
		if err := repo.GormDb().QueryAccountByName(ctx, accountInfo); err != nil {
			return &pb.QuerySingleAccountInfoResp{Res: GenDBDealRes(err.Error())}, nil
		}

		// 刷新进缓存
		//go func() {
		_ = cache.RedisEngine().InsertAccountInfo(ctx, accountInfo)
		//}()
	}

	// 返回数据
	account = &pb.Account{
		Id:         accountInfo.Id,
		UserName:   accountInfo.UserName,
		Password:   accountInfo.Password,
		NickName:   accountInfo.NickName,
		Email:      accountInfo.Email,
		CreateTime: accountInfo.CreateTime}*/
	return &model.Account{}, nil
}

func QuerySingleAccountInfo(ctx context.Context, accountId uint64) (*model.QueryAccountResp, error) {
	/*// 校验参数
	if req.UserName == "" {
		return &pb.QuerySingleAccountInfoResp{Res: GenInValidParameterRes()}, nil
	}

	// 缓存不命中查询数据库
	var (
		account     *pb.Account
		accountInfo = &model.Account{UserName: req.UserName}
		err         error
	)

	if accountInfo, err = cache.RedisEngine().QuerySingleAccountInfo(ctx, accountInfo.UserName); err != nil {
		return &pb.QuerySingleAccountInfoResp{Res: GenRedisDealRes(err.Error())}, nil
	}
	if accountInfo == nil {
		// 查询数据库
		accountInfo = &model.Account{UserName: req.UserName}
		if err := repo.GormDb().QueryAccountByName(ctx, accountInfo); err != nil {
			return &pb.QuerySingleAccountInfoResp{Res: GenDBDealRes(err.Error())}, nil
		}

		// 刷新进缓存
		//go func() {
		_ = cache.RedisEngine().InsertAccountInfo(ctx, accountInfo)
		//}()
	}

	// 返回数据
	account = &pb.Account{
		Id:         accountInfo.Id,
		UserName:   accountInfo.UserName,
		Password:   accountInfo.Password,
		NickName:   accountInfo.NickName,
		Email:      accountInfo.Email,
		CreateTime: accountInfo.CreateTime}*/
	return &model.QueryAccountResp{}, nil
}

func ModifyAccountInfo(ctx context.Context, req *model.ModifyAccountInfoReq) (*model.ModifyAccountInfoResp, error) {
	// 2. 先删除缓存
	/*if err := cache.RedisEngine().RemoveAccountInfo(ctx, accountInfo); err != nil {
		logger.Errorf("cache remove account fail with: %+v", err)
	}

	// 3. 数据库更新
	if err := repo.GormDb().UpdateAccount(ctx, accountInfo); err != nil {
		logger.Errorf("account update fail with: %+v", err)
	}

	// 4. 从数据库获取最新数据，写入缓存， 写入失败也返回成功，缓存没有数据，会去查询数据库
	//go func() {
	if err := repo.GormDb().QueryAccountByName(ctx, accountInfo); err == nil {
		_ = cache.RedisEngine().InsertAccountInfo(ctx, accountInfo)
	}*/
	return &model.ModifyAccountInfoResp{}, nil
}
