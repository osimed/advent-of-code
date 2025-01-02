package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sep = [][]byte{
	[]byte("mul("),
	[]byte("do()"),
	[]byte("don't()"),
}

func tokenize(data []byte, atEOF bool) (int, []byte, error) {
	number := func(i, j int) bool {
		if j-i < 1 || j-i > 3 {
			return false
		}
		for k := i; k < j; k++ {
			x := data[k]
			if x < '0' || x > '9' {
				return false
			}
		}
		return true
	}
	findFirst := func(dx int) (int, int) {
		x, f := -1, -1
		for s, sep := range sep {
			i := bytes.Index(data[dx:], sep)
			if i >= 0 && (i < f || f == -1) {
				x, f = s, i
			}
		}
		if x == -1 {
			return x, f
		}
		return x, f + dx + len(sep[x])
	}
	dx := 0
	for dx < len(data) {
		x, i := findFirst(dx)
		if x < 0 {
			return 0, nil, nil
		}
		if x == 1 || x == 2 {
			return i, sep[x], nil
		}
		j := bytes.IndexByte(data[i:], ',') + i
		if j < i || !number(i, j) {
			dx = i
			continue
		}
		k := bytes.IndexByte(data[j:], ')') + j
		if k < j || !number(j+1, k) {
			dx = j
			continue
		}
		return k + 1, data[i:k], nil
	}
	return 0, nil, nil
}

type mul struct {
	x1, x2   int
	disabled bool
}

func scan(in *os.File) []mul {
	sc := bufio.NewScanner(in)
	sc.Split(tokenize)
	instrs := []mul{}
	var disabled bool
	for sc.Scan() {
		if sc.Text() == "do()" {
			disabled = false
			continue
		}
		if sc.Text() == "don't()" {
			disabled = true
			continue
		}
		m := strings.Split(sc.Text(), ",")
		x1, _ := strconv.Atoi(m[0])
		x2, _ := strconv.Atoi(m[1])
		instr := mul{x1, x2, disabled}
		instrs = append(instrs, instr)
	}
	return instrs
}

func part1(instrs []mul) int {
	var result int = 0
	for _, mul := range instrs {
		result += mul.x1 * mul.x2
	}
	return result
}

func part2(instrs []mul) int {
	var result int = 0
	for _, mul := range instrs {
		if !mul.disabled {
			result += mul.x1 * mul.x2
		}
	}
	return result
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	instrs := scan(in)
	in.Close()

	result1 := part1(instrs)
	fmt.Printf("part1: %d\n", result1)

	result2 := part2(instrs)
	fmt.Printf("part2: %d\n", result2)
}
