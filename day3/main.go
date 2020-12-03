package main

import "fmt"

func p1() {
	trees := 0
	pos := 0

	for _, line := range input {
		last := pos % len(line)

		if string(line[last]) == "#" {
			trees++
		}

		pos += 3
	}

	fmt.Printf("Trees length is : %d \n", trees)
}

func p2(right int, down int) int {
	trees := 0
	pos := 0

	for index, line := range input {
		if index%down != 0 {
			continue
		}

		last := pos % len(line)

		if string(line[last]) == "#" {
			trees++
		}

		pos += right
	}

	return trees
}

func main() {
	p1()

	sum := p2(1, 1) * p2(3, 1) * p2(5, 1) * p2(7, 1) * p2(1, 2)

	fmt.Printf("Total is: %d \n", sum)

}
