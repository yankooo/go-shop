package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yankooo/school-eco/be/constant"
	"github.com/yankooo/school-eco/be/handler"
	"github.com/yankooo/school-eco/be/model"
	"github.com/yankooo/school-eco/be/transport/http/m_token"
	"github.com/yankooo/school-eco/be/utils"
	"net/http"
)

type bookSeller struct {
}

func SendResponse(c *gin.Context, httpCode int, retCode int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"ret_code": retCode,
		"message":  msg,
		"data":     data,
	})
}

func (b *bookSeller) RegisterAccount(c *gin.Context) {
	// 参数校验
	var registerReq = &model.RegisterReq{}
	if err := c.BindJSON(registerReq); err != nil {
		SendResponse(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error(), nil)
		return
	}

	fmt.Printf("req %+v\n", *registerReq)

	var (
		err          error
		registerResp *model.RegisterResp

		account = &model.Account{
			Mobile:     registerReq.Mobile,
			NickName:   registerReq.NickName,
			OpenId:     registerReq.OpenId,
			Email:      registerReq.Email,
			Avatar:     registerReq.Avatar,
			Gender:     registerReq.Gender,
			School:     registerReq.School,
			Major:      registerReq.Major,
			CreateTime: utils.GetTimeNowUnix(),
			UpdateTime: utils.GetTimeNowUnix(),
		}
	)

	if registerResp, err = handler.RegisterAccount(context.Background(), account); err != nil {
		// 操作出错
		SendResponse(c, http.StatusInternalServerError, constant.WEB_SERVER_DEAL_ERROR, err.Error(), nil)
		return
	}
	SendResponse(c, http.StatusOK, constant.DEAL_SUCCESS, constant.GetMsg(constant.DEAL_SUCCESS), registerResp)
}

func (b *bookSeller) Login(c *gin.Context) {
	// 参数校验
	var loginReq = &model.LoginReq{}
	if err := c.BindJSON(loginReq); err != nil {
		SendResponse(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error(), nil)
		return
	}

	var (
		err     error
		account *model.Account
	)

	if account, err = handler.Login(context.Background(), loginReq); err != nil {
		// 操作出错
		SendResponse(c, http.StatusInternalServerError, constant.WEB_SERVER_DEAL_ERROR, err.Error(), nil)
		return
	}
	// 生成token, 放在cookie中返回
	generateToken(c, account)
}

// 生成令牌
func generateToken(c *gin.Context, account *model.Account) {
	token, err := m_token.GenerateToken(account.Id, account.OpenId)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, constant.WEB_SERVER_DEAL_ERROR, err.Error(), nil)
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	SendResponse(c, http.StatusOK, constant.DEAL_SUCCESS, "login success", model.LoginResp{
		ResCode: constant.Success,
		Phone:   account.Mobile,
	})
	return
}

func (b *bookSeller) QuerySingleAccountInfo(c *gin.Context) {
	// 参数校验
	var id uint64
	accountId, ok := c.Get("account-id")
	if id, ok = accountId.(uint64); !ok {
		SendResponse(c, http.StatusBadRequest, constant.INVALID_PARAMS, "", nil)
		return
	}

	var (
		err                  error
		queryAccountInfoResp *model.QueryAccountResp
	)

	// rpc调用，检查用户合法性
	if queryAccountInfoResp, err = handler.QuerySingleAccountInfo(context.Background(), id); err != nil {
		// rpc 调用出错
		SendResponse(c, http.StatusInternalServerError, constant.WEB_SERVER_DEAL_ERROR, err.Error(), nil)
		return
	}

	SendResponse(c, http.StatusOK, constant.DEAL_SUCCESS, "query success", queryAccountInfoResp)
}

func (b *bookSeller) ModifyAccountInfo(c *gin.Context) {
	// 参数校验
	var modifyInfoReq = &model.ModifyAccountInfoReq{}
	if err := c.BindJSON(modifyInfoReq); err != nil {
		SendResponse(c, http.StatusBadRequest, constant.INVALID_PARAMS, err.Error(), nil)
		return
	}

	// handler调用，修改账号信息
	var (
		err                   error
		modifyAccountInfoResp *model.ModifyAccountInfoResp
	)
	if modifyAccountInfoResp, err = handler.ModifyAccountInfo(context.Background(), modifyInfoReq); err != nil {
		// rpc 调用出错
		SendResponse(c, http.StatusInternalServerError, constant.WEB_SERVER_DEAL_ERROR, err.Error(), nil)
		return
	}

	SendResponse(c, http.StatusOK, constant.DEAL_SUCCESS, constant.GetMsg(constant.DEAL_SUCCESS), modifyAccountInfoResp)
}
