package main

import (
	"strconv"
	"testing"
)

var (
	// for int type map look up faster slice when element size between 25 and 30 or higher
	sizeInt = 30
	// for string type map look up faster when element size between 8 and 9 or higher
	sizeStr = 9
)

func BenchmarkSearchSliceInt(b *testing.B) {
	// init slice
	sl := make([]int, sizeInt)
	for i := 0; i < sizeInt; i++ {
		sl[i] = i
	}
	n := sizeInt - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchSliceInt(n, sl)
	}
}

func BenchmarkSearchMapInt(b *testing.B) {
	// init map
	mp := make(map[int]struct{}, sizeInt)
	for i := 0; i < sizeInt; i++ {
		mp[i] = struct{}{}
	}
	n := sizeInt - 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMapInt(n, mp)
	}
}

func BenchmarkSearchSliceStr(b *testing.B) {
	// init slice
	sl := make([]string, sizeStr)
	for i := 0; i < sizeStr; i++ {
		sl[i] = strconv.Itoa(i)
	}
	s := strconv.Itoa(sizeStr - 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchSliceStr(s, sl)
	}
}

func BenchmarkSearchMapStr(b *testing.B) {
	// init map
	mp := make(map[string]struct{}, sizeStr)
	for i := 0; i < sizeStr; i++ {
		k := strconv.Itoa(i)
		mp[k] = struct{}{}
	}
	s := strconv.Itoa(sizeStr - 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchMapStr(s, mp)
	}
}

func searchSliceInt(n int, sl []int) bool {
	for _, v := range sl {
		if n == v {
			return true
		}
	}
	return false
}

func searchMapInt(n int, mp map[int]struct{}) bool {
	if _, ok := mp[n]; ok {
		return true
	}
	return false
}

func searchSliceStr(s string, sl []string) bool {
	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}

func searchMapStr(s string, mp map[string]struct{}) bool {
	if _, ok := mp[s]; ok {
		return true
	}
	return false
}
