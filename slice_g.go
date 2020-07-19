package rgba32

// G returns the value for the green channel.
func (receiver Slice) G() (green uint32) {
	if ByteSize != len(receiver)  {
		return 0
	}

	u8 := receiver[OffsetGreen]

	u32 := uint32(u8) * (0xffff/0xff)

	return u32
}

