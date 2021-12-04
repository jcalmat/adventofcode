package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movement string

const (
	Forward Movement = "forward"
	Down    Movement = "down"
	Up      Movement = "up"
)

func main() {
	content, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("failed to open input.json")
		return
	}

	steps := make([]map[Movement]int, 0)
	err = json.Unmarshal(content, &steps)

	var depth, distance, aim int
	for _, step := range steps {
		for k, v := range step {
			switch k {
			case Forward:
				distance += v
				depth += aim * v
			case Down:
				aim += v
			case Up:
				aim -= v
			}
		}
	}

	fmt.Println(depth * distance)
}
