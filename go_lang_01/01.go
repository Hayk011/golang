package main

import "fmt"

func main() {
	array := []float64{1, 2, 3, 4, 5, 6, 7}
	n := 3
	fmt.Print(sum(array))
	fmt.Print(evenOrOdd(n))
	fff := oddNumGenerator()
	fmt.Print(fff())
	fmt.Print(fibonachi(5))
	x := 5
	y := 2
	reverseVariables(&x, &y)
	fmt.Print("\n",x,y)

}

func sum(arr []float64) float64 {
	total := 0.0
	for _, item := range arr {
		total += item
	}
	return total
}

func evenOrOdd(n int) (int, bool) {
	half := n / 2
	if half%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

func oddNumGenerator() func() uint {
	num := uint(1)
	return func() (req uint) {
		req = num
		num += 2
		return
	}
}

func fibonachi(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonachi(n-1) + fibonachi(n-2)
	}
}

func reverseVariables (x *int, y *int ){
	z := 0
	z = *x
	*x = *y
	*y = z
}




