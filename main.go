package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	str := os.Args[1]
	splitedStr := strings.Split(str, " ")
	fmt.Println(len(splitedStr))
}
