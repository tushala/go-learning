package utils

import (
	"math"
)

func Is_WanNum(num int) bool {
	factor_sum := 0
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			factor_sum += i
			factor_sum += int(num / i)
		}
	}
	return num == factor_sum+1
}
func Is_Palindrome(str string) bool {
	length := len(str)
	for i := 0; i < int(length)/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}
