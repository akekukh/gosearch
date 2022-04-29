package strwriter

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{
			name: "strings",
			args: []interface{}{"foo", "bar"},
			want: "foobar",
		},
		{
			name: "string, numbers, nil",
			args: []interface{}{"foo", 1, nil},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Print(w, tt.args...)
			if got := w.String(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
