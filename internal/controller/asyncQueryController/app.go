package asyncQueryController

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/constData/errorCode"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/controllerErrorHandler"
	"github.com/gin-gonic/gin"
)

func QueryAsyncResponse(context *gin.Context) {
	userInfoInterface, exist := context.Get(constData.UserInfoContextKey)
	if !exist {
		controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.GetUserInfoError))
		return
	}

	var uid uint64
	if userInfo, ok := userInfoInterface.(*userModel.User); ok {
		uid = userInfo.Id
	} else {
		controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.GetUserInfoError))
		return
	}

	asyncRequestId := context.Param("id")
	if len(asyncRequestId) < 2 {
		controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.AsyncOperationError))
		return
	}
	if asyncResponse, ok := global.AsyncOperationMap[asyncRequestId]; ok {
		if asyncResponse.Uid == uid {
			controllerErrorHandler.HandlerError(context, asyncResponse, nil)
		} else {
			controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.PermissionError))
		}
		return
	} else {
		controllerErrorHandler.HandlerError(context, nil, errors.New(errorCode.AsyncOperationNotFount))
		return
	}

}
