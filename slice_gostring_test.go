package rgba32_test

import (
	"github.com/reiver/go-rgba32"

	"testing"
)

func TestSlice_GoString(t *testing.T) {

	tests := []struct{
		R uint8
		G uint8
		B uint8
		A uint8
		Expected string
	}{
		{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
			Expected: "rgba32.Slice{0,0,0,0}",
		},
		{
			R: 11,
			G: 22,
			B: 33,
			A: 44,
			Expected: "rgba32.Slice{11,22,33,44}",
		},
		{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
			Expected: "rgba32.Slice{255,255,255,255}",
		},


		{
			R:   1,
			G:   1,
			B:   1,
			A: 255,
			Expected: "rgba32.Slice{1,1,1,255}",
		},
		{
			R: 222,
			G:  56,
			B:  43,
			A: 255,
			Expected: "rgba32.Slice{222,56,43,255}",
		},
		{
			R:  57,
			G: 181,
			B:  74,
			A: 255,
			Expected: "rgba32.Slice{57,181,74,255}",
		},
		{
			R: 255,
			G: 199,
			B:   6,
			A: 255,
			Expected: "rgba32.Slice{255,199,6,255}",
		},
		{
			R: 0,
			G: 111,
			B: 184,
			A: 255,
			Expected: "rgba32.Slice{0,111,184,255}",
		},
		{
			R: 118,
			G:  38,
			B: 113,
			A: 255,
			Expected: "rgba32.Slice{118,38,113,255}",
		},
		{
			R:  44,
			G: 181,
			B: 233,
			A: 255,
			Expected: "rgba32.Slice{44,181,233,255}",
		},
	}

	for testNumber, test := range tests {

		var buffer [rgba32.ByteSize]uint8
		{
			buffer[rgba32.OffsetRed]   = test.R
			buffer[rgba32.OffsetGreen] = test.G
			buffer[rgba32.OffsetBlue]  = test.B
			buffer[rgba32.OffsetAlpha] = test.A
		}

		var rgba rgba32.Slice = rgba32.Slice(buffer[:])

		{
			expected := test.Expected

			actual := rgba.GoString()

			if expected != actual {
				t.Errorf("For test #%d, the actual rgba() string is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

	}
}
