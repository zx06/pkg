package lwt

import "sync"

// nolint:govet
// 取决于每个goroutine的执行顺序, 每个append进去的i可能为0-4之间的任意一个值
// 这里面还有另外一个问题, 并发读写res的时候, 可能会出现数据不一致的情况
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
		mu  sync.Mutex
		res []int
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			res = append(res, j)
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
		mu  sync.Mutex
		res []int
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		ic := i
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			res = append(res, ic)
		}()
	}
	wg.Wait()
	return res
}
