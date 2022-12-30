package wordCounter

import (
	"bytes"
	"io"
	"testing"
)

func TestCounter_Count(t *testing.T) {
	regex := "hola"
	type fields struct {
		reader io.Reader
		lines  bool
		regex  *string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Given a valid input get a the number of words",
			fields: fields{reader: bytes.NewBufferString("hola mundo\n"), regex: nil},
			want:   2,
		},
		{name: "Given a valid input with the lines flag set get a the number of words",
			fields: fields{reader: bytes.NewBufferString("hola mundo\n"), lines: true, regex: nil},
			want:   1,
		}, {name: "Given a valid input with a regex get a the number of words that match the regex",
			fields: fields{reader: bytes.NewBufferString("hola mundo\n"), lines: false, regex: &regex},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := BuildCounter(tt.fields.reader, tt.fields.lines, tt.fields.regex)
			if got := c.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
