package rgba32

// B returns the value for the blue channel.
func (receiver Slice) A() (alpha uint32) {
	if ByteSize != len(receiver)  {
		return 0
	}

	u8 := receiver[OffsetAlpha]

	u32 := uint32(u8) * (0xffff/0xff)

	return u32
}

