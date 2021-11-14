/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 11:13
 * @Desc:
 */

package initializator

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func checkHostNetType() {
	hostNetType := viper.GetString(ConfHostNetTypeConfigKey)
	if len(hostNetType) < 1 {
		log.WithField("op", "startup").Panic("host net type not set")
	}
	if !(hostNetType == "tap" || hostNetType == "dpdk") {
		log.WithField("op", "startup").Panic("host net type setting error")
	}
}

func checkHttpListenConfig() {
	listenPort := viper.GetInt(ConfServerListenPortKey)
	if !(listenPort > 0 && listenPort < 65536) {
		log.WithField("op", "startup").Panic("listen port out of range")
	}
}
