package main

import "fmt"

func main() {
	test_array := [3]string{"hi", "what", "happen"}

	for i := 0; i <= 2; i++ {
		fmt.Printf("%d", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("%s", test_array[j])
		}
	}
}
