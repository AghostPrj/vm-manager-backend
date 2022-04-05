package asyncQueryController

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryAsyncResponse(context *gin.Context) {
	userInfoInterface, exist := context.Get(constData.UserInfoContextKey)
	if !exist {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.BaseResponse{
			Code: 500,
			Desc: "get user info error",
		})
		return
	}

	var uid uint64
	if userInfo, ok := userInfoInterface.(*userModel.User); ok {
		uid = userInfo.Id
	} else {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.BaseResponse{
			Code: 500,
			Desc: "get user info error",
		})
		return
	}

	asyncRequestId := context.Param("id")
	if len(asyncRequestId) < 2 {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.BaseResponse{
			Code: 400,
			Desc: "async operation error",
		})
		return
	}
	if asyncResponse, ok := global.AsyncOperationMap[asyncRequestId]; ok {
		if asyncResponse.Uid == uid {
			context.JSON(http.StatusOK, response.BaseResponse{
				Code: 0,
				Data: asyncResponse,
			})
		} else {
			context.JSON(http.StatusForbidden, response.BaseResponse{
				Code: 403,
			})
		}
		return
	} else {
		context.AbortWithStatusJSON(http.StatusNotFound, response.BaseResponse{
			Code: 404,
			Desc: "async operation not found",
		})
		return
	}

}
