package main

import (
	"fmt"

	"github.com/telnet2/gopkg/bench"
)

func main() {
	var length = 0
	length += len(bench.HeaderKey)
	length += len(bench.Headers.Key)
	fmt.Println(length)
}
