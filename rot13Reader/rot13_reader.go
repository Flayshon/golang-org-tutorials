package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
    switch {
    case x >= 65 && x <= 77:
        fallthrough
    case x >= 97 && x <= 109:
        x += 13
    case x >= 78 && x <= 90:
        fallthrough
    case x >= 110 && x <= 122:
        x -= 13
    }
    return x
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)

	for i := 0; i <= n; i++ {
        b[i] = rot13(b[i])
    }
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
