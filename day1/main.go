package main

import "fmt"

const target = 2020

func p1() {
	for _, num1 := range input {
		for _, num2 := range input {
			if num1+num2 == target {
				fmt.Printf("The numbers are %d and %d. Answer is: %d \n", num1, num2, num1*num2)
				return
			}
		}
	}
}

func p2() {
	for _, num1 := range input {
		for _, num2 := range input {
			for _, num3 := range input {
				if num1+num2+num3 == target {
					fmt.Printf("The numbers are %d, %d and %d. Answer is: %d \n", num1, num2, num3, num1*num2*num3)
					return
				}
			}
		}
	}
}

func main() {
	p1()
	p2()
}
