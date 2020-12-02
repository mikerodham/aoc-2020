package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1-3 a: abcde - VALID
// 1-3 b: cdefg - INVALID
// 2-9 c: ccccccccc - VALID
// rule letter: pw

func p1() {
	valid := 0
	for _, rule := range input {
		sections := strings.Split(rule, " ")

		count := sections[0]
		letter := sections[1][0:1]
		password := sections[2]

		numbers := strings.Split(count, "-")

		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		if strings.Count(password, letter) >= num1 {
			if strings.Count(password, letter) <= num2 {
				valid++
			}
		}
	}

	fmt.Printf("The amount of valid passwords is: %d \n", valid)
}

func p2() {
	valid := 0
	for _, rule := range input {
		sections := strings.Split(rule, " ")

		count := sections[0]
		letter := sections[1][0:1]
		password := sections[2]

		numbers := strings.Split(count, "-")

		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		if string(password[num1-1]) == letter {
			if string(password[num2-1]) == letter {
				continue
			} else {
				valid++
			}
		}

		if string(password[num2-1]) == letter {
			valid++
		}
	}

	fmt.Printf("The amount of valid passwords for part 2 is: %d \n", valid)
}

func main() {
	p1()
	p2()
}
