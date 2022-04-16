/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 11:13
 * @Desc:
 */

package initializator

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func checkHttpListenConfig() {
	listenPort := viper.GetInt(constData.ConfServerListenPortKey)
	if !(listenPort > 0 && listenPort < 65536) {
		log.WithField("op", "startup").Panic("listen port out of range")
	}
}
