/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:02
 * @Desc:
 */

package userModel

type User struct {
	Id       uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id"`
	Name     string `gorm:"type:varchar(128);Column:username;NOT NULL;uniqueIndex"`
	Password string `gorm:"type:varchar(256);Column:password;NOT NULL"`
	Salt     string `gorm:"type:char(8);Column:salt;NOT NULL"`
	Otp      string `gorm:"type:varchar(512);Column:otp;NOT NULL"`
}
