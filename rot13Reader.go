// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
// For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
// Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.



package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read (p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
    var b byte
    for i := 0; i < n; i++ {
        if c := p[i]; c >= 'A' && c <= 'z' {
            if c >= 'a' {
                b = byte('a')
            } else {
                b = byte('A')
            }
            // rotate by 13 places
            p[i] = (c - b + 13) % 26 + b
        }
    }
    return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
