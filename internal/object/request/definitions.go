/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:08
 * @Desc:
 */

package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OtpCode  string `json:"otp_code"`
}

func (r *LoginRequest) Check() bool {
	if len(r.Username) < 1 || len(r.Password) < 1 || len(r.OtpCode) < 1 {
		return false
	}
	return true
}
