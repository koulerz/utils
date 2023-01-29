# utils
utils for go

- kstring 字符串相关
- kslice 切片相关
- kmap 映射相关
- kfile 文件操作相关
- kcommand 命令行相关
- kvalid 验证相关
- krand 随机数相关

# details
- kstrings
- kslice
  - InSlice() bool
  - DeleteByKey() []any
  - DeleteByValue() []any
- kmap
- kfile
  - IsFile() bool
  - IsDir() bool
  - IsFileExist() bool
  - IsDirExist() bool
  - IsPathExist() bool
  - CreateFile() *File
  - CreateDir() bool
  - DeleteFile() bool
  - DeleteDir() bool
  - GetDirs() []string
  - GetFiles() []string
  - GetAllDirs() []string
  - GetAllFiles() []string
- kconsole
  - IsCommandExist(name string) bool
- kvalidator
  - IsString() bool
  - IsInt() bool
  - IsMobile() bool
- krand
  - GetRandInt() int
  - GetRandString() string

# other
- [time](https://github.com/golang-module/carbon)
- [encode/encry](https://github.com/golang-module/dongle)

# 参考
- [golang-slices](https://pkg.go.dev/golang.org/x/exp/slices)
