package main

import "fmt"

// constant di luar fungsi
const LENGTH = 10
const WIDTH = 5

func main(){

	// constant di dalam fungsi
	const PIE float32 = 3.14

    fmt.Println(LENGTH)
    fmt.Println(WIDTH)
	fmt.Println(PIE)
}