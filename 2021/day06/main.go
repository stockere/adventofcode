package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	ages, err := getInput("2021/day06/input.csv")
	if err != nil {
		panic(err)
	}
	var school []*Lanternfish
	for _, age := range ages {
		fish := Lanternfish{age: age}
		school = append(school, &fish)
	}
	days := 80
	for i := 1; i <= days; i++ {
		var newFish []*Lanternfish
		for _, fish := range school {
			baby := fish.passDay()
			if baby != nil {
				newFish = append(newFish, baby)
			}
		}
		school = append(school, newFish...)
	}
	fmt.Println("Number of fish after day", days, len(school))
}

type Lanternfish struct {
	age int
}

func (l *Lanternfish) passDay() *Lanternfish {
	if l.age > 0 {
		l.age = l.age - 1
		return nil
	} else {
		l.age = 6
		return &Lanternfish{age: 8}
	}
}

func getInput(filename string) ([]int, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var nums []int
	r := csv.NewReader(csvFile)
	for {
		ages, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, age := range ages {
			num, err := strconv.Atoi(age)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
	}
	return nums, nil
}
