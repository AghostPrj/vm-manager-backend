/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:03
 * @Desc:
 */

package userController

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/object/request"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/AghostPrj/vm-manager-backend/internal/service/userService"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/requestUtils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	loginReq := request.LoginRequest{}
	err := requestUtils.DecodeJsonRequestBody(context, &loginReq)
	if err != nil {
		return
	}
	if !loginReq.Check() {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.BaseResponse{
			Code: 500,
			Desc: "body error",
		})
		return
	}

	authCode, err := userService.Login(loginReq.Username, loginReq.Password, loginReq.OtpCode)

	if err != nil || authCode == nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, response.BaseResponse{
			Code: 401,
			Desc: "login error",
		})
		return
	}
	context.JSON(200, response.BaseResponse{
		Code: 0,
		Data: *authCode,
	})
	return
}

func Logout(context *gin.Context) {
	headers := context.Request.Header
	authCode := headers.Get(constData.AuthCodeHeaderKey)
	if len(authCode) > 0 {
		userService.Logout(authCode)
	}
	context.Status(http.StatusOK)
}
