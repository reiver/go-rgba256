package rgba32_test

import (
	"github.com/reiver/go-rgba32"

	"image"
	"image/color"
	"math/rand"
	"time"

	"testing"
)

func TestSlice_At(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var buffer [rgba32.ByteSize]uint8
		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var rgba rgba32.Slice = rgba32.Slice(buffer[:])

		var img image.Image = rgba

		for subTestNumber:=0; subTestNumber<50; subTestNumber++ {

			x := randomness.Int()
			if 0 == x%2 {
				x = -x
			}

			y := randomness.Int()
			if 0 == y%2 {
				y = -y
			}

			var c color.Color = img.At(x,y)

			aR,aG,aB,aA := c.RGBA()

			eR := uint32(buffer[0]) * (0xffff/0xff)
			eG := uint32(buffer[1]) * (0xffff/0xff)
			eB := uint32(buffer[2]) * (0xffff/0xff)
			eA := uint32(buffer[3]) * (0xffff/0xff)

			if eR != aR || eG != aG || eB != aB || eA != aA {
				t.Errorf("For test #%d & sub-test #%d, the actual color was not what was expected.", testNumber, subTestNumber)
				t.Logf("(x,y) = (%d,%d)", x,y)
				t.Logf("rgba(%d,%d,%d,%d)", buffer[0], buffer[1], buffer[2], buffer[3])
				t.Logf("EXPECTED (r,g,b,a) = (%d,%d,%d,%d)", eR, eG, eB, eA)
				t.Logf("ACTUAL   (r,g,b,a) = (%d,%d,%d,%d)", aR, aG, aB, aA)
				continue
			}
		}
	}
}

func TestSlice_Bounds(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	for testNumber:=0; testNumber<10; testNumber++ {

		var buffer [rgba32.ByteSize]uint8
		for i:=0; i<len(buffer); i++ {
			buffer[i] = uint8(randomness.Intn(256))
		}

		var rgba rgba32.Slice = rgba32.Slice(buffer[:])

		var img image.Image = rgba

		bounds := img.Bounds()

		for subTestNumber:=0; subTestNumber<50; subTestNumber++ {

			x := randomness.Int()
			if 0 == x%2 {
				x = -x
			}

			y := randomness.Int()
			if 0 == y%2 {
				y = -y
			}

			if x < bounds.Min.X || bounds.Max.Y < x || y < bounds.Min.Y || bounds.Max.Y < y {
				t.Errorf("For test #%d & sub-test #%d, found something outside the bounds.", testNumber, subTestNumber)
				t.Logf("(x,y) = (%d,%d)", x,y)
				continue
			}
		}
	}
}
