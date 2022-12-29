package wordCounter

import (
	"bufio"
	"io"
)

type Counter struct {
	Reader io.Reader
}

func BuildCounter(reader io.Reader) *Counter {
	return &Counter{Reader: reader}
}

func (c *Counter) Count() int {
	scanner := bufio.NewScanner(c.Reader)
	scanner.Split(bufio.ScanWords)
	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}
