/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/5/31 14:54
 * @Desc:
 */

package constData

import "github.com/AghostPrj/vm-manager-backend/internal/object/response"

var (
	DataErrorResponse = response.BaseResponse{
		Code: 400,
		Desc: "request data error",
	}
	SystemErrorResponse = response.BaseResponse{
		Code: 500,
		Desc: "system error",
	}
)
