package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	oxygenEngine := Engine{
		PosValue: "1",
		NegValue: "0",
	}

	co2Engine := Engine{
		PosValue: "0",
		NegValue: "1",
	}

	oxygen := oxygenEngine.computeRating(input)
	co2 := co2Engine.computeRating(input)

	a, _ := strconv.ParseInt(oxygen, 2, 64)
	b, _ := strconv.ParseInt(co2, 2, 64)
	fmt.Printf("With an oxygen generator rating of %d and a CO2 scrubber rating of %d, we can assume that the submarine's life support rating is %d.\n", a, b, a*b)
}

func computeSumMap(input []string) []int {
	sumMap := make([]int, len(input[0]))

	// fill the sumMap
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input); j++ {
			n, _ := strconv.Atoi(string(input[j][i]))
			sumMap[i] += n
		}
	}

	return sumMap
}

type Engine struct {
	PosValue string
	NegValue string
}

func (e Engine) computeRating(input []string) string {
	var ret string

	sumMap := computeSumMap(input)
	l := len(sumMap)
	for i := 0; i < l; i++ {
		sumMap := computeSumMap(input)

		if float32(sumMap[i]) >= float32(len(input))/2 {
			ret += e.PosValue
		} else {
			ret += e.NegValue
		}

		tmpinput := make([]string, 0)

		for _, i := range input {
			if strings.HasPrefix(i, ret) {
				tmpinput = append(tmpinput, i)
			}
		}

		if len(tmpinput) == 1 {
			ret = tmpinput[0]
			break
		}

		input = tmpinput
	}

	return ret
}
