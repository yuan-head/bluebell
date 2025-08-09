package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// 处理注册请求函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数绑定失败，返回错误信息
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是Validation错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 不是Validation错误，直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidPassword, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUp with invalid param", zap.Error(fmt.Errorf("两次密码不一致")))
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//}

	// 2. 业务处理
	// 业务规则校验通过，继续处理注册逻辑
	if err := logic.SignUp(p); err != nil {
		fmt.Println("SignUp err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func LoginHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamLogin)
	// 2. 业务处理
	if err := c.ShouldBindJSON(p); err != nil {
		// 参数绑定失败，返回错误信息
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是Validation错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 不是Validation错误，直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务规则校验通过，继续处理登录逻辑
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		fmt.Println("Login err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, token)
}
