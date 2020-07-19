package rgba32_test

import (
	"github.com/reiver/go-rgba32"

	"math/rand"
	"time"

	"testing"
)

func TestSlice_RGBA(t *testing.T) {

	for testNumber:=0; testNumber<10; testNumber++ {

		var r, g, b, a uint8
		{
			randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

			r = uint8(randomness.Intn(256))
			g = uint8(randomness.Intn(256))
			b = uint8(randomness.Intn(256))
			a = uint8(randomness.Intn(256))
		}


		var buffer [rgba32.ByteSize]uint8
		{
			buffer[rgba32.OffsetRed]   = r
			buffer[rgba32.OffsetGreen] = g
			buffer[rgba32.OffsetBlue]  = b
			buffer[rgba32.OffsetAlpha] = a
		}

		var rgba rgba32.Slice = rgba32.Slice(buffer[:])

		{
			aR, aG, aB, aA := rgba.RGBA()

			eR := uint32(r) * (0xffff/0xff)
			eG := uint32(g) * (0xffff/0xff)
			eB := uint32(b) * (0xffff/0xff)
			eA := uint32(a) * (0xffff/0xff)

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d, the actual color is not what was expected.", testNumber)
				t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
				t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
				continue
			}
		}
	}
}
