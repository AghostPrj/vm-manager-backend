/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:02
 * @Desc:
 */

package userModel

import (
	"github.com/AghostPrj/vm-manager-backend/internal/utils/cryptUtils"
	"time"
)

type User struct {
	Id            uint64    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id;<-:false"`
	Name          string    `gorm:"type:varchar(128);Column:username;NOT NULL;uniqueIndex"`
	Password      string    `gorm:"type:varchar(256);Column:password;NOT NULL"`
	Salt          string    `gorm:"type:char(8);Column:salt;NOT NULL"`
	Otp           string    `gorm:"type:varchar(512);Column:otp;NOT NULL"`
	LastOperation time.Time `gorm:"-"`
}

func (user *User) CheckLogin(password string, otpCode string) bool {

	hash := user.CalcPasswordHash(password)

	if hash != user.Password {
		return false
	}

	return cryptUtils.ValidateTotp(user.Otp, otpCode)

}
func (user *User) CalcPasswordHash(password string) string {
	return cryptUtils.HmacSha256(password, user.Salt)
}
