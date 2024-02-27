package timesort_test

import (
	"fmt"
	"sort"
	"testing"

	timesort "github.com/trueaniki/time-sort"
)

func BenchmarkTimeSort(b *testing.B) {
	arr := []int{3, 2, 1}
	fmt.Println(timesort.Sort(arr, 10))
}

func BenchmarkRegularSort(b *testing.B) {
	arr := []int{3, 2, 1}
	sort.Ints(arr)
	fmt.Println(arr)
}
