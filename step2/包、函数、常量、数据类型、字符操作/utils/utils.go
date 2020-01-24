package utils

import (
	"math"
)

func Is_Prime(num int) bool {
	//是否为质数
	if num < 2 {
		return false
	} else if num == 2 {
		return true
	} else if num%2 == 0 {
		return false
	} else {
		for i := 3; i < int(math.Sqrt(float64(num)))+1; i += 2 {
			if num%i == 0 {
				return false
			}
		}
	}
	return true
}

func Is_Narcissus_Num(num int) bool {
	//是否为水仙花数
	var num_c int = num
	var c int = 1
	var n_sum int = 0
	var r int
	for num > c*10 {
		c *= 10
	}
	for c >= 1 {
		r = num / c
		num -= r * c
		c /= 10
		n_sum += r * r * r
	}
	return num_c == n_sum
}
func Factorial_sum(num int) int {
	sum := 0
	mul := 1
	for i := 1; i <= num; i++ {
		mul *= i
		sum += mul
	}
	return sum
}
