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

	matrix := make([][]int, len(input[0]))

	// fill the matrix by rotating the input
	// Exemple:
	// Input:
	// [
	//     "01011",
	//     "01001"
	// ]
	// Matrix:
	// [
	//     00,
	//     11
	//     00,
	//     10,
	//     11
	// ]
	for i := 0; i < len(input[0]); i++ {
		matrix[i] = make([]int, 0)
		for j := 0; j < len(input); j++ {
			n, _ := strconv.Atoi(string(input[j][i]))
			matrix[i] = append(matrix[i], n)
		}
	}

	var gamma, epsilon string

	for _, i := range matrix {
		var zero, one int
		for _, j := range i {
			switch j {
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
