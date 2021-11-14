/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:32
 * @Desc:
 */

package global

import (
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/robfig/cron"
)

var (
	AuthMap = make(map[string]*userModel.User)
	Cron    = cron.New()
)
