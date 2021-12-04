package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("failed to open input.json")
		return
	}

	input := make([]string, 0)
	err = json.Unmarshal(content, &input)
	if err != nil {
		fmt.Println("failed to unmarshal input.json")
		return
	}

	sumMap := make([]int, len(input[0]))

	// fill the sumMap
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			n, _ := strconv.Atoi(string(input[j][i]))
			sumMap[i] += n
		}
	}

	var gamma, epsilon string

	for _, v := range sumMap {
		if v < len(input)/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	a, _ := strconv.ParseInt(gamma, 2, 64)
	b, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(a * b)
}
