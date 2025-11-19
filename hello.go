package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := bufio.NewScanner(os.Stdin)
	x.Scan()
	y, _ := strconv.ParseInt(x.Text(), 10, 64)
	fmt.Println(int64(y))

}
