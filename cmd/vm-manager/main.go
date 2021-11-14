/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 10:49
 * @Desc:
 */

package main

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/initializator"
	"github.com/AghostPrj/vm-manager-backend/internal/router"
	"github.com/AghostPrj/vm-manager-backend/internal/task"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	initializator.InitApp()
	task.InitAndStartCron()
	ginRouter := router.BuildGinRouter()
	log.WithFields(log.Fields{
		"op":   "startup",
		"host": viper.GetString(constData.ConfServerListenHostKey),
		"port": viper.GetString(constData.ConfServerListenPortKey),
	}).Info("http server start")
	router.StartGinServer(ginRouter)

}
