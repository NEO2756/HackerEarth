package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func main() {

	var T int
	var num string

	fmt.Scanf("%d", &T)
	arr := [10]string{"6", "2", "5", "5", "4", "5", "6", "3", "7", "6"}
	for i := 0; i < T; i++ {
		fmt.Scanf("%s", &num)
		//fmt.Printf("%s", num)
		maxStick := 0
		for _, char := range num {
			e, _ := strconv.Atoi(string(char))
			stick, _ := strconv.Atoi(arr[e])
			maxStick += stick //arr[digit(N, i)]
		}

		if maxStick%2 == 0 {
			fmt.Println(strings.Repeat("1", maxStick/2))
		} else {
			maxStick -= 3 //these three will form 7
			fmt.Printf("%s%s\n", "7", strings.Repeat("1", maxStick/2))
		}
	}
}
