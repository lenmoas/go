package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

//func main() {
//	codeCount := map[rune]int{}
//	count("가나다나", codeCount)
//	for _, key := range []rune{'가', '나', '다'} {
//		fmt.Println(string(key), codeCount[key])
//	}
//}

func main() {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}
