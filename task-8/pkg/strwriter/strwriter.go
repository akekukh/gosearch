package strwriter

import "io"

// Print - print only strings
func Print(w io.Writer, args ...interface{}) {
	for _, arg := range args {
		if val, ok := arg.(string); ok {
			w.Write([]byte(val))
		}
	}
}
