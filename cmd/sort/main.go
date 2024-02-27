package main

import (
	"fmt"
	"time"

	timesort "github.com/trueaniki/time-sort"
)

func main() {
	start := time.Now()
	arr := []int{3, 2, 1}
	fmt.Println(timesort.Sort(arr, 70))
	fmt.Println(time.Since(start))
}
