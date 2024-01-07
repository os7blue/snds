package util

import (
	"crypto/rand"
	"encoding/hex"
)

type globalUtil struct {
}

func (globalUtil) GenerateRandomHex(length int) (string, error) {
	// 生成指定长度的随机字节串
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// 将字节串转换为16进制字符串
	hexString := hex.EncodeToString(bytes)
	return hexString, nil
}
