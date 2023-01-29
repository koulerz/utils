package kfile

import (
	"os"
)

// IsPathExist 判断路径是否存在
func IsPathExist(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsFileExists 判断文件是否存在
func IsFileExists(path string) bool {
	return IsFile(path)
}

// IsDirExist 判断目录是否存在
func IsDirExist(path string) bool {
	return IsDir(path)
}

// IsDir 判断是否为目录
func IsDir(path string) bool {
	if path == "" {
		return false
	}
	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile 判断是否为文件
func IsFile(path string) bool {
	if path == "" {
		return false
	}
	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}
