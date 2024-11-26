package utils

import (
	"fmt"
	"os"
	"time"
)

var showTime bool = true

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func ReadAllFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
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
