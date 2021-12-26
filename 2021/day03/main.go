package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part two
	bins, err := readBinaryToArray("2021/day03/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	o2Rating, err := recursiveFilter(0, bins, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o2Rating)
	c02Rating, err := recursiveFilter(0, bins, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c02Rating)
	fmt.Println(multiplyRates(o2Rating, c02Rating))
	/*
	Part one
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
	 */
}

// recursiveFilter will return the o2 rating (mostCommon = true) or c02 rating
// (mostCommon = false)
func recursiveFilter(position int, binaries []string, mostCommon bool) (string, error) {
	if position >= len(binaries[0]) {
		return "", fmt.Errorf("cannot filter for position out of range")
	}
	var c string
	var err error
	if mostCommon {
		c, err = mostCommonCharAt(position, binaries)
	} else {
		c, err = leastCommonCharAt(position, binaries)
	}
	if err != nil {
		return "", err
	}
	matches, err := filterMatchesAt(position, binaries, c)
	if err != nil {
		return "", err
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	return recursiveFilter(position + 1, matches, mostCommon)
}

func filterMatchesAt(position int, binaries []string, match string) ([]string, error) {
	var matches []string
	for _, binary := range binaries {
		c, err := charAt(binary, position)
		if err != nil {
			return nil, err
		}
		if c == match {
			matches = append(matches, binary)
		}
	}
	return matches, nil
}

func leastCommonCharAt(position int, binaries []string) (string, error) {
	c, err := mostCommonCharAt(position, binaries)
	if err != nil {
		return "", err
	}
	if c == "0" {
		return "1", nil
	} else {
		return "0", nil
	}
}

func mostCommonCharAt(position int, binaries []string) (string, error) {
	zeros := 0
	ones := 0
	for _, binary := range binaries {
		c, err := charAt(binary, position)
		if err != nil {
			return "", err
		}
		if c == "0" {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return "0", nil
	} else {
		return "1", nil
	}
}

func charAt(binary string, position int) (string, error) {
	switch string(binary[position]) {
	case "0": return "0", nil
	case "1": return "1", nil
	default:
		return "", fmt.Errorf("unrecognized character in string %v, expecting 0s and 1s", binary)
	}
}

func readBinaryToArray(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var binaries []string
	var binary string
	for {
		_, err := fmt.Fscanln(file, &binary)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		binaries = append(binaries, binary)
	}
	return binaries, nil
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
