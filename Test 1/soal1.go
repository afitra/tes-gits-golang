package main

import (
	"fmt"
	"strconv"
)

func getSequence(num int) string {
	var result = ""
	for i := 0; i < num; i++ {
		result += strconv.Itoa(((i * i) + i + 2) / 2)
		if i != num-1 {
			result += "-"
		}
	}

	return result
}

func main() {

	var result = getSequence(7)
	fmt.Println(result)
}
