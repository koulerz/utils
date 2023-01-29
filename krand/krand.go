package krand

import (
	"math/rand"
	"time"
)

// RandStr 返回一个指定长度的随机字符串
func RandStr(n int) string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	// 确保随机字符串第一位和最后一位不为 0
	for {
		if result[0] == '0' {
			result[0] = bytes[rand.Intn(len(bytes))]
		}

		if result[n-1] == '0' {
			result[n-1] = bytes[rand.Intn(len(bytes))]
		}

		if result[0] != '0' && result[n-1] != '0' {
			break
		}
	}

	return string(result)
}

// RandAlphabetStr 返回一个指定长度的随机字母字符串
func RandAlphabetStr(n int) string {
	bytes := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	return string(result)
}

// RandUppercaseStr 返回一个指定长度的随机大写字母字符串
func RandUppercaseStr(n int) string {
	bytes := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	return string(result)
}

// RandLowercaseStr 返回一个指定长度的随机小写字母字符串
func RandLowercaseStr(n int) string {
	bytes := []byte("abcdefghijklmnopqrstuvwxyz")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	return string(result)
}
