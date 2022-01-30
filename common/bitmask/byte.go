package bitmask

type Byte byte

func (b Byte) Has(bb Byte) bool {
	return (b & bb) != 0
}

func (b *Byte) Set(bb Byte) {
	*b |= bb
}

func (b *Byte) Clear(bb Byte) {
	*b &= ^bb
}

func (b *Byte) Toggle(bb Byte) {
	*b ^= bb
}
