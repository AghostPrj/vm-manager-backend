/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:32
 * @Desc:
 */

package global

import (
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/AghostPrj/vm-manager-backend/internal/object/response"
	"github.com/robfig/cron"
)

var (
	AuthMap           = make(map[string]*userModel.User)
	AsyncOperationMap = make(map[string]*response.AsyncResponsePayload)
	Cron              = cron.New()
)
