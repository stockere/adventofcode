package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	nums, max, err := getInput("2021/day07/input.txt")
	if err != nil {
		panic(err)
	}
	leastFuelCost := math.MaxInt
	var optimalPosition int
	for i := 0; i <= max; i++ {
		var fuelCost int
		for num, freq := range nums {
			var moveCost int
			if i == num {
				moveCost = 0
			} else {
				moveCost = int(math.Abs(float64(i - num)))
			}
			fuelCost += moveCost * freq
		}
		if fuelCost < leastFuelCost {
			leastFuelCost = fuelCost
			optimalPosition = i
		}
	}
	fmt.Printf("Cheapest move is spending %v fuel to move all crabs to position %v\n",
		leastFuelCost, optimalPosition)
}

func getInput(filename string) (map[int]int, int, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	nums := make(map[int]int)
	var max int
	r := csv.NewReader(csvFile)
	for {
		ages, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, 0, err
		}
		for _, age := range ages {
			num, err := strconv.Atoi(age)
			if err != nil {
				return nil, 0, err
			}
			if num > max {
				max = num
			}
			value, _ := nums[num]
			nums[num] = value + 1
		}
	}
	return nums, max, nil
}
