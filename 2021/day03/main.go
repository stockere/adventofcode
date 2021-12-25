package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	freqs, err := countFrequencies("2021/day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(freqs)
	gamma, epsilon, err := convertToRates(freqs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Gamma:", gamma)
	fmt.Println("Epsilon", epsilon)
	powerConsumption, err := multiplyRates(gamma, epsilon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(powerConsumption)
}

func countFrequencies(filename string) (map[int]map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	freqs := make(map[int]map[string]int)
	var binary string
	for {
		_, err := fmt.Fscanln(file, &binary)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for i, c := range binary {
			// check if we have a map for this index and create it if not
			if _, ok := freqs[i]; !ok {
				freqs[i] = make(map[string]int)
			}
			// add a count of 1 for the character in this position
			freqs[i][string(c)] += 1
		}
	}
	return freqs, nil
}

func convertToRates(freqs map[int]map[string]int) (string, string, error) {
	gamma := strings.Builder{}
	epsilon := strings.Builder{}
	for i := 0; i < len(freqs); i++ {
		freq := freqs[i]
		zeros, _ := freq["0"]
		ones, _ := freq["1"]
		if zeros == ones {
			return "", "", fmt.Errorf("error: 0's and 1's have same frequency")
		} else if zeros > ones {
			gamma.WriteRune('0')
			epsilon.WriteRune('1')
		} else {
			gamma.WriteRune('1')
			epsilon.WriteRune('0')
		}
	}
	return gamma.String(), epsilon.String(), nil
}

func multiplyRates(gamma string, epsilon string) (int, error) {
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println("Gamma:", g)
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return 0, err
	}
	fmt.Println("Epsilon:", e)
	return int(g * e), nil
}

