/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 14:13
 * @Desc:
 */

package userService

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/constData"
	"github.com/AghostPrj/vm-manager-backend/internal/global"
	"github.com/AghostPrj/vm-manager-backend/internal/model/userModel"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/cryptUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/randomUtils"
	"time"
)

func Logout(authCode string) {
	delete(global.AuthMap, authCode)
}

func Login(username string, password string, otpCode string) (*string, error) {
	user := userModel.User{}
	err := global.DBClient.Where(&userModel.User{Name: username}).First(&user).Error
	if err != nil {
		return nil, err
	}
	if user.Id < 1 || len(user.Name) < 1 {
		return nil, errors.New("user not exists")
	}

	loginSuccess := user.CheckLogin(password, otpCode)

	if !loginSuccess {
		return nil, err
	}
	authCode := saveAuthData(user)
	return &authCode, nil

}

func saveAuthData(user userModel.User) string {
	user.LastOperation = time.Now()
	authCode := randomUtils.RandStringWithLength(32)

	global.AuthMap[authCode] = &user

	return authCode
}

func CreateNewUserWithPassword(name string, password string) (*userModel.User, error) {
	totp, err := cryptUtils.GenerateTotp(name, constData.ApplicationName)
	if err != nil {
		return nil, err
	}
	user := userModel.User{
		Name: name,
		Salt: randomUtils.RandStringWithLength(8),
		Otp:  totp.String(),
	}
	user.Password = user.CalcPasswordHash(password)
	err = global.DBClient.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserWithoutPassword(name string) (*userModel.User, *string, error) {
	password := randomUtils.RandString()
	user, err := CreateNewUserWithPassword(name, password)
	if err != nil {
		return nil, nil, err
	}
	return user, &password, nil
}

func CheckDefaultUserExist() bool {
	user := userModel.User{}
	err := global.DBClient.Where(&userModel.User{Name: constData.DefaultUserName}).First(&user).Error
	if err != nil {
		return false
	}
	return user.Id >= 1
}
