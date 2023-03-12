package slicex

import (
	"strconv"

	"github.com/minlib/go-util/json"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Slice
func Slice[E any](s ...E) []E {
	return []E(s)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[E any](s []E, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return slices.Contains(s, v)
}

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[E comparable](s1, s2 []E) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Delete removes the elements s[i:i+1] from s, returning the modified slice
func Delete[S ~[]E, E any](s S, i int) S {
	return slices.Delete(s, i, i+1)
}

// Distinct 去重
func Distinct[S ~[]E, E comparable](s S) S {
	var res S
	if len(s) > 0 {
		m := map[E]struct{}{}
		for _, v := range s {
			if _, found := m[v]; !found {
				m[v] = struct{}{}
				res = append(res, v)
			}
		}
	}
	return res
}

// Subtract returns the elements in `s1` that aren't in `s2`
func Subtract[S ~[]E, E comparable](s1, s2 S) S {
	var s S
	if len(s1) > 0 {
		var mp = make(map[E]struct{}, len(s2))
		for _, v := range s2 {
			mp[v] = struct{}{}
		}
		for _, v := range s1 {
			if _, found := mp[v]; !found {
				s = append(s, v)
			}
		}
	}
	return s
}

// Intersect 返回交集并去重
func Intersect[S ~[]E, E comparable](s1, s2 S) S {
	var s S
	if len(s1) > 0 && len(s2) > 0 {
		var mp = make(map[E]struct{}, 0)
		for _, v := range s2 {
			mp[v] = struct{}{}
		}
		for _, v := range s1 {
			if _, found := mp[v]; found {
				s = append(s, v)
			}
		}
	}
	return Distinct(s)
}

// Union 返回并集并去重
func Union[S ~[]E, E comparable](s1, s2 S) S {
	s := append(s1, s2...)
	return Distinct(s)
}

// Sum 求和
func Sum[S ~[]E, E constraints.Ordered](s S) E {
	var e E
	for _, v := range s {
		e += v
	}
	return e
}

// EqualIgnoreOrder 先排序再比较，不同排序的切片对比返回true
func EqualIgnoreOrder[E constraints.Ordered](s1, s2 []E) bool {
	slices.Sort(s1)
	slices.Sort(s2)
	return Equal(s1, s2)
}

// IntToString 整型切片转为字符串切片
func IntToString[S ~[]E, E constraints.Integer](s S) []string {
	var res []string
	for i := range s {
		res = append(res, strconv.FormatInt(int64(s[i]), 10))
	}
	return res
}

// StringToInt 字符串切片转为整型切片，E为转换后的整型类型
func StringToInt[E constraints.Integer](s []string) ([]E, error) {
	var res []E
	for i := range s {
		if val, err := strconv.Atoi(s[i]); err != nil {
			return nil, err
		} else {
			res = append(res, E(val))
		}
	}
	return res, nil
}

// LongToInt64 convert long slice to int64 slice
func LongToInt64(s []*json.Long) []int64 {
	var rs []int64
	for _, v := range s {
		if v != nil {
			rs = append(rs, v.Int64)
		}
	}
	return rs
}

// Int64ToLong convert int64 slice to long slice
func Int64ToLong(s []int64) []*json.Long {
	return json.NewLongSlice(s)
}
