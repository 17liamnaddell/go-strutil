package main

import (
	"fmt"
	"github.com/skilstak/go-strutil"
	"os"
)

const usage string = "usage: slug STRING"

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}
	s := os.Args[1]
	fmt.Println(strutil.ToSlug(s))
}
