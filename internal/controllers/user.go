package controllers

import (
	"api_admin_server/api/user_server"
	"api_admin_server/pkg/resp"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserController struct {
}

func (s *UserController) Login(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	if email == "" || password == "" {
		resp.RespParamErr(c, "账号或密码不能为空")
		return
	}
	userServer := user_server.GetClient()
	req := &user_server.AdminLoginReq{Email: email, Password: password}
	ret, err := userServer.AdminLogin(context.Background(), req)
	if err != nil {
		resp.RespInternalErr(c, "账号或密码错误")
		return
	}
	resp.RespOk(c, ret)
	return
}

//
func (s *UserController) Register(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	if email == "" || code == "" {
		resp.RespParamErr(c, "账号或验证码不能为空")
		return
	}
	userServer := user_server.GetClient()
	req := &user_server.AdminRegisterReq{Email: email, Code: code}
	ret, err := userServer.AdminRegister(context.Background(), req)
	if err != nil {
		resp.RespInternalErr(c, "验证码无效")
		return
	}
	resp.RespOk(c, ret)
	return
}

//
func (s *UserController) SendEmailCode(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		resp.RespParamErr(c, "邮箱发送失败，请检查邮箱是否正确")
		return
	}
	from := c.Query("from")
	client := user_server.GetClient()
	ret, err := client.AdminSendEmailCode(context.Background(), &user_server.AdminSendEmailCodeReq{Email: email, From: from})
	if err != nil {
		resp.RespInternalErr(c, "邮箱发送失败，请检查邮箱是否正确")
		return
	}
	if ret.Code != 200 {
		resp.RespParamErr(c, ret.Msg)
		return
	}
	resp.RespOk(c)
	return
}

func (s *UserController) DataFull(c *gin.Context) {
	req := new(user_server.AdminDataFullReq)
	err := c.BindJSON(req)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	if req.Uid <= 0 || req.Name == "" || req.Password == "" {
		resp.RespParamErr(c)
		return
	}
	client := user_server.GetClient()
	res, err := client.AdminDataFull(context.Background(), req)
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	if res.Code != 200 {
		resp.RespGeneralErr(c, res.Msg)
		return
	}
	resp.RespOk(c)
	return
}

func (s *UserController) PasswordBack(c *gin.Context) {
	email := c.Query("email")
	code := c.Query("code")
	password := c.Query("password")
	if email == "" || code == "" || password == "" {
		resp.RespParamErr(c)
		return
	}
	if len(strings.TrimSpace(password)) < 7 {
		resp.RespGeneralErr(c, "密码最少需要6位")
		return
	}
	client := user_server.GetClient()
	res, err := client.AdminPasswordBack(context.Background(), &user_server.AdminPasswordBackReq{
		Email:    email,
		Code:     code,
		Password: password,
	})
	if err != nil {
		resp.RespGeneralErr(c, "验证码错误")
		return
	}
	if res.Code != 200 {
		resp.RespGeneralErr(c, "验证码错误")
		return
	}
	resp.RespOk(c)
	return
}
