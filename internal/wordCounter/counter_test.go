package wordCounter

import (
	"bytes"
	"io"
	"testing"
)

func TestCounter_Count(t *testing.T) {
	type fields struct {
		reader io.Reader
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				Reader: tt.fields.reader,
			}
			if got := c.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
