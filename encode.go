package base64encoding

import "github.com/4kills/base64encoding/datatypes"

func (enc Encoder64) encode(b []byte) string {
	return string(bitsToBase64(datatypes.FromBytes(b), enc.valMap))
}

func bitsToBase64(bits datatypes.BitArray, valMap []byte) []byte {
	log64 := 6
	runs := bits.Len() / log64
	remainder := bits.Len() % log64
	overflow := 0
	if remainder != 0 {
		overflow = 1
	}

	str := make([]byte, runs+overflow)

	if remainder != 0 {
		str[0] = valMap[nextNBits(bits, 0, remainder)]
	}

	for i := 0; i < runs; i++ {
		pos := nextNBits(bits, remainder+log64*i, log64)
		str[overflow+i] = valMap[pos] //gets the ASCII of the position in the code
	}
	return str
}

// assert 0 <= n <= 8
func nextNBits(a datatypes.BitArray, idx, n int) byte {
	var b byte
	for i := 0; i < n; i++ {
		bit := a.Get(idx + i)
		if !bit {
			continue
		}
		b |= 0x80 >> i
	}
	return b >> (8 - n)
}
