package prime

import (
	"sync"
	"time"
)

func PrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}

	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}

func GoPrimesAndSleep(n int, sleep time.Duration) []int {

	var wg sync.WaitGroup
	out := make(chan int)

	res := []int{}

	go func() {
		for k := 2; k < n; k++ {
			wg.Add(1)
			go func(kk int) {
				defer wg.Done()
				for i := 2; i < n; i++ {
					if kk%i == 0 {
						time.Sleep(sleep)
						if kk == i {
							out <- kk
						}
						break
					}
				}
			}(k)
		}
		wg.Wait()
		close(out)
	}()
	for val := range out {
		res = append(res, val)
	}

	return res
}
