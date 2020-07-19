package rgba32

// RGBA makes Slice fit the Go built-in color.Color interface.
func (receiver Slice) RGBA() (r, g, b, a uint32) {
	r = receiver.R()
	g = receiver.G()
	b = receiver.B()
	a = receiver.A()

	return r, g, b, a
}

