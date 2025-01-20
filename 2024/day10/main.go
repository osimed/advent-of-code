package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

func scan(in *os.File) [][]int {
	sc := bufio.NewScanner(in)
	m := make([][]int, 0)
	for sc.Scan() {
		l := len(sc.Bytes())
		r := make([]int, l)
		for i, b := range sc.Bytes() {
			r[i] = int(b - '0')
		}
		m = append(m, r)
	}
	return m
}

func paths(m [][]int) iter.Seq2[int, [10][2]int] {
	count, path := 0, [10][2]int{}
	var loop func(y, x, h int, yield func(int, [10][2]int) bool)
	loop = func(y, x, h int, yield func(int, [10][2]int) bool) {
		path[h] = [2]int{y, x}
		if h == 9 {
			if !yield(count, path) {
				return
			} else {
				count++
			}
		}
		positions := [][2]int{
			{y + 0, x - 1}, {y + 0, x + 1},
			{y - 1, x + 0}, {y + 1, x + 0},
		}
		for _, p := range positions {
			if p[0] < 0 || p[0] >= len(m) {
				continue
			}
			if p[1] < 0 || p[1] >= len(m[0]) {
				continue
			}
			if m[p[0]][p[1]] != h+1 {
				continue
			}
			if h < 9 {
				loop(p[0], p[1], h+1, yield)
			}
		}
	}
	return func(yield func(int, [10][2]int) bool) {
		for y, row := range m {
			for x, height := range row {
				if height != 0 {
					continue
				}
				loop(y, x, 0, yield)
			}
		}
	}
}

func part1(m [][]int) {
	f := map[[4]int]struct{}{}
	for _, v := range paths(m) {
		p := [4]int{
			v[0][0], v[0][1],
			v[9][0], v[9][1],
		}
		f[p] = struct{}{}
	}
	fmt.Printf("part1: %v\n", len(f))
}

func part2(m [][]int) {
	var count int
	for count = range paths(m) {
	}
	fmt.Printf("part2: %v\n", count+1)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	m := scan(in)
	fmt.Println("Hoof It")
	part1(m)
	part2(m)
}
