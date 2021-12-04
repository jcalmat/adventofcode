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

	var depth, distance int
	for _, step := range steps {
		for k, v := range step {
			switch k {
			case Forward:
				distance += v
			case Down:
				depth += v
			case Up:
				depth -= v
			}
		}
	}

	fmt.Println(depth * distance)
}
