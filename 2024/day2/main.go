package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan(in *os.File) ([][]int, error) {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanLines)

	reports := [][]int{}
	for {
		if !sc.Scan() {
			break
		}
		_levels := strings.Fields(sc.Text())
		levels := make([]int, len(_levels))
		for i := range levels {
			level, err := strconv.Atoi(_levels[i])
			if err != nil {
				return nil, err
			}
			levels[i] = level
		}
		reports = append(reports, levels)
	}
	return reports, nil
}

func IsSafe(report []int, n int) bool {
	var set, inc bool
	for i := 0; i < len(report)-1; i++ {
		if i == n {
			continue
		}
		if i+1 == n && i+2 >= len(report) {
			continue
		}
		dx := report[i] - report[i+1]
		if i+1 == n && i+2 < len(report) {
			dx = report[i] - report[i+2]
		}
		if dx == 0 || dx > 3 || dx < -3 {
			return false
		}
		if set && inc && dx > 0 {
			return false
		}
		if set && !inc && dx < 0 {
			return false
		}
		set, inc = true, dx < 0
	}
	return true
}

func part1(reports [][]int) {
	safe := 0
	for _, report := range reports {
		if IsSafe(report, -1) {
			safe++
		}
	}
	fmt.Printf("part1: %v\n", safe)
}

func part2(reports [][]int) {
	safe := 0
	for _, report := range reports {
		if IsSafe(report, -1) {
			safe++
			goto exit
		}
		for i := range report {
			if IsSafe(report, i) {
				safe++
				goto exit
			}
		}
	exit:
	}
	fmt.Printf("part2: %v\n", safe)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	reports, err := scan(in)
	if err != nil {
		panic(err)
	}
	in.Close()

	fmt.Println("Red-Nosed Reports")
	part1(reports)
	part2(reports)
}
