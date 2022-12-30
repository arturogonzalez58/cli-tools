package main

import (
	"flag"
	"fmt"
	"github.com/arturogonzalez58/cli-tools/internal/wordCounter"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	regex := flag.String("r", "", "Regex to count each word")
	flag.Parse()
	counter := wordCounter.BuildCounter(os.Stdin, *lines, regex)
	fmt.Println(counter.Count())
}
