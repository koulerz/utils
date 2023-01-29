package kslice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReverse(t *testing.T) {
	// test odd length
	var s1 = []string{"a", "c", "z", "b", "f"}
	var e1 = []string{"f", "b", "z", "c", "a"}
	var res1 = Reverse(s1)
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of Reverse(%#v) should be %v but %v", s1, e1, res1)
	}

	// test even length
	var s2 = []string{"a", "c", "z", "b", "f", "d"}
	var e2 = []string{"d", "f", "b", "z", "c", "a"}
	var res2 = Reverse(s2)
	if !reflect.DeepEqual(res2, e2) {
		t.Errorf("the result of Reverse(%#v) should be %v but %v", s2, e2, res2)
	}

	// test slice capacity
	var s3 = make([]string, 6, 10)
	s3 = []string{"a", "c", "z", "b", "f", "d"}
	var e3 = []string{"d", "f", "b", "z", "c", "a"}
	var res3 = Reverse(s3)
	if !reflect.DeepEqual(res3, e3) {
		t.Errorf("the result of Reverse(%#v) should be %v but %v", s3, e3, res3)
	}
}

func TestDeleteKeys(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = []int{1, 3, 4}
	var res1 = DeleteKeys(s1, []int{0, 2, 5})
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of DeleteKeys() should be %v but %v", e1, res1)
	}
}

func TestDeleteValues(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5, 1, 3, 5}
	var e1 = []int{2, 3, 4, 3}
	var res1 = DeleteValues(s1, []int{0, 1, 5})
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of DeleteValues() should be %v but %v", e1, res1)
	}
}

func TestUnique(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5, 1, 3, 5}
	var e1 = []int{0, 1, 2, 3, 4, 5}
	var res1 = Unique(s1)
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of Unique() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = []int{0, 1, 2, 3, 4, 5}
	var res2 = Unique(s2)
	if !reflect.DeepEqual(res2, e2) {
		t.Errorf("the result of Unique() should be %v but %v", e2, res2)
	}
}

func TestFilter(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5, 1, 3, 5}
	var e1 = []int{2, 3, 4, 3}
	var res1 = Filter(s1, func(i int) bool {
		if i == 0 || i == 1 || i == 5 {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of Filter() should be %v but %v", e1, res1)
	}

	var s2 = []string{"a", "b", "c", "d", "e", "a"}
	var e2 = []string{"b", "d"}
	var res2 = Filter(s2, func(s string) bool {
		if s == "a" || s == "c" || s == "e" {
			return false
		}
		return true
	})
	if !reflect.DeepEqual(res2, e2) {
		t.Errorf("the result of Filter() should be %v but %v", e2, res2)
	}
}

func TestMap(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5, 1, 3, 5}
	var e1 = []int{10, 11, 12, 13, 14, 15, 11, 13, 15}
	var res1 = Map(s1, func(i int) int {
		return i + 10
	})
	if !reflect.DeepEqual(res1, e1) {
		t.Errorf("the result of Map() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5, 1, 3, 5}
	var e2 = []string{"0", "1", "2", "3", "4", "5", "1", "3", "5"}
	var res2 = Map(s2, func(i int) string {
		return strconv.Itoa(i)
	})
	if !reflect.DeepEqual(res2, e2) {
		t.Errorf("the result of Map() should be %v but %v", e2, res2)
	}
}

func TestEqual(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = true
	var res1 = Equal(s1, []int{0, 1, 2, 3, 4, 5})
	if res1 != e1 {
		t.Errorf("the result of Equal() should be %v but %v", e1, res1)
	}

	var s3 = make([]int, 0, 10)
	s3 = []int{0, 1, 2, 3, 4, 5}
	var e3 = true
	var res3 = Equal(s3, []int{0, 1, 2, 3, 4, 5})
	if res3 != e3 {
		t.Errorf("the result of Equal() should be %v but %v", e3, res3)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = false
	var res2 = Equal(s2, []int{0, 1, 2, 3, 5, 4})
	if res2 != e2 {
		t.Errorf("the result of Equal() should be %v but %v", e2, res2)
	}
}

func TestEqualFunc(t *testing.T) {
	var s1 = []func() int{func() int { return 1 }}
	var s2 = []func() string{func() string { return "1" }}
	var e1 = true
	var res1 = EqualFunc(s1, s2, func(v1 func() int, v2 func() string) bool {
		return strconv.Itoa(v1()) == v2()
	})
	if res1 != e1 {
		t.Errorf("the result of EqualFunc() should be %v but %v", e1, res1)
	}

	var s3 = []int{1, 2, 3, 4}
	var e3 = true
	var res3 = EqualFunc(s3, []int{6, 7, 8, 9}, func(v1, v2 int) bool {
		return (v1 > 5) == (v2 < 5)
	})
	if res3 != e3 {
		t.Errorf("the result of EqualFunc() should be %v but %v", e3, res3)
	}
}

func TestContains(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = true
	var res1 = Contains(s1, 0)
	if res1 != e1 {
		t.Errorf("the result of Contains() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = false
	var res2 = Contains(s2, 6)
	if res2 != e2 {
		t.Errorf("the result of Contains() should be %v but %v", e2, res2)
	}

	var s3 = []string{"a", "b", "c"}
	var e3 = true
	var res3 = Contains(s3, "a")
	if res3 != e3 {
		t.Errorf("the result of Contains() should be %v but %v", e3, res3)
	}
}

func TestContainsFunc(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = true
	var res1 = ContainsFunc(s1, func(i int) bool {
		return i < 5
	})
	if res1 != e1 {
		t.Errorf("the result of ContainsFunc() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = false
	var res2 = ContainsFunc(s2, func(i int) bool {
		return i > 5
	})
	if res2 != e2 {
		t.Errorf("the result of ContainsFunc() should be %v but %v", e2, res2)
	}
}

func TestContainsAll(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = true
	var res1 = ContainsAll(s1, []int{1, 2, 3})
	if res1 != e1 {
		t.Errorf("the result of ContainsAll() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = false
	var res2 = ContainsAll(s2, []int{4, 5, 6})
	if res2 != e2 {
		t.Errorf("the result of ContainsAll() should be %v but %v", e2, res2)
	}
}

func TestUnion(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = []int{0, 1, 2, 3, 4, 5}
	var res1 = Union(s1, []int{1, 2, 3})
	if len(res1) != len(e1) {
		t.Errorf("the result of Union() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = []int{0, 1, 2, 3, 4, 5, 6}
	var res2 = Union(s2, []int{4, 5, 6})
	if len(res2) != len(e2) {
		t.Errorf("the result of Union() should be %v but %v", e2, res2)
	}
}

func TestIntersect(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = []int{0, 3}
	var res1 = Intersect(s1, []int{0, 3, 6, 7})
	if len(res1) != len(e1) {
		t.Errorf("the result of Intersect() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = []int{}
	var res2 = Intersect(s2, []int{6, 7})
	if len(res2) != len(e2) {
		t.Errorf("the result of Intersect() should be %v but %v", e2, res2)
	}
}

func TestDifference(t *testing.T) {
	var s1 = []int{0, 1, 2, 3, 4, 5}
	var e1 = []int{1, 2, 4, 5, 6, 7}
	var res1 = Difference(s1, []int{0, 3, 6, 7})
	if len(res1) != len(e1) {
		t.Errorf("the result of Difference() should be %v but %v", e1, res1)
	}

	var s2 = []int{0, 1, 2, 3, 4, 5}
	var e2 = []int{}
	var res2 = Difference(s2, []int{0, 1, 2, 3, 4, 5})
	if len(res2) != len(e2) {
		t.Errorf("the result of Difference() should be %v but %v", e2, res2)
	}
}
