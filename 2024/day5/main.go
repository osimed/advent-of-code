package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func scan(in *os.File) (map[int][]int, [][]int) {
	rules := map[int][]int{}
	updates := [][]int{}
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			break
		}
		rule := strings.Split(sc.Text(), "|")
		n1, _ := strconv.Atoi(rule[0])
		n2, _ := strconv.Atoi(rule[1])
		r1 := rules[n1]
		r1 = append(r1, n2)
		rules[n1] = r1
	}
	for sc.Scan() {
		update := strings.Split(sc.Text(), ",")
		_update := make([]int, len(update))
		for i, v := range update {
			n, _ := strconv.Atoi(v)
			_update[i] = n
		}
		updates = append(updates, _update)
	}
	return rules, updates
}

func inOrder(rules map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if r, ok := rules[update[j]]; ok {
				if slices.Contains(r, update[i]) {
					return false
				}
			}
		}
	}
	return true
}

func part1(rules map[int][]int, updates [][]int) {
	var result int
	for _, update := range updates {
		if inOrder(rules, update) {
			result += update[len(update)/2]
		}
	}
	fmt.Printf("part1: %d\n", result)
}

func part2(rules map[int][]int, updates [][]int) {
	var result int
	for _, update := range updates {
		if inOrder(rules, update) {
			continue
		}
		slices.SortFunc(update, func(a, b int) int {
			if r, ok := rules[a]; ok {
				if slices.Contains(r, b) {
					return -1
				}
			}
			if r, ok := rules[b]; ok {
				if slices.Contains(r, a) {
					return 1
				}
			}
			return 0
		})
		result += update[len(update)/2]
	}
	fmt.Printf("part2: %d\n", result)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	rules, updates := scan(in)
	fmt.Println("Print Queue")
	part1(rules, updates)
	part2(rules, updates)
}
