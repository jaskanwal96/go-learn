package challenge3_4

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (int, error) {
	n, err := reader.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]

		switch {
		case 'a' <= b && b <= 'z':
			p[i] = 'a' + (b-'a'+13)%26
		case 'A' <= b && b <= 'Z':
			p[i] = 'A' + (b-'A'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
