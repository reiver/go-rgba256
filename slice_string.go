package rgba32

import (
	"fmt"
)

func (receiver Slice) String() string {

	r := receiver[OffsetRed]
	g := receiver[OffsetGreen]
	b := receiver[OffsetBlue]
	a := receiver[OffsetAlpha]

	return fmt.Sprintf("rgba(%d,%d,%d,%d)", r,g,b,a)
}
