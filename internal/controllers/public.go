package controllers

import (
	"api_admin_server/api/file_server"
	"api_admin_server/pkg/resp"
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type PublicController struct {
}

func (s *PublicController) UploadFile(c *gin.Context) {
	fileHead, err := c.FormFile("file")
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	file, err := fileHead.Open()
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	client := file_server.GetClient()
	ret, err := client.UploadLocal(context.Background(), &file_server.UploadLocalReq{FileBytes: fileBytes})
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	if ret.Res.Code != 200 {
		resp.RespGeneralErr(c, ret.Res.Msg)
		return
	}
	resp.RespOk(c, ret)
	return
}
