package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan(in *os.File) []string {
	sc := bufio.NewScanner(in)
	antennas := make([]string, 0)
	for sc.Scan() {
		antennas = append(antennas, sc.Text())
	}
	return antennas
}

func part1(antennas []string) {
	antinodes := make(map[[2]int]bool)
	for y0, r0 := range antennas {
		for x0, antenna0 := range r0 {
			if antenna0 == '.' {
				continue
			}
			for y1, r1 := range antennas {
				for x1, antenna1 := range r1 {
					if antenna1 != antenna0 {
						continue
					}
					if y0 == y1 && x0 == x1 {
						continue
					}
					y, x := 2*y0-y1, 2*x0-x1
					if y >= 0 && y < len(antennas) {
						if x >= 0 && x < len(antennas[0]) {
							antinodes[[2]int{y, x}] = true
						}
					}
				}
			}
		}
	}
	fmt.Printf("part1: %d\n", len(antinodes))
}

func part2(antennas []string) {
	antinodes := make(map[[2]int]bool)
	for y0, r0 := range antennas {
		for x0, antenna0 := range r0 {
			if antenna0 == '.' {
				continue
			}
			for y1, r1 := range antennas {
				for x1, antenna1 := range r1 {
					if antenna1 != antenna0 {
						continue
					}
					for i := 0; y0 != y1 && x0 != x1; i++ {
						y, x := y0+i*(y0-y1), x0+i*(x0-x1)
						if y < 0 || y >= len(antennas) {
							break
						}
						if x < 0 || x >= len(antennas[y]) {
							break
						}
						antinodes[[2]int{y, x}] = true
					}
				}
			}
		}
	}
	fmt.Printf("part2: %d\n", len(antinodes))
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	antennas := scan(in)
	fmt.Println("Resonant Collinearity")
	part1(antennas)
	part2(antennas)
}
