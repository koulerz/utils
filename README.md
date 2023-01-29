# utils

utils for go

- kstring 字符串相关
- kslice 切片相关
- kmap 映射相关
- kfile 文件操作相关
- kcmd 命令行相关
- kvalid 验证相关
- krand 随机数相关

# kstring

# kslice

- **Reverse[T any](s []T) []T** 逆转切片的值
- **DeleteKeys[T any](s []T, ks []int) []T** 删除指定 key 的元素
- **DeleteValues[T comparable](s []T, vs []T) []T** 删除指定 values 的元素
- **Unique[T comparable](s []T) []T** 移除重复的元素
- **Filter[T any](s []T, f func(T) bool) []T** 根据自定义方法过滤切片元素
- **Map[E, T any](s []E, f func(E) T) []T** 将自定义方法应用于每个切片元素并返回
- **Equal[E comparable](s1, s2 []E) bool** 判断两个切片是否相等
- **EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool** 判断两个切片是否相等
- **Contains[E comparable](s []E, v E) bool** 判断切片是否包含指定元素
- **ContainsFunc[E any](s []E, f func(E) bool) bool** 判断切片是否包含指定元素
- **ContainsAll[T comparable](s1, s2 []T) bool** 判断 s1 切片是否完全包含 s2 切片
- **Union[T comparable](s1, s2 []T) []T** 返回切片并集
- **Intersect[T comparable](s1, s2 []T) []T** 返回切片交集
- **Difference[T comparable](s1, s2 []T) []T** 返回切片差集

# kmap

- **Keys[K comparable, T any](m map[K]T) []K** 返回 keys 切片
- **Values[K comparable, T any](m map[K]T) []T** 返回 values 切片

# kfile

- **IsPathExist(path string) bool** 判断路径是否存在
- **IsFileExists(path string) bool** 判断文件是否存在
- **IsDirExist(path string) bool** 判断目录是否存在
- **IsDir(path string) bool** 判断是否为目录
- **IsFile(path string) bool** 判断是否为文件

# kcmd

# kvalid

- **IsValidUUID(s string) bool** 验证是否为有效的 UUID

# krand

- **RandStr(n int) string** 返回一个指定长度的随机字符串
- **RandAlphabetStr(n int) string** 返回一个指定长度的随机字母字符串
- **RandUppercaseStr(n int) string** 返回一个指定长度的随机大写字母字符串
- **RandLowercaseStr(n int) string** 返回一个指定长度的随机小写字母字符串

# TODO

- kfile
  - CreateFile() \*File
  - CreateDir() bool
  - DeleteFile() bool
  - DeleteDir() bool
  - GetDirs() []string
  - GetFiles() []string
  - GetAllDirs() []string
  - GetAllFiles() []string

# libs

- **date/time** https://github.com/golang-module/carbon
- **encode/encry** https://github.com/golang-module/dongle

# 参考

- [golang-slices](https://pkg.go.dev/golang.org/x/exp/slices)
