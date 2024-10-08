package kfile

import (
	"bytes"
	"os"
	"path/filepath"
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
// Deprecated: Use IsFileExist instead
func IsFileExists(path string) bool {
	return IsFile(path)
}

// IsFileExist 判断文件是否存在
func IsFileExist(path string) bool {
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

// ReadLines 读取文件行
func ReadLines(name string) ([][]byte, error) {
	cont, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return bytes.Split(cont, []byte("\n")), nil
}

// WriteFile 向文件写入内容
//
// 若文件不存在，将以 0666 权限创建新文件
func WriteFile(name string, data []byte) error {
	err := CreateDir(filepath.Dir(name))
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, 0666)
}

// AppendFile 向文件追加内容
//
// 若文件不存在，将以 0666 权限创建新文件
// 若目录不存在，将以 0755 权限创建目录
func AppendFile(name string, data []byte) error {
	err := CreateDir(filepath.Dir(name))
	if err != nil {
		return err
	}
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

// CreateFile 创建文件
//
// 若文件已存在则覆盖已存在的文件
// 若目录不存在则以 0755 权限创建
func CreateFile(path string) error {
	err := CreateDir(filepath.Dir(path))
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

// CreateDir 创建目录
//
// 若路径不存在则以 0755 权限创建
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}
