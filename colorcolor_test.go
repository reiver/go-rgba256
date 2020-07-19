package rgba32_test

import (
	"github.com/reiver/go-rgba32"

	"image/color"

	"testing"
)

func TestImageImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x color.Color = rgba32.Slice{}

	if nil == x {
		t.Error("This should never happen")
	}
}
