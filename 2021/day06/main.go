package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	days := 256
	ages, err := getInput("2021/day06/input.csv")
	if err != nil {
		panic(err)
	}
	var school []*Lanternfish
	for _, age := range ages {
		fish := Lanternfish{age: age, daysLeft: days}
		school = append(school, &fish)
	}
	runningFishTally := 0
	// fishRecord is a lookup table of Lanternfish
	// we've calculated the progeny outcome for
	fishRecord := make(map[Lanternfish]int)
	for _, fish := range school {
		progeny, ok := fishRecord[*fish]
		if !ok {
			progeny = fish.getProgenyCount(fishRecord)
		}
		runningFishTally += progeny
	}
	fmt.Println(runningFishTally)
	fmt.Println("Time:", time.Since(start))
}

type Lanternfish struct {
	age int
	daysLeft int
}

// advanceTimer progresses time for this fish until it
// either produces another fish (resetting this fish's timer
// and returning the new fish) or remaining days run out,
// whichever is first
func (f *Lanternfish) advanceTimer() *Lanternfish {
	if f.daysLeft == 0 {
		return nil
	}
	if f.age >= f.daysLeft{
		f.age = f.age - f.daysLeft
		f.daysLeft = 0
		return nil
	}
	f.daysLeft = f.daysLeft - f.age - 1
	f.age = 6
	return &Lanternfish{
		age:      8,
		daysLeft: f.daysLeft,
	}
}

// getProgenyCount advances time by the number of days this Lanternfish has left
// and returns the total number of fish that will exist when daysLeft == 0
// each unique Lanternfish encountered by this algorithm will be added to the
// lookup table fishRecord for efficiency
func (f *Lanternfish) getProgenyCount(fishRecord map[Lanternfish]int) int {
	if f.daysLeft == 0 {
		return 1
	}
	if progeny, ok := fishRecord[*f]; ok {
		return progeny
	}
	// grab a copy because f.advanceTimer() mutates the fish before we store it
	fishSnapshot := *f
	if newFishy := f.advanceTimer(); newFishy == nil {
		progeny := f.getProgenyCount(fishRecord)
		fishRecord[fishSnapshot] = progeny
		return progeny
	} else {
		progeny := f.getProgenyCount(fishRecord) + newFishy.getProgenyCount(fishRecord)
		fishRecord[fishSnapshot] = progeny
		return progeny
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
