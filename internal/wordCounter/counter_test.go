package wordCounter

import (
	"bytes"
	"io"
	"testing"
)

func TestCounter_Count(t *testing.T) {
	type fields struct {
		reader io.Reader
		lines  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Given a valid input get a the number of words",
			fields: fields{reader: bytes.NewBufferString("hola mundo\n")},
			want:   2,
		},
		{name: "Given a valid input with the lines flag set get a the number of words",
			fields: fields{reader: bytes.NewBufferString("hola mundo\n"), lines: true},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := BuildCounter(tt.fields.reader, tt.fields.lines, nil)
			if got := c.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
