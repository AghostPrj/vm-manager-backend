/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2023/4/18 14:39
 * @Desc:
 */

package errorCode

const (
	DataError        = "data error"
	GetUserInfoError = "get user info error"

	AsyncOperationError    = "async operation error"
	PermissionError        = "permission error"
	AsyncOperationNotFount = "async operation not found"

	LoginFailedError = "login failed"

	NoDpdkSupportError          = "no dpdk support"
	DeviceNotFoundError         = "device not found"
	DeviceDriverNotSupportError = "driver not support"
)
