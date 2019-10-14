package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, result int

	fmt.Scanf("%d", &n)

	if n < 0 {
		fmt.Print("Please input positive number")
		return
	}

	for i := 1; i <= n; i++ {
		numberLen := len(strconv.Itoa(i))
		numString := string(strconv.Itoa(i))

		if numberLen == 1 {
			result = i + getDigit(numString, 0)
		} else if numberLen == 2 {
			result = i + getDigit(numString, 0) + getDigit(numString, 1)
		} else if numberLen == 3 {
			result = i + getDigit(numString, 0) + getDigit(numString, 1) + getDigit(numString, 2)
		} else if numberLen == 4 {
			result = i + getDigit(numString, 0) + getDigit(numString, 1) + getDigit(numString, 2) + getDigit(numString, 3)
		}

		fmt.Printf("Hasil %d", result)
	}

}

func getDigit(number string, position int) int {
	res := string(number[position])

	numInt, err := strconv.Atoi(res)

	if err != nil {
		fmt.Printf("Fatal error")
	}

	return numInt
}
