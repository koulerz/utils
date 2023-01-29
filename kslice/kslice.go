package kslice

// Reverse 逆转切片的值
func Reverse[T any](s []T) []T {
	var length = len(s)
	var loop int
	if length%2 == 0 {
		loop = (length / 2) - 1
	} else {
		loop = ((length + 1) / 2) - 1
	}
	for i, j := 0, length; i <= loop; i++ {
		j--
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// DeleteKeys 删除指定 key 的元素
func DeleteKeys[T any](s []T, ks []int) []T {
	var res = make([]T, 0, len(s)-len(ks))
	for k, v := range s {
		if !Contains(ks, k) {
			res = append(res, v)
		}
	}
	return res
}

// DeleteValues 删除指定 values 的元素
func DeleteValues[T comparable](s []T, vs []T) []T {
	var res = make([]T, 0, len(s)-len(vs))
	for _, v := range s {
		if !Contains(vs, v) {
			res = append(res, v)
		}
	}
	return res
}

// Unique 移除重复的元素
func Unique[T comparable](s []T) []T {
	var record = map[T]int{}
	for _, v := range s {
		record[v]++
	}
	var res = make([]T, 0, len(s))
	for k := range record {
		res = append(res, k)
	}
	return res
}

// Filter 根据自定义方法过滤切片元素
func Filter[T any](s []T, f func(T) bool) []T {
	var res = make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// Map 将自定义方法应用于每个切片元素并返回
func Map[E, T any](s []E, f func(E) T) []T {
	var res = make([]T, 0, len(s))
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

// Equal 判断两个切片是否相等
func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// Equal 判断两个切片是否相等
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Contains 判断切片是否包含指定元素
func Contains[E comparable](s []E, v E) bool {
	for _, vs := range s {
		if v == vs {
			return true
		}
	}
	return false
}

// ContainsFunc 判断切片是否包含指定元素
func ContainsFunc[E any](s []E, f func(E) bool) bool {
	for _, vs := range s {
		if f(vs) {
			return true
		}
	}
	return false
}

// ContainsAll 判断 s1 切片是否完全包含 s2 切片
func ContainsAll[T comparable](s1, s2 []T) bool {
	return len(Intersect(s1, s2)) == len(Unique(s2))
}

// Union 返回切片并集
func Union[T comparable](s1, s2 []T) []T {
	return Unique(append(s1, s2...))
}

// Intersect 返回切片交集
func Intersect[T comparable](s1, s2 []T) []T {
	s1 = Unique(s1)
	s2 = Unique(s2)

	var length = len(s1)
	if len(s1) > len(s2) {
		length = len(s2)
	}

	var record = map[T]int{}
	for _, v := range s1 {
		record[v]++
	}
	for _, v := range s2 {
		record[v]++
	}

	var res = make([]T, 0, length)
	for k, v := range record {
		if v == 2 {
			res = append(res, k)
		}
	}

	return res
}

// Difference 返回切片差集
func Difference[T comparable](s1, s2 []T) []T {
	s1 = Unique(s1)
	s2 = Unique(s2)

	var record = map[T]int{}
	for _, v := range s1 {
		record[v]++
	}
	for _, v := range s2 {
		record[v]++
	}

	var res = make([]T, 0, len(s1)+len(s2))
	for k, v := range record {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}
