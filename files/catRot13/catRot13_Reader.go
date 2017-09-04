/* cat_rot13.go */

package catRot13

type Reader interface {
	Read(b []byte) (ret int, err error)
	String() string
}
