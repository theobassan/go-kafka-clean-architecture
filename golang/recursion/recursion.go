package recursion

import "fmt"

func factorialLoop(num int) int {
	result := 1
	for ; num > 0; num-- {
		result *= num
	}
	return result
}

func factorialRecursive(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorialRecursive(num-1)
}

func factorialHO(num int, next func(int)) {
	if num == 0 {
		next(1)
	} else {
		factorialHO(num-1, func(result int) {
			next(num * result)
		})
	}
}

func factorialCR(num int) func(next func(int)) {
	return func(next func(int)) {
		if num == 0 {
			next(1)
		} else {
			factorialCR(num - 1)(func(result int) {
				next(num * result)
			})
		}
	}
}

func factorialCR2(next func(int)) func(num int) {
	return func(num int) {
		if num == 0 {
			next(1)
		} else {
			factorialCR(num - 1)(func(result int) {
				next(num * result)
			})
		}
	}
}

func factorialCR3() func(next func(int)) func(num int) {
	return func(next func(int)) func(num int) {
		return func(num int) {
			if num == 0 {
				next(1)
			} else {
				factorialCR(num - 1)(func(result int) {
					next(num * result)
				})
			}
		}
	}
}

func printResult(name string, result int) {
	fmt.Println(name, result)
}

func Run() {
	factorialLoopResult := factorialLoop(4)
	printResult("factorialLoop", factorialLoopResult)

	factorialRecursiveResult := factorialRecursive(4)
	printResult("factorialRecursive", factorialRecursiveResult)

	factorialHO(4, func(result int) {
		printResult("factorialHO", result)
	})

	printCRResultFunc := func(result int) {
		printResult("factorialCR", result)
	}
	factorialCRResult := factorialCR(4)
	factorialCRResult(printCRResultFunc)

	printCR2ResultFunc := func(result int) {
		printResult("factorialCR2", result)
	}
	factorialCR2Result := factorialCR2(printCR2ResultFunc)
	factorialCR2Result(4)

	printCR3ResultFunc := func(result int) {
		printResult("factorialCR3", result)
	}
	factorialCR3Result := factorialCR3()(printCR3ResultFunc)
	factorialCR3Result(4)
}
