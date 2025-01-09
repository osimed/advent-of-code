package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan(in *os.File) (vec2, [][]byte) {
	sc := bufio.NewScanner(in)
	p, area := vec2{}, [][]byte{}
	for y := 0; sc.Scan(); y++ {
		for x, v := range sc.Bytes() {
			if v == '^' {
				p.y, p.x = y, x
			}
		}
		a := []byte(sc.Text())
		area = append(area, a)
	}
	return p, area
}

type vec2 struct {
	y, x int
}

func (v vec2) turn() vec2 {
	switch v {
	case vec2{-1, 0}:
		return vec2{0, 1}

	case vec2{0, 1}:
		return vec2{1, 0}

	case vec2{1, 0}:
		return vec2{0, -1}

	case vec2{0, -1}:
		return vec2{-1, 0}
	}
	return v
}

func (v vec2) in(area [][]byte) bool {
	y, x, ly, lx := v.y, v.x, len(area), len(area[0])
	return y >= 0 && x >= 0 && y < ly && x < lx
}

func (v vec2) walk(area [][]byte, d vec2) bool {
	o := make(map[vec2]vec2)
	for {
		if area[v.y][v.x] == '.' {
			area[v.y][v.x] = 'X'
		}
		n := vec2{v.y + d.y, v.x + d.x}
		if !n.in(area) {
			break
		}
		if area[n.y][n.x] == '#' {
			od, ok := o[n]
			if ok && od == d {
				return false
			}
			o[n], d = d, d.turn()
			continue
		}
		v = vec2{v.y + d.y, v.x + d.x}
	}
	return true
}

func part1(p vec2, area [][]byte) {
	d, count := vec2{-1, 0}, 1
	p.walk(area, d)
	for _, line := range area {
		for _, point := range line {
			if point == 'X' {
				count++
			}
		}
	}
	fmt.Printf("part1: %d\n", count)
}

func part2(p vec2, area [][]byte) {
	d, count := vec2{-1, 0}, 0
	p.walk(area, d)
	a := make([][]byte, len(area))
	for y := range area {
		l := len(area[y])
		a[y] = make([]byte, l)
		for x := range area[y] {
			a[y][x] = area[y][x]
		}
	}
	for y := range area {
		for x := range area[y] {
			if area[y][x] == 'X' {
				a[y][x] = '#'
				if !p.walk(a, d) {
					count++
				}
				a[y][x] = '.'
			}
		}
	}
	fmt.Printf("part2: %d\n", count)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	p, area := scan(in)
	fmt.Println("Guard Gallivant")
	part1(p, area)
	part2(p, area)
}
