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

func GenerateTotp(account string, issuer string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
		Algorithm:   otp.AlgorithmSHA256,
	})
	return key, err
}
