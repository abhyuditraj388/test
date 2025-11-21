package main

import "fmt"

type dog struct {
	name string
	age  int
}

func (d dog) bark() string {
	return d.name + " is barking"
}

func main() {

	tommy := dog{"tommy", 25}

	fmt.Println(tommy.bark())

}
