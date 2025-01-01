package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func scan(in *os.File) ([]int, []int) {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	arr1, arr2 := []int{}, []int{}

	for {
		if !sc.Scan() {
			break
		}
		n1, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		arr1 = append(arr1, n1)

		if !sc.Scan() {
			break
		}
		n2, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		arr2 = append(arr2, n2)
	}
	return arr1, arr2
}

func part1(arr1 []int, arr2 []int) {
	slices.Sort(arr1)
	slices.Sort(arr2)

	result := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			result += arr1[i] - arr2[i]
		} else {
			result += arr2[i] - arr1[i]
		}
	}
	fmt.Printf("part1: %d\n", result)
}

func part2(arr1 []int, arr2 []int) {
	m := map[int][2]int{}

	for _, v1 := range arr1 {
		n, ok := m[v1]
		if ok {
			n[1] += 1
			m[v1] = n
		} else {
			x := 0
			for _, v2 := range arr2 {
				if v1 == v2 {
					x++
				}
			}
			m[v1] = [2]int{x * v1, 1}
		}
	}

	result := 0
	for _, v := range m {
		result += v[0] * v[1]
	}
	fmt.Printf("part2: %d\n", result)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	arr1, arr2 := scan(in)
	in.Close()

	fmt.Println("Historian Hysteria")
	part1(arr1, arr2)
	part2(arr1, arr2)
}
