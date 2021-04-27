package main

import "fmt"

func main() {
	for i := 0; i <=27 ; i++ {
		if isPowerOfThree(i){
			fmt.Println(i)
		}
	}
}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n % 3 == 0 {
		return isPowerOfThree(n/3)
	}
	return false
}