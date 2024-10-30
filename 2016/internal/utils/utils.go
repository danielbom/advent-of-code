package utils

import (
	"fmt"
	"time"
)

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func TimeIt(name, format string, f func() any) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	fmt.Printf("%s "+format+" [%v]\n", name, result, elapsed)
}
