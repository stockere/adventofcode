package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	numbers 	[]int
	boards 		[]*Board
	gameWon 	bool
}

type Board struct {
	rows    []map[int]int
	columns []map[int]int
	winner  bool
}

func main() {
	game, err := loadGame("2021/day04/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	for _, num := range game.numbers {
		err := game.markNumber(num)
		if err != nil {
			panic(err)
		}
		if game.gameWon {
			fmt.Println("Game won for number:", num)
			fmt.Println("Winning score:", game.score(num))
			break
		}
	}
}

func loadGame(filename string) (*Game, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var numbers string
	// parse the 'random' numbers
	buffer := bufio.NewReader(file)
	line, _, err := buffer.ReadLine()
	numbers = string(line)
	if err != nil {
		return nil, err
	}
	fmt.Println(numbers)
	strs := strings.Split(numbers, ",")
	var n []int
	for _, s := range strs {
		m, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		n = append(n, m)
	}
	// parse the boards
	var boards []*Board
	for {
		board, err := makeBoard(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	return &Game{
		numbers: n,
		boards:  boards,
		gameWon: false,
	}, nil
}

func makeBoard(buffer *bufio.Reader) (*Board, error) {
	var rows []map[int]int
	//var columns []map[int]int
	columns := make([]map[int]int, 5)
	lines := 0
	for lines < 5 {
		line, _, err := buffer.ReadLine()
		if err != nil {
			return nil, err
		}
		numbers := string(line)
		if len(numbers) == 0 {
			continue
		}
		numbers = strings.TrimSpace(numbers)
		numbers = strings.ReplaceAll(numbers, "  ", " ")
		strs := strings.Split(numbers, " ")
		//fmt.Printf("Length: %v, array: %v\n", len(strs), strs)
		row := make(map[int]int)
		for i, s := range strs {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			// row
			row[n] = 0
			// column
			if columns[i] == nil {
				columns[i] = make(map[int]int)
			}
			columns[i][n] = 0
		}
		rows = append(rows, row)
		lines++
	}
	return &Board{rows: rows, columns: columns}, nil
}

func (g *Game) markNumber(number int) error {
	if g.boards == nil {
		return fmt.Errorf("nil boards")
	}
	for _, board := range g.boards {
		// mark the rows
		for _, row := range board.rows {
			if _, ok := row[number]; ok {
				delete(row, number)
				if len(row) == 0 {
					board.winner = true
					g.gameWon = true
				}
			}
		}
		// mark the columns
		for _, column := range board.columns {
			if _, ok := column[number]; ok {
				delete(column, number)
				if len(column) == 0 {
					board.winner = true
					g.gameWon = true
				}
			}
		}
	}
	return nil
}

func (g *Game) score(num int) int {
	for _, board := range g.boards {
		if board.winner {
			sum := 0
			for _, row := range board.rows {
				for key := range row {
					sum += key
				}
			}
			return num * sum
		}
	}
	return 0
}
