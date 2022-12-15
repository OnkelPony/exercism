package variablelengthquantity

import (
	"errors"
	"fmt"
)

const (
	bits   = 7
	msbVal = 128
)

func DecodeVarint(inputs []byte) ([]uint32, error) {
	if inputs[len(inputs)-1] >= msbVal {
		return []uint32{}, errors.New("Ses posral,ne?")
	}
	var result []uint32
	var resParts []byte
	for _, input := range inputs {
		resParts = append(resParts, input)
		if input < msbVal {
			var varInt uint32
			for i, octet := range resParts {
				if octet >= msbVal {
					bsVal := (len(resParts) - i - 1) * bits
					varInt += (uint32(octet) & (msbVal - 1)) << bsVal
				} else {
					varInt += uint32(octet)
					result = append(result, varInt)
					resParts = []byte{}
				}
			}
		}
	}
	return result, nil
}

func EncodeVarint(inputs []uint32) []byte {
	var result []byte
	for _, input := range inputs {
		binLen := len(fmt.Sprintf("%b", input))
		lsB := byte(input % msbVal)
		if binLen <= bits {
			result = append(result, lsB)
			continue
		} else {
			noOcts := (binLen - 1) / bits
			for i := noOcts; i > 0; i-- {
				octet := input>>(i*bits)&(msbVal-1) + msbVal
				result = append(result, byte(octet))
			}
			result = append(result, lsB)
		}
	}
	return result
}
