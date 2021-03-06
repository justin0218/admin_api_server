package middleware

import (
	"api_admin_server/api/auth_server"
	"api_admin_server/pkg/resp"
	"api_admin_server/store"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var config = new(store.Config)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if new(store.Config).Get().Runmode == "debug" {
			c.Next()
			return
		}
		token := c.Query("token")
		uid := c.Query("uid")
		uidInt64, err := strconv.ParseInt(uid, 10, 64)
		if err != nil {
			resp.RespCode(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		conn := auth_server.GetClient()
		req := &auth_server.VerifyTokenReq{}
		req.TokenType = auth_server.TokenType_ADMIN
		req.Uid = uidInt64
		req.Token = token
		ret, err := conn.VerifyToken(context.Background(), req)
		if err != nil || ret.Code != 200 {
			resp.RespCode(c, http.StatusUnauthorized, "未授权")
			c.Abort()
			return
		}
		c.Set("uid", ret.Uid)
		c.Next()
	}
}
