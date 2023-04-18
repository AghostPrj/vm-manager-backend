package router

import (
	"github.com/AghostPrj/vm-manager-backend/internal/controller/asyncQueryController"
	"github.com/AghostPrj/vm-manager-backend/internal/controller/netController"
	"github.com/AghostPrj/vm-manager-backend/internal/controller/userController"
	"github.com/AghostPrj/vm-manager-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func buildApiV1(router *gin.Engine) {
	groupV1 := router.Group("/api/v1")

	groupV1.POST("/login", userController.Login)
	groupV1.Any("/logout", userController.Logout)

	buildApiV1AdminApis(groupV1.Group("/admin"))
	buildApiV1AsyncResultApis(groupV1.Group("/async"))

}

func buildApiV1AsyncResultApis(router *gin.RouterGroup) {
	router.Use(middleware.CheckUserLoginMiddleware)
	router.GET("/result/:id", asyncQueryController.QueryAsyncResponse)
}

func buildApiV1AdminApis(router *gin.RouterGroup) {
	router.Use(middleware.CheckUserLoginMiddleware)

	router.GET("/net/type", netController.GetNetTypeList)
	router.GET("/net/system/nic", netController.GetSystemNicList)
	router.POST("/net/dpdk/bind/:domain/:driver", netController.BindNicDriver)

	router.GET("/vm")
	router.GET("/vm/:id")
}
