package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution(fees []int, records []string) []int {
	temp := map[string]map[string][]string{}
	parking_fee := []int{}
	parking_time_list := []int{}
	car_list := []string{}

	var fee int
	var round_up int

	for _, str := range records {
		record_list := strings.Split(str, " ")
		if temp[record_list[1]] == nil {
			temp[record_list[1]] = make(map[string][]string)
		}
		if temp[record_list[1]][record_list[2]] == nil {
			temp[record_list[1]][record_list[2]] = make([]string, 0)
		}
		temp[record_list[1]][record_list[2]] = append(temp[record_list[1]][record_list[2]], record_list[0])
	}

	for key := range temp {
		car_list = append(car_list, key)
		if len(temp[key]["IN"]) > len(temp[key]["OUT"]) {
			temp[key]["OUT"] = append(temp[key]["OUT"], "23:59")
		}
	}

	sort.Strings(car_list)

	for _, car_no := range car_list {
		var parking_time int
		for i := range temp[car_no]["OUT"] {
			in_hour, _ := strconv.Atoi(temp[car_no]["IN"][i][0:2])
			out_hour, _ := strconv.Atoi(temp[car_no]["OUT"][i][0:2])
			in_minute, _ := strconv.Atoi(temp[car_no]["IN"][i][3:5])
			out_minute, _ := strconv.Atoi(temp[car_no]["OUT"][i][3:5])
			in_time_conv_minute := (in_hour * 60) + (in_minute)
			out_time_conv_minute := (out_hour * 60) + (out_minute)
			temp_parking_time := out_time_conv_minute - in_time_conv_minute
			parking_time += temp_parking_time
		}
		parking_time_list = append(parking_time_list, parking_time)
	}

	for _, time := range parking_time_list {
		switch {
		case time > fees[0]:
			if (time-fees[0])%fees[2] != 0 {
				round_up = ((time - fees[0]) / fees[2]) + 1
			} else {
				round_up = (time - fees[0]) / fees[2]
			}
			fee = fees[1] + round_up*fees[3]
		case time <= fees[0]:
			fee = fees[1]
		}
		parking_fee = append(parking_fee, fee)

	}
	return parking_fee
}

func main() {
	fees := []int{180, 5000, 10, 600}
	records := []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT",
		"07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"}

	result := solution(fees, records)
	fmt.Println(result)
}
