package wordCounter

import (
	"bufio"
	"io"
	"regexp"
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
	var r *regexp.Regexp
	if c.regex != nil {
		r, _ = regexp.Compile(*c.regex)
	}
	scanner := bufio.NewScanner(c.reader)
	if !c.lines {
		scanner.Split(bufio.ScanWords)
	}
	counter := 0
	for scanner.Scan() {
		token := scanner.Text()
		if c.regex != nil {
			if r.MatchString(token) {
				counter++
			}
		} else {
			counter++
		}
	}
	return counter
}
