package rgba32

// RGBA makes Slice fit the Go built-in color.Color interface.
func (receiver Slice) RGBA() (r, g, b, a uint32) {
	if ByteSize != len(receiver)  {
		return 0,0,0,0
	}

	u8r := receiver[indexRed]
	u8g := receiver[indexGreen]
	u8b := receiver[indexBlue]
	u8a := receiver[indexAlpha]

	r = uint32(u8r) * (0xffff/0xff)
	g = uint32(u8g) * (0xffff/0xff)
	b = uint32(u8b) * (0xffff/0xff)
	a = uint32(u8a) * (0xffff/0xff)

	return r, g, b, a
}

