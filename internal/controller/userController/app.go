/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:03
 * @Desc:
 */

package userController

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/constData/errorCode"
	"github.com/AghostPrj/vm-manager-backend/internal/object/request"
	"github.com/AghostPrj/vm-manager-backend/internal/service/userService"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/controllerErrorHandler"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/requestUtils"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	loginReq := request.LoginRequest{}
	err := requestUtils.DecodeJsonRequestBody(context, &loginReq)
	if err != nil {
		return
	}
	if !loginReq.Check() {
		controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.DataError))
		return
	}

	authCode, err := userService.Login(loginReq.Username, loginReq.Password, loginReq.OtpCode)

	controllerErrorHandler.HandlerError(context, authCode, err)
}

func Logout(context *gin.Context) {
	headers := context.Request.Header
	authCode := headers.Get(constData.AuthCodeHeaderKey)
	if len(authCode) > 0 {
		userService.Logout(authCode)
	}
	controllerErrorHandler.HandlerError(context, nil, nil)
}
