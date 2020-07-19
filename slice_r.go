package rgba32

// R returns the value for the red channel.
func (receiver Slice) R() (red uint32) {
	if ByteSize != len(receiver)  {
		return 0
	}

	u8 := receiver[OffsetRed]

	u32 := uint32(u8) * (0xffff/0xff)

	return u32
}

