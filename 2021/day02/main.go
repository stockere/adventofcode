package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	commands, err := processPositionChanges("2021/day02/input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Aggregated commands: %+v\n", *commands)
	fmt.Println("Multipled position:", commands.getMultipliedPosition())
}

func processPositionChanges(filename string) (*commands, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var direction string
	var position int
	commands := &commands{}
	for {
		_, err := fmt.Fscanln(file, &direction, &position)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch direction {
		case "forward":
			commands.forward += position
			commands.depth += commands.aim * position
		case "up":
			commands.aim -= position
		case "down":
			commands.aim += position
		default:
			return nil, fmt.Errorf("error: unknown direction %v\n", direction)
		}
	}
	return commands, nil
}

type commands struct {
	forward int
	aim int
	depth int
}

func (c *commands) getMultipliedPosition() int {
	return c.depth * c.forward
}
