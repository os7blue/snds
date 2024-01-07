package util

import "path/filepath"

type fileUtil struct {
}

// GetFileNameWithoutExtension 获取文件名（不带后缀）
func (*fileUtil) GetFileNameWithoutExtension(filePath string) string {
	base := filepath.Base(filePath)
	// 使用filepath.Ext获取文件后缀，然后去掉后缀
	return base[:len(base)-len(filepath.Ext(base))]
}
