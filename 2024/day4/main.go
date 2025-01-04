package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func scan(in io.Reader) [][]byte {
	sc := bufio.NewScanner(in)
	var fields [][]byte
	for sc.Scan() {
		m := []byte(sc.Text())
		fields = append(fields, m)
	}
	return fields
}
func part1(fields [][]byte) {
	getPoint := func(ix, iy int) byte {
		x, y := len(fields[0]), len(fields)
		if ix >= 0 && iy >= 0 && ix < x && iy < y {
			return fields[iy][ix]
		}
		return '#'
	}
	var checkDirection func(int, int, int, [2]int) bool
	checkDirection = func(x, y, depth int, d [2]int) bool {
		if depth > 3 {
			return true
		}
		var c byte
		var dx, dy int
		switch depth {
		case 0:
			c = 'X'
		case 1:
			c = 'M'
		case 2:
			c = 'A'
		case 3:
			c = 'S'
		}
		dx, dy = depth*int(d[0]), depth*int(d[1])
		if getPoint(x+dx, y+dy) == c {
			return checkDirection(x, y, depth+1, d)
		}
		return false
	}
	dirs := [8][2]int{
		{1, 1}, {1, 0}, {1, -1}, {0, -1},
		{-1, -1}, {-1, 0}, {-1, 1}, {0, 1},
	}
	var result int
	for y, field := range fields {
		for x := range field {
			for _, d := range dirs {
				if checkDirection(x, y, 0, d) {
					result++
				}
			}
		}
	}
	fmt.Printf("part1: %d\n", result)
}

func part2(fields [][]byte) {
	var result int
	for y := 1; y < len(fields)-1; y++ {
		for x := 1; x < len(fields[y])-1; x++ {
			if fields[y][x] != 'A' {
				continue
			}
			nw, se := fields[y+1][x-1], fields[y-1][x+1]
			ne, sw := fields[y+1][x+1], fields[y-1][x-1]
			if (nw == 'M' && se == 'S') || (nw == 'S' && se == 'M') {
				if (ne == 'M' && sw == 'S') || (ne == 'S' && sw == 'M') {
					result++
				}
			}
		}
	}
	fmt.Printf("part2: %d\n", result)
}
func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	fields := scan(in)
	in.Close()

	fmt.Println("Ceresult Search")
	part1(fields)
	part2(fields)
}
