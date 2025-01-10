package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func scan(in *os.File) map[int][]int {
	sc := bufio.NewScanner(in)
	eqs := make(map[int][]int)
	for sc.Scan() {
		eq := strings.Split(sc.Text(), ":")
		val, _ := strconv.Atoi(eq[0])
		_nums := strings.Fields(eq[1])
		nums := make([]int, len(_nums))
		for i, n := range _nums {
			nums[i], _ = strconv.Atoi(n)
		}
		eqs[val] = nums
	}
	return eqs
}

func seq(c []byte, l int) iter.Seq2[int, []byte] {
	comb := 1
	for i := 0; i < l; i++ {
		comb *= len(c)
	}
	return func(yield func(int, []byte) bool) {
		for i := 0; i < comb; i++ {
			d, buf := i, make([]byte, l)
			for j := 0; j < l; j++ {
				digit := d % len(c)
				buf[j] = c[digit]
				d /= len(c)
			}
			if !yield(i, buf) {
				break
			}
		}
	}
}

func eval(c []byte, eqs map[int][]int) int64 {
	var count int64
	for val, eq := range eqs {
		l := len(eq) - 1
		for _, s := range seq(c, l) {
			result := eq[0]
			for i, v := range s {
				switch v {
				case '+':
					result += eq[i+1]
				case '*':
					result *= eq[i+1]
				case '|':
					n1 := strconv.Itoa(result)
					n2 := strconv.Itoa(eq[i+1])
					result, _ = strconv.Atoi(n1 + n2)
				}
			}
			if result == val {
				count += int64(val)
				break
			}
		}
	}
	return count
}

func part1(eqs map[int][]int) {
	c := []byte{'+', '*'}
	count := eval(c, eqs)
	fmt.Printf("part1: %d\n", count)
}

func part2(eqs map[int][]int) {
	c := []byte{'+', '*', '|'}
	count := eval(c, eqs)
	fmt.Printf("part2: %d\n", count)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	eqs := scan(in)
	fmt.Println("Bridge Repair")
	part1(eqs)
	part2(eqs)
}
