/* cat_rot13_impl.go */

package catRot13

type Rotate13 struct {
	source Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= ('Z'-13) {
		b += 13
	}
	if b >= ('Z'-12) && b < 'Z' {
		b = 'A' + (b - ('Z' - 12))
	}
	if b >= 'a' && b <= ('z'-13) {
		b += 13
	}
	if b >= ('z'-12) && b < 'z' {
		b = 'a' + (b - ('z' - 12))
	}
	return b
}

func NewRotate13(source Reader) *Rotate13 {
	return &Rotate13{source}
}

func (r13 *Rotate13) Read(b []byte) (ret int, err error) {
	s := r13.source
	r, e := s.Read(b)
	for i := 0; i < r; i++ {
		b[i] = rot13(b[i])
	}
	return r, e
}

func (r13 *Rotate13) String() string {
	return r13.source.String()
}
