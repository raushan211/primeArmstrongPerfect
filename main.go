package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter three digit no.")
	fmt.Scanln(&num)
	fmt.Println(primeArmstrongPerfect(num))
}
func primeArmstrongPerfect(num int) string {
	var result string
	var remainder int
	var multiply int
	var count int
	var sum int
	for i := 2; i < num/2; i++ {
		{
			if num%i == 0 {
				count++
				break
			}
		}
		if count == 0 && num != 1 {
			result = "prime no."
		}
	}

	temp := num
	for num > 0 {

		remainder = num % 10
		multiply = remainder*remainder*remainder + multiply
		num = num / 10

	}
	if multiply == temp {
		result = " is armstrong no."
	}
	for i := 1; i < num; i++ {
		{
			remainder = num % i
			if remainder == 0 {
				sum = sum + i
			}

		}

		if sum == num {
			result = " perfect no."

		}
	}

	return result
}
