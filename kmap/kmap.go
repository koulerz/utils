package kmap

// Keys 返回 keys 切片
func Keys[K comparable, T any](m map[K]T) []K {
	var res = make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// Values 返回 values 切片
func Values[K comparable, T any](m map[K]T) []T {
	var res = make([]T, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
