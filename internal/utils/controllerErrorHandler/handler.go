/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/12/2 19:33
 * @Desc:
 */

package controllerErrorHandler

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/constData/errorCode"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

var (
	ErrorMap = make(map[string]*response.BaseResponse)
)

func Init() {
	ErrorMap[errorCode.DataError] = &constData.DataErrorResponse
	addError(400, errorCode.GetUserInfoError)

	addError(400, errorCode.AsyncOperationError)
	addError(403, errorCode.PermissionError)
	addError(404, errorCode.AsyncOperationNotFount)

	addError(401, errorCode.LoginFailedError)

	addError(400, errorCode.NoDpdkSupportError)
	addError(400, errorCode.DeviceNotFoundError)
	addError(400, errorCode.DeviceDriverNotSupportError)

	addError(400, errorCode.NoOvsSupportError)
}

func addError(code int64, errorString string) {
	ErrorMap[errorString] = &response.BaseResponse{
		Code: code,
		Desc: errorString,
	}
}

func HandlerError(context *gin.Context, data interface{}, err error) {
	if err != nil {
		if r, ok := ErrorMap[err.Error()]; ok {
			context.AbortWithStatusJSON(http.StatusOK, r)
			return
		} else {
			pc, _, _, getCallerResult := runtime.Caller(1)
			if !getCallerResult {
				log.WithFields(log.Fields{
					"op":         "HandlerError",
					"step":       "get_caller",
					"err":        err,
					"get_caller": false,
				}).Error()
			} else {
				funcName := runtime.FuncForPC(pc).Name()
				log.WithFields(log.Fields{
					"op":     "HandlerError",
					"source": funcName,
					"err":    err,
				}).Error()
			}
			context.AbortWithStatusJSON(http.StatusOK, constData.SystemErrorResponse)
			return
		}
	} else {
		context.JSON(http.StatusOK, response.BaseResponse{
			Code: 0,
			Desc: "success",
			Data: data,
		})
		return
	}
}
