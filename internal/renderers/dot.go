package renderers

import (
	"io"

	pub "github.com/raishey/plinko/pkg/renderers"
)

type Dot = pub.Dot

func NewDot(w io.Writer) *Dot {
	return pub.NewDot(w)
}

func DotFileToImg(from, to, format string) error {
	return pub.DotFileToImg(from, to, format)
}
