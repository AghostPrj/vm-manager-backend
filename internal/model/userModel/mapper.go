/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:25
 * @Desc:
 */

package userModel

import (
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/cryptUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/randomUtils"
)

func CalcPasswordHash(password string, salt string) string {
	return cryptUtils.HmacSha256(password, salt)
}

func CreateNewUserWithPassword(name string, password string) (*User, error) {
	totp, err := cryptUtils.GenerateTotp(name, global.ApplicationName)
	if err != nil {
		return nil, err
	}
	user := User{
		Name: name,
		Salt: randomUtils.RandStringWithLength(8),
		Otp:  totp.String(),
	}
	user.Password = CalcPasswordHash(password, user.Salt)
	err = global.DBClient.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserWithoutPassword(name string) (*User, *string, error) {
	password := randomUtils.RandString()
	user, err := CreateNewUserWithPassword(name, password)
	if err != nil {
		return nil, nil, err
	}
	return user, &password, nil
}

func CheckDefaultUserExist() bool {
	user := User{}
	err := global.DBClient.Where(&User{Name: global.DefaultUserName}).First(&user).Error
	if err != nil {
		return false
	}
	return user.Id >= 1
}
