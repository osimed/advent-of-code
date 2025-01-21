package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scan(in *os.File) []int {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	var stones []int
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		stones = append(stones, n)
	}
	return stones
}

var seen = map[[2]int]int{}

func blink(stone, b int) int {
	if b <= 0 {
		return 1
	}
	if stone == 0 {
		return blink(1, b-1)
	}
	_s := [2]int{stone, b}
	if x, ok := seen[_s]; ok {
		return x
	}
	d := len(strconv.Itoa(stone))
	if d%2 == 1 {
		seen[_s] = blink(stone*2024, b-1)
		return seen[_s]
	}
	var p int = 1
	for a := 0; a < d/2; a++ {
		p *= 10
	}
	s1, s2 := stone/p, stone%p
	seen[_s] = blink(s1, b-1) + blink(s2, b-1)
	return seen[_s]
}

func part1(stones []int) {
	var count int = 0
	for i := 0; i < len(stones); i++ {
		count += blink(stones[i], 25)
	}
	fmt.Printf("part1: %d\n", count)
}

func part2(stones []int) {
	var count int = 0
	for i := 0; i < len(stones); i++ {
		count += blink(stones[i], 75)
	}
	fmt.Printf("part2: %d\n", count)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	stones := scan(in)
	fmt.Println("Plutonian Pebbles")
	part1(stones)
	part2(stones)
}
