package kfile

import "os"

// IsPathExist 判断路径是否存在
func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsFileExist 判断文件是否存在
func IsFileExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return !s.IsDir()
}

// IsDirExist 判断目录是否存在
func IsDirExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return s.IsDir()
}
