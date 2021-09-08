package main

import "fmt"

func solution(price, money, count int) int64 {
	sum := 0
	for i := 0; i <= count; i++ {
		sum += (price * i)
	}
	if money >= sum {
		return 0
	} else {
		total := money - sum
		return int64(-total)
	}

}

func main() {
	price := 3
	money := 20
	count := 4
	result := solution(price, money, count)
	fmt.Println(result)
}
