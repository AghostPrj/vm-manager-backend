/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 15:02
 * @Desc:
 */

package task

import (
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/spf13/viper"
	"time"
)

func InitAndStartCron() {
	buildRefreshAuthTask()
	global.Cron.Start()
}

func buildRefreshAuthTask() {
	global.Cron.AddFunc("*/30 * * * * ?", func() {
		needDeleteList := make([]string, 0)
		for key, user := range global.AuthMap {
			if user.LastOperation.Add(viper.GetDuration(constData.ConfAuthExpireTimeKey)).Before(time.Now()) {
				needDeleteList = append(needDeleteList, key)
			}
		}
		for _, s := range needDeleteList {
			delete(global.AuthMap, s)
		}
	})
}
