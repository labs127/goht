package goht

import (
	"flag"
	"fmt"
)

var (
	data = flag.String("data", "", "Location data test yaml file")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(*data)
}
