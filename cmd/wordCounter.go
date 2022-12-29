package main

import (
	"fmt"
	"github.com/arturogonzalez58/cli-tools/internal/wordCounter"
	"os"
)

func main() {
	counter := wordCounter.BuildCounter(os.Stdin)
	fmt.Println(counter.Count())
}
