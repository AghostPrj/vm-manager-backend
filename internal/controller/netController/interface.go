/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 14:37
 * @Desc:
 */

package netController

import (
	"github.com/AghostPrj/vm-manager-backend/internal/service/networkService"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/controllerErrorHandler"
	"github.com/gin-gonic/gin"
)

func GetSystemNicList(context *gin.Context) {
	nic, err := networkService.GetSystemNic()

	controllerErrorHandler.HandlerError(context, nic, err)
	return
}

func BindNicDriver(context *gin.Context) {
	err := networkService.BindNicDriver(context.Param("domain"), context.Param("driver"))

	controllerErrorHandler.HandlerError(context, nil, err)
	return
}
