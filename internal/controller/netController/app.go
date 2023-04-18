/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 16:00
 * @Desc:
 */

package netController

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/controllerErrorHandler"
	"github.com/gin-gonic/gin"
)

func GetNetTypeList(context *gin.Context) {
	result := make([]string, 0)
	result = append(result, constData.HostNetTypeTap, constData.HostNetTypeDpdk)

	controllerErrorHandler.HandlerError(context, &result, nil)
	return
}
