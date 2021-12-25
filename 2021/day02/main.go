package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	commands, err := getPositionChanges("2021/day02/input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", *commands)
	fmt.Println(commands.getMultipliedPosition())
}

func getPositionChanges(filename string) (*commands, error) {
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
		case "up":
			commands.up += position
		case "down":
			commands.down += position
		default:
			return nil, fmt.Errorf("error: unknown direction %v\n", direction)
		}
	}
	return commands, nil
}

type commands struct {
	forward int
	up int
	down int
}

func (c *commands) getDepth() int {
	return c.down - c.up
}

func (c *commands) getForwardPosition() int {
	return c.forward
}

func (c *commands) getMultipliedPosition() int {
	return c.getDepth() * c.getForwardPosition()
}
