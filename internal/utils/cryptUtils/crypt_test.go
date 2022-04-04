/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:39
 * @Desc:
 */

package cryptUtils

import (
	"testing"
)

func TestHmacSha256(t *testing.T) {
	sha256 := HmacSha256("aaaaa", "aa")
	if sha256 != "574b1d75c67447d55da4b88ba283d866ddb267143575778d15c936993e69c98d" {
		t.Fatalf("calc hmac sha256 error")
	}
}

func TestOtp(t *testing.T) {
	totp, err := GenerateTotp("test-account", "issuer")
	if err != nil {
		t.Fatalf("generate totp error:  %s", err.Error())
	}
	code, err := generateTotpCode(totp.Secret())
	if err != nil {
		t.Fatalf("generate totp code error:  %s", err.Error())
	}
	if !ValidateTotp(totp.String(), code) {
		t.Fatalf("validate totp error")
	}
}
