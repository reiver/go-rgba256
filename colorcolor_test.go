package rgba256_test

import (
	"github.com/reiver/go-rgba256"

	"image/color"

	"testing"
)

func TestImageImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x color.Color = rgba256.Slice{}

	if nil == x {
		t.Error("This should never happen")
	}
}
