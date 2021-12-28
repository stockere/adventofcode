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

	pointCount := make(map[Point]int)
	for _, line := range lines {
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

func (l *Line) isDiagonal() bool {
	return !(l.isHorizontal() || l.isVertical())
}

func (l *Line) isHorizontal() bool {
	return l.point1.y == l.point2.y
}

func (l *Line) isVertical() bool {
	return l.point1.x == l.point2.x
}

func (l *Line) getLineCoordinates() ([]Point, error) {
	var points []Point
	if l.isHorizontal() {
		y := l.point1.y
		var x int
		var xmax int
		if l.point1.x < l.point2.x {
			x = l.point1.x
			xmax = l.point2.x
		} else {
			x = l.point2.x
			xmax = l.point1.x
		}
		for ; x <= xmax; x++ {
			point := Point{
				x: x,
				y: y,
			}
			points = append(points, point)
		}
	}
	if l.isVertical() {
		x := l.point1.x
		var y int
		var ymax int
		if l.point1.y < l.point2.y {
			y = l.point1.y
			ymax = l.point2.y
		} else {
			y = l.point2.y
			ymax = l.point1.y
		}
		for ; y <= ymax; y++ {
			point := Point{
				x: x,
				y: y,
			}
			points = append(points, point)
		}
	}
	if l.isDiagonal() {
		var x int
		var xmax int
		var y int
		var increaseY bool
		if l.point1.x < l.point2.x {
			x = l.point1.x
			xmax = l.point2.x
			// y value needs to increase if point2.y is greater or decrease if it's smaller
			y = l.point1.y
			increaseY = l.point2.y > y
		} else {
			x = l.point2.x
			y = l.point2.y
			xmax = l.point1.x
			increaseY = l.point1.y > y
		}
		for ; x <= xmax; x++ {
			point := Point{
				x: x,
				y: y,
			}
			points = append(points, point)
			if increaseY {
				y++
			} else {
				y--
			}
		}
	}
	return points, nil
}