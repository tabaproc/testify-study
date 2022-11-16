package main

import (
	"fmt"
	"math"
)

// 二乗を返す
func square(num int) int {
	return int(math.Pow(float64(num), 2))
}

func main() {
	fmt.Printf("二乗した結果: %d\n", square(6))
}
