package base64encoding

import (
	"errors"
)

func (enc Encoder64) decode(s string) ([]byte, error) {
	if len(s) < 1 {
		return nil, errors.New("base64decoding error: string is empty")
	}

	bits, err := base64ToBits([]byte(s), []byte(enc.posMap))
	if err != nil {
		return nil, err
	}

	return shift(bits), nil
}

func shift(bits BitArray) []byte {
	l := bits.Len()
	start := l%8
	bs := make([]byte, (l - start)/8)

	idx := 0
	for i := start; i < l; i+=8 {
		var b byte
		for j := 0; j < 8; j++ {
			bit := bits.Get(i+j)
			if !bit {
				continue
			}
			b |= 0x80 >> j
		}

		bs[idx] = b
		idx++
	}
	return bs
}

func base64ToBits(s, posMap []byte) (BitArray, error) {
	bitLen := 6
	bits := NewBitArray(len(s) * bitLen)

	for i, v := range s {
		num, err := findValue(v, posMap)
		if err != nil {
			return BitArray{}, err
		}

		curPart := i * bitLen
		for j := 0; j < bitLen; j++ {
			newBit := false
			if (0x20 >> j) & num > 0 {
				newBit = true
			}

			bits.Set(curPart + j, newBit)
		}
	}

	return bits, nil
}

func findValue(s byte, posMap []byte) (int, error) {
	position := posMap[s]
	idx := int(position - 1)
	if position == 0 {
		return idx, errors.New("base64decoding: semantic: string was invalid, character not found in codeset")
	}
	return idx, nil
}
