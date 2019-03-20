package gofast

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
)

func For(st, fi, in int, function interface{}) {
	var wg sync.WaitGroup

	threads := runtime.NumCPU()
	ops := int(math.Abs(float64(st-fi))) / in

	if ops <= 0 {
		return
	}

	if threads > ops {
		threads = ops
	}

	for t := 0; t < threads; t++ {
		wg.Add(1)

		start := fi / threads * t
		finish := fi / threads * (t + 1)

		if t == threads-1 {
			finish = fi
		}

		fmt.Println(start, finish)
		go func(s, f int) {
			defer wg.Done()
			for ; s < f; s += in {
				reflect.ValueOf(function).Call([]reflect.Value{
					reflect.ValueOf(s), reflect.ValueOf(f)})
			}
		}(start, finish)
	}

	wg.Wait()
}
