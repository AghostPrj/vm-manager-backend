/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:48
 * @Desc:
 */

package router

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func BuildGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	if viper.GetBool(constData.ConfDebugFlagKey) {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
			AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization, x-bm-auth-code"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"},
		}))
	}
	buildApiV1(router)
	return router
}

func StartGinServer(router *gin.Engine) {
	router.Run(viper.GetString(constData.ConfServerListenHostKey) +
		":" + viper.GetString(constData.ConfServerListenPortKey))
}
