package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	x := bufio.NewScanner(os.Stdin)
	x.Scan()
	fmt.Println(x.Text())

}
