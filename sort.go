package timesort

import (
	"sync"
	"time"
)

// Sort sorts the given array of integers in ascending order.
// It takes an additional parameter k, which is the time factor.
// Bigger k means faster sorting.
// But for big k, the result is not guaranteed to be deterministic.
func Sort(arr []int, k int) []int {
	m := sync.Mutex{}
	res := make([]int, 0, len(arr))

	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v int) {
			time.Sleep(time.Duration(v) * (time.Millisecond / time.Duration(k)))
			m.Lock()
			res = append(res, v)
			m.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
	return res
}
