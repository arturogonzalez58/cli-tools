package main

import (
	"flag"
	"fmt"
	"github.com/arturogonzalez58/cli-tools/internal/wordCounter"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	counter := wordCounter.BuildCounter(os.Stdin, *lines, nil)
	fmt.Println(counter.Count())
}
