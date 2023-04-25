package slice

import (
	"hash/fnv"
	"math/rand"
	"sort"
	"strings"
)

func Union(a, b []string) []string {
	set := make(map[string]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := set[v]; !ok {
			a = append(a, v)
		}
	}
	return a
}

func DifferenceRemovedAdded(a, b []string) (removed []string, added []string) {
	var amap = map[string]struct{}{}
	var bmap = map[string]struct{}{}

	for _, item := range a {
		amap[item] = struct{}{}
	}

	for _, item := range b {
		if _, exists := amap[item]; !exists {
			added = append(added, item)
		}
		bmap[item] = struct{}{}
	}

	for _, item := range a {
		if _, exists := bmap[item]; !exists {
			removed = append(removed, item)
		}
	}
	return
}

func FindPos[T comparable](s []T, v T) int {
	for i, sv := range s {
		if sv == v {
			return i
		}
	}
	return -1
}

func Find[T comparable](s []T, cond func(T) bool) int {
	for i, sv := range s {
		if cond(sv) {
			return i
		}
	}
	return -1
}

// Difference returns the elements in `a` that aren't in `b`.
func Difference(a, b []string) []string {
	var diff = make([]string, 0, len(a))
	for _, a1 := range a {
		if FindPos(b, a1) == -1 {
			diff = append(diff, a1)
		}
	}
	return diff
}

func Insert[T any](s []T, pos int, v ...T) []T {
	if len(s) <= pos {
		return append(s, v...)
	}
	if pos == 0 {
		return append(v, s[pos:]...)
	}
	return append(s[:pos], append(v, s[pos:]...)...)
}

// Remove reuses provided slice capacity. Provided s slice should not be used after without reassigning to the func return!
func Remove[T comparable](s []T, v T) []T {
	var n int
	for _, x := range s {
		if x != v {
			s[n] = x
			n++
		}
	}
	return s[:n]
}

// RemoveIndex reuses provided slice capacity. Provided s slice should not be used after without reassigning to the func return!
func RemoveIndex[T any](s []T, idx int) []T {
	var n int
	for i, x := range s {
		if i != idx {
			s[n] = x
			n++
		}
	}
	return s[:n]
}

func Filter[T any](vals []T, cond func(T) bool) []T {
	var result = make([]T, 0, len(vals))
	for i := range vals {
		if cond(vals[i]) {
			result = append(result, vals[i])
		}
	}
	return result
}

func FilterMut[T any](vals []T, cond func(T) bool) []T {
	result := vals[:0]
	for i := range vals {
		if cond(vals[i]) {
			result = append(result, vals[i])
		}
	}
	return result
}

func GetRandomString(s []string, seed string) string {
	rand.Seed(int64(hash(seed)))
	return s[rand.Intn(len(s))]
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func SortedEquals(s1, s2 []string) bool {
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

func UnsortedEquals(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Sorted := make([]string, len(s1))
	s2Sorted := make([]string, len(s2))
	copy(s1Sorted, s1)
	copy(s2Sorted, s2)
	sort.Strings(s1Sorted)
	sort.Strings(s2Sorted)

	return SortedEquals(s1Sorted, s2Sorted)
}

func HasPrefix(value, prefix []string) bool {
	if len(value) < len(prefix) {
		return false
	}
	for i, p := range prefix {
		if value[i] != p {
			return false
		}
	}
	return true
}

func Copy(s []string) []string {
	res := make([]string, len(s))
	copy(res, s)
	return res
}

func Intersection(a, b []string) (res []string) {
	sort.Strings(a)
	sort.Strings(b)
	aIdx := 0
	bIdx := 0
	for aIdx < len(a) && bIdx < len(b) {
		cmp := strings.Compare(a[aIdx], b[bIdx])
		switch cmp {
		case 0:
			res = append(res, a[aIdx])
			aIdx++
			bIdx++
		case -1:
			aIdx++
		case 1:
			bIdx++
		}
	}
	return
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Subtract[T comparable](a, b []T) []T {
	mb := make(map[T]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []T
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
