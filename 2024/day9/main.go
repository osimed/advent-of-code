package main

import (
	"bufio"
	"fmt"
	"os"
)

const dot int = -1

func scan(in *os.File) []int {
	sc := bufio.NewScanner(in)
	if !sc.Scan() {
		panic("scan error")
	}
	var id int
	var dmap []int
	for i, b := range sc.Bytes() {
		for j := byte('0'); j < b; j++ {
			if i%2 == 0 {
				dmap = append(dmap, id)
			} else {
				dmap = append(dmap, dot)
			}
		}
		if i%2 == 1 {
			id++
		}
	}
	return dmap
}

func part1(dmap []int) {
	for i, j := 0, len(dmap)-1; ; {
		if i >= j {
			break
		}
		if dmap[i] != dot {
			i++
			continue
		}
		if dmap[j] == dot {
			j--
			continue
		}
		dmap[i], dmap[j] = dmap[j], dot
		i, j = i+1, j-1
	}
	var count int
	for i := 0; dmap[i] != dot; i++ {
		count += i * int(dmap[i])
	}
	fmt.Printf("part1: %v\n", count)
}

func part2(dmap []int) {
	type block struct {
		pos, len int
	}
	files := []block{}
	for i := 0; i < len(dmap); i++ {
		if dmap[i] != dot {
			p, l, id := i, 0, dmap[i]
			for i < len(dmap) && dmap[i] == id {
				i, l = i+1, l+1
			}
			files = append(files, block{p, l})
			i--
		}
	}
	for i := len(files) - 1; i >= 0; i-- {
		if files[i].pos == dot {
			continue
		}
		sp, sl, f := -1, -1, files[i]
		for j := 0; j < f.pos; j++ {
			if dmap[j] != dot {
				continue
			}
			for h := j; h < len(dmap); h++ {
				if dmap[h] == dot {
					continue
				}
				if h-j < f.len {
					break
				}
				sp, sl = j, h-j
				goto moveFile
			}
		}
	moveFile:
		if sl > 0 {
			for h := 0; h < f.len; h++ {
				dmap[h+sp] = dmap[h+f.pos]
				dmap[h+f.pos] = dot
			}
			files[i].pos = dot
		}
	}
	var count int
	for i := 0; i < len(dmap); i++ {
		if dmap[i] != dot {
			count += int(dmap[i]) * i
		}
	}
	fmt.Printf("part2: %d\n", count)
}

func main() {
	in, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	dmap := scan(in)
	fmt.Println("Disk Fragmenter")
	dmap1, dmap2 := []int{}, []int{}
	dmap1 = append(dmap1, dmap...)
	dmap2 = append(dmap2, dmap...)
	part1(dmap1)
	part2(dmap2)
}
