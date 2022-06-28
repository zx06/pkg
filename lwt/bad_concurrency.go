package lwt

import "sync"


// nolint:govet
// 取决于每个goroutine的执行顺序, 每个append进去的i可能为0-4之间的任意一个值
// BadLoopDataRace
func BadLoopDataRace() []int {
	var (
		wg  sync.WaitGroup
		res []int
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res = append(res, i)
		}()
	}
	wg.Wait()
	return res
}

// 方案1, 作为函数的参数
// BadLoopDataRaceFix1
func BadLoopDataRaceFix1() []int {
	var (
		wg  sync.WaitGroup
		res []int
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res = append(res, i)
		}(i)
	}
	wg.Wait()
	return res
}

// 方案2, 复制一份再用
// BadLoopDataRaceFix2
func BadLoopDataRaceFix2() []int {
	var (
		wg  sync.WaitGroup
		res []int
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		ic := i
		go func() {
			defer wg.Done()
			res = append(res, ic)
		}()
	}
	wg.Wait()
	return res
}
