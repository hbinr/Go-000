package hash

import (
	"crypto/md5"
	"encoding/hex"

)

// keyMySecret 盐值
const keyMySecret = "user.service.sadasd1234nsdcui123"
// MD5String 密码加密
func MD5String(oPassword string) string {
	h := md5.New()
	h.Write([]byte(keyMySecret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
