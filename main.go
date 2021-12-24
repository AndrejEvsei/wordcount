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
	count := 0

	for _, i := range splitedStr {
		if i != "" {
			count++
		}

	}
	fmt.Println(count)
}
