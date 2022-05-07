package goroutine

import (
	"fmt"
	"sync"
)

func nFibonacciLoop() func(next func(float64)) func(n float64) {
	return func(next func(float64)) func(n float64) {
		return func(n float64) {
			f := make([]float64, int(n+1), int(n+2))
			if n < 2 {
				f = f[0:2]
			}
			f[0] = 0
			f[1] = 1
			for i := 2; i <= int(n); i++ {
				f[i] = f[i-1] + f[i-2]
			}
			next(f[int(n)])
		}
	}
}

func nFibonacciRec() func(next func(float64)) func(n float64) {
	return func(next func(float64)) func(n float64) {
		return func(n float64) {
			if n <= 1 {
				next(n)
			}
			nFibonacciRec()(func(result float64) {
				nFibonacciRec()(func(result2 float64) {
					next(result + result2)
				})(n - 2)
			})(n - 1)
		}
	}
}

func printResult(name string, result float64) {
	fmt.Println(name, result)
}

func Run() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	n := float64(10)

	fibonacciLoopResultFunc := func(result float64) {
		defer wg.Done()
		printResult("nFibonacciLoop", result)
	}
	fibonacciLoop := nFibonacciLoop()(fibonacciLoopResultFunc)

	fibonacciRecResultFunc := func(result float64) {
		defer wg.Done()
		printResult("nFibonacciRec", result)
	}
	fibonacciRec := nFibonacciRec()(fibonacciRecResultFunc)

	go fibonacciLoop(n)
	go fibonacciRec(n)

	wg.Wait()
	fmt.Println("Done")
}
