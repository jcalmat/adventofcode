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

	matrix := make([][]int, len(input))

	// fill the matrix
	for i, l := range input {
		matrix[i] = make([]int, 0)
		for _, r := range l {

			n, _ := strconv.Atoi(string(r))
			matrix[i] = append(matrix[i], n)
		}
	}

	var gamma, epsilon string

	for i := 0; i < len(matrix[0]); i++ {
		var zero, one int
		for j := 0; j < len(matrix); j++ {
			switch matrix[j][i] {
			case 1:
				one++
			case 0:
				zero++
			default:
				fmt.Println("invalid payload")
				return
			}
		}
		if zero > one {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}

	for _, r := range gamma {
		switch r {
		case '0':
			epsilon += "1"
		case '1':
			epsilon += "0"
		}
	}

	a, _ := strconv.ParseInt(gamma, 2, 64)
	b, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(a * b)
}
