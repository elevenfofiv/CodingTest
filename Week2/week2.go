package main

import "fmt"

func solution(scores [][]int) string {
	var max, min, repeat, sum, average int
	var result string
	var result_people string
	i_reshape := make([][]int, len(scores))

	for i := 0; i < len(scores); i++ {
		i_reshape[i] = make([]int, len(scores[i]))
		for index, value := range scores {
			i_reshape[i][index] = value[i]
		}
	}

	for index, i := range i_reshape {
		max = i[index]
		min = i[index]
		repeat = 0
		sum = 0
		for _, j := range i {
			if j > max {
				max = j
			}
			if j < min {
				min = j
			}
			if j == i[index] {
				repeat += 1
			}
		}

		if (i[index] > min) && (i[index] < max) {
			repeat = 0
		}
		if repeat == 1 {
			i = append(i[:index], i[index+1:]...)
		}

		for _, val := range i {
			sum += val
		}

		average = sum / len(i)

		switch {
		case (average >= 90):
			result_people = "A"
		case (average >= 80) && (average < 90):
			result_people = "B"
		case (average >= 70) && (average < 80):
			result_people = "C"
		case (average >= 50) && (average < 70):
			result_people = "D"
		default:
			result_people = "F"
		}
		result += result_people
	}
	return result
}

func main() {
	score := [][]int{{100, 90, 98, 88, 65}, {50, 45, 99, 85, 77}, {47, 88, 95, 80, 67}, {61, 57, 100, 80, 65}, {24, 90, 94, 75, 65}}
	result := solution(score)
	fmt.Println(result)

}
