package main

import "fmt"

type counterFunc func() int

func main() {
    goods := []string{"x", "y", "x", "x", "x", "y", "x"}
	xCountVal, yCountVal := 0, 0
	goodsCounterMap := map[string]counterFunc {
		"x": counter(),
		"y": counter(),
	}
	for _, good := range goods {
		if good == "x" {
			xCountVal = goodsCounterMap[good]()
		}
		if good == "y" {
			yCountVal = goodsCounterMap[good]()
		}
	}
	fmt.Println("xCountVal:", xCountVal)
	fmt.Println("yCountVal:", yCountVal)
}

func counter() counterFunc {
	x := 0
	return func() int {
		x += 1
		return x
	}
}