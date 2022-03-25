package main

import "fmt"

func main() {
	n := createSlice()
	checkOddsAndEvens(n)
}

func createSlice() []int {
	var n []int
	for i := 0; i <= 10; i++ {
		n = append(n, i)
	}
	return n
}

func checkOddsAndEvens(n []int) {
	for _, v := range n {
		if v % 2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		} 
	}
}
