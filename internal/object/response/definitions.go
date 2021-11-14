/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:11
 * @Desc:
 */

package response

type BaseResponse struct {
	Code int64       `json:"code"`
	Desc string      `json:"desc,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
