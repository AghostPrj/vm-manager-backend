/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:39
 * @Desc:
 */

package cryptUtils

import (
	"github.com/AghostPrj/vm-manager-backend/internal/utils/cryptUtils"
	"testing"
)

func TestHmacSha256(t *testing.T) {
	sha256 := cryptUtils.HmacSha256("aaaaa", "aa")
	if sha256 != "574b1d75c67447d55da4b88ba283d866ddb267143575778d15c936993e69c98d" {
		t.Fatalf("calc hmac sha256 error")
	}
}
