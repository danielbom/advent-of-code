package utils

import (
	"fmt"
	"time"
)

var showTime bool = true

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TimeIt(name, format string, f func() any) {
	if showTime {
		start := time.Now()
		result := f()
		elapsed := time.Since(start)
		fmt.Printf("%s "+format+" [%v]\n", name, result, elapsed)
	} else {
		result := f()
		fmt.Printf("%s "+format+"\n", name, result)
	}
}

func DisableTime() {
	showTime = false
}
