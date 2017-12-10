package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func check_alpha(b byte) int {
	switch {
	case (b > 'a' && b < 'n') || (b > 'A' && b < 'N'):
		return 13
	case (b > 'M' && b <= 'Z') || (b > 'm' && b <= 'z'):
		return -13
	default:
		return 0
	}
}

func (rt13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rt13.r.Read(b)
	for x := 0; x < n; x++ {
		var oset int = check_alpha(b[x])
		switch oset {
		case 13:
			b[x] += 13
		case -13:
			b[x] -= 13
		}
	}
	fmt.Printf("n=%d err=%v len(b)=%d\n", n, err, len(b))
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
