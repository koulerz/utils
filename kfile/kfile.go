package kfile

import (
	"bytes"
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

// ReadFile 读文件
func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// ReadLines 按行读取文件
func ReadLines(name string) ([][]byte, error) {
	cont, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return bytes.Split(cont, []byte("\n")), nil
}

// WriteFile 向文件写入内容
// 若文件不存在，将以 0666 权限创建新文件
func WriteFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0666)
}

// AppendFile 向文件追加内容
// 若文件不存在，将以 0666 权限创建新文件
func AppendFile(name string, data []byte) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
