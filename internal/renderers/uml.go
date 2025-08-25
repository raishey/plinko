package renderers

import (
	"io"

	pub "github.com/raishey/plinko/pkg/renderers"
)

type UML = pub.UML

func NewUML(w io.Writer) *UML {
	return pub.NewUML(w)
}
