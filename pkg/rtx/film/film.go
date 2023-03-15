package film

import (
	"io"
	"sergioffpc/rtx/pkg/rtx/color"
)

type Film interface {
	Set(x, y int, s color.Spectrum)
	Write(w io.Writer)
}
