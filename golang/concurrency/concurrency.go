package concurrency

import (
	"fmt"
)

func nFibonacciRec(n float64) float64 {
	if n <= 1 {
		return n
	}
	return nFibonacciRec(n-1) + nFibonacciRec(n-2)
}

func nFibonacciRecConc(n float64) float64 {

	c1 := make(chan float64)
	c2 := make(chan float64)

	var answer float64
	if n <= 1 {
		answer = n
	} else {
		go func() {
			c2 <- nFibonacciRecConc(n - 1)
		}()
		go func() {
			c1 <- nFibonacciRecConc(n - 2)
		}()

		answer = <-c2 + <-c1
	}
	close(c1)
	close(c2)

	return answer
}

type goRoutineManager struct {
	goRoutineCnt chan bool
}

func (g *goRoutineManager) Run(f func()) {
	select {
	case g.goRoutineCnt <- true:
		go func() {
			f()
			<-g.goRoutineCnt
		}()
	default:
		f()
	}
}

func NewGoRoutineManager(goRoutineLimit int) *goRoutineManager {
	return &goRoutineManager{
		goRoutineCnt: make(chan bool, goRoutineLimit),
	}
}

func nFibonacciRecConcManager(n float64, gm *goRoutineManager) float64 {

	c1 := make(chan float64, 1)
	c2 := make(chan float64, 1)

	var answer float64
	if n <= 1 {
		answer = n
	} else {

		gm.Run(func() {
			c2 <- nFibonacciRecConcManager(n-1, gm)
		})

		gm.Run(func() {
			c1 <- nFibonacciRecConcManager(n-2, gm)
		})

		answer = <-c2 + <-c1
	}
	close(c1)
	close(c2)

	return answer
}

func Run() {
	n := float64(10)

	fibonacciRec := nFibonacciRec(n)
	fmt.Println("nFibonacciRec", fibonacciRec)

	fibonacciRecConc := nFibonacciRecConc(n)
	fmt.Println("nFibonacciRecConc", fibonacciRecConc)

	gm := NewGoRoutineManager(100)
	fibonacciRecConcManager := nFibonacciRecConcManager(n, gm)
	fmt.Println("nFibonacciRecConcManager", fibonacciRecConcManager)

}
