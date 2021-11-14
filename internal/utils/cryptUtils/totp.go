/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:19
 * @Desc:
 */

package cryptUtils

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	Algorithm = otp.AlgorithmSHA1
)

func GenerateTotp(account string, issuer string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
		Algorithm:   Algorithm,
	})
	return key, err
}

func CheckTotp(totpUrl string, code string) bool {
	key, err := otp.NewKeyFromURL(totpUrl)
	if err != nil {
		return false
	}
	return totp.Validate(code, key.Secret())

}
