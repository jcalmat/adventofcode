package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	c := 0

	content, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("failed to open input.json")
		return
	}

	input := make([]int, 0)
	err = json.Unmarshal(content, &input)

	for i := range input {
		if len(input) == i+1 {
			break
		}
		if input[i] < input[i+1] {
			c++
		}
	}
	fmt.Println(c)
}
