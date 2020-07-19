package rgba32

import (
	"fmt"
)

func (receiver Slice) GoString() string {

	r := receiver[OffsetRed]
	g := receiver[OffsetGreen]
	b := receiver[OffsetBlue]
	a := receiver[OffsetAlpha]

	return fmt.Sprintf("rgba32.Slice{%d,%d,%d,%d}", r,g,b,a)
}
