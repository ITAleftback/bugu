/**
 * @Author: Anpw
 * @Description:
 * @File:  md5
 * @Version: 1.0.0
 * @Date: 2021/5/29 15:48
 */

package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
 
