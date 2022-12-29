package wordCounter

import (
	"bufio"
	"io"
)

type Counter struct {
	reader io.Reader
	lines  bool
	regex  *string
}

func BuildCounter(reader io.Reader, lines bool, regex *string) *Counter {
	return &Counter{reader: reader, lines: lines, regex: regex}
}

func (c *Counter) Count() int {
	scanner := bufio.NewScanner(c.reader)
	if !c.lines {
		scanner.Split(bufio.ScanWords)
	}
	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}
