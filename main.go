package main

import (
	"flag"
	"fmt"
)

var (
	yaml = flag.String("data", "", "Location data test yaml file")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(*yaml)
}
