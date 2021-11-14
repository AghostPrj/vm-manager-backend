/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2021/11/14 13:12
 * @Desc:
 */

package cryptUtils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(data string, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
