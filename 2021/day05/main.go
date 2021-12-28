package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readCoordinates("2021/day05/input.txt")
	if err != nil {
		panic(err)
	}

	var straightLines []*Line
	for _, line := range lines {
		if line.isStraight() {
			straightLines = append(straightLines, line)
		}
	}
	pointCount := make(map[Point]int)
	for _, line := range straightLines {
		points, err := line.getLineCoordinates()
		if err != nil {
			panic(err)
		}
		for _, point := range points {
			value, _ := pointCount[point]
			pointCount[point] = value + 1
		}
	}
	overlap := 0
	for _, count := range pointCount {
		if count >= 2 {
			overlap++
		}
	}
	fmt.Println(overlap)
}

func readCoordinates(filename string) ([]*Line, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	buffer := bufio.NewReader(file)
	var lines []*Line
	for {
		bytes, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		str := string(bytes)
		line, err := buildLine(str)
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func buildLine(input string) (*Line, error) {
	points := strings.Split(input, " -> ")

	p1 := strings.Split(points[0], ",")
	p1x, _ := strconv.Atoi(p1[0])
	p1y, _ := strconv.Atoi(p1[1])
	point1 := Point{
		x: p1x,
		y: p1y,
	}

	p2 := strings.Split(points[1], ",")
	p2x, _ := strconv.Atoi(p2[0])
	p2y, _ := strconv.Atoi(p2[1])
	point2 := Point{
		x: p2x,
		y: p2y,
	}

	return &Line{
		point1: point1,
		point2: point2,
	}, nil
}

type Line struct {
	point1 Point
	point2 Point
}

type Point struct {
	x int
	y int
}

// isStraight is true if the line is not diagonal
// (yes diagonal lines are also straight, naming is hard)
func (l *Line) isStraight() bool {
	return l.isHorizontal() || l.isVertical()
}

func (l *Line) isHorizontal() bool {
	return l.point1.y == l.point2.y
}

func (l *Line) isVertical() bool {
	return l.point1.x == l.point2.x
}

func (l *Line) getLineCoordinates() ([]Point, error) {
	if !l.isStraight() {
		return nil, fmt.Errorf("can't return coordinates of diagonal line")
	}
	var points []Point
	if l.isHorizontal() {
		y := l.point1.y
		var i int
		var j int
		if l.point1.x < l.point2.x {
			i = l.point1.x
			j = l.point2.x
		} else {
			i = l.point2.x
			j = l.point1.x
		}
		for ; i <= j; i++ {
			point := Point{
				x: i,
				y: y,
			}
			points = append(points, point)
		}
	}
	if l.isVertical() {
		x := l.point1.x
		var i int
		var j int
		if l.point1.y < l.point2.y {
			i = l.point1.y
			j = l.point2.y
		} else {
			i = l.point2.y
			j = l.point1.y
		}
		for ; i <= j; i++ {
			point := Point{
				x: x,
				y: i,
			}
			points = append(points, point)
		}
	}
	return points, nil
}