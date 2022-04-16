package middleware

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/ggg17226/aghost-go-base/pkg/utils/collectionUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckUserLoginMiddleware(c *gin.Context) {
	headers := c.Request.Header
	keyList := collectionUtils.GetKeyListFromHeaderMap(&headers)

	var authCodePos int
	authCodePos = -1
	for i, s := range keyList {
		if s == constData.AuthCodeHeaderKey {
			authCodePos = i
			break
		}
	}

	if authCodePos < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.BaseResponse{
			Code: 401,
			Desc: "no auth key",
		})
		return
	}

	authCode := headers.Get(constData.AuthCodeHeaderKey)
	if _, ok := global.AuthMap[authCode]; ok {
		global.AuthMap[authCode].LastOperation = time.Now()
		c.Set(constData.UserInfoContextKey, global.AuthMap[authCode])
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.BaseResponse{
			Code: 401,
			Desc: "auth code expired",
		})
		return
	}
}
