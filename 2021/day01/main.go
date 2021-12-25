package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Part 1 small input (expect function to return 7)
	// depths := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	// Part 2 small input (expect function to return 5)
	// depths := []int{607,618,618,617,647,716,769,792}
	depths, err := csvToSlice("2021/day01/input.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Increasing pairs count:", getIncreaseCountPairs(depths))
	fmt.Println("Increasing trios count:", getIncreaseCountTrios(depths))
}

// getIncreaseCountPairs returns the count of how many times two adjacent numbers
// in the given slice are an "increasing pair" (meaning the second number
// in the pair is greater than the first number)
func getIncreaseCountPairs(measurements []int) int {
	increaseCount := 0
	for i := 1; i < len(measurements); i++ {
		if v := measurements[i] - measurements[i-1]; v > 0 {
			increaseCount++
		}
	}
	return increaseCount
}

// getIncreaseCountTrios returns the count of how many times three adjacent numbers
// in the given slice are an "increasing trio" (meaning the sum of the values in the
// second trio are greater than the sum of the values in the trio immediately
// preceding it)
func getIncreaseCountTrios(measurements []int) int {
	increaseCount := 0
	trio1 := measurements[0] + measurements[1] + measurements[2]
	for i := 3; i < len(measurements); i++ {
		trio2 := measurements[i] + measurements[i-1] + measurements[i-2]
		if trio2 > trio1 {
			increaseCount++
		}
		trio1 = trio2
	}
	return increaseCount
}

// csvToSlice reads a csv file and attempts to parse the first column
// into a slice of ints
func csvToSlice(filename string) ([]int, error) {
	// open the file
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// parse the file
	var nums []int
	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		num, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}