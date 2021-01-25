package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	fishData := []float64{2, 3, 3, 4, 4, 4, 4, 5, 5, 6}
	fmt.Printf("合計: %f\n", sum(fishData))
	fmt.Printf("平均: %f\n", stat.Mean(fishData, nil))
	fmt.Printf("不偏分散: %f\n", stat.Variance(fishData, nil))
	fmt.Printf("不偏標準偏差: %f\n", stat.StdDev(fishData, nil))
}

func sum(xs []float64) float64 {
	s := 0.0
	for _, x := range xs {
		s += x
	}
	return s
}
