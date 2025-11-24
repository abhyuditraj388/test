package main

import "fmt"

func main() {

	fmt.Println(div(2, 0))

}

func div(a, b float64) float64 {
	return a / b
}
