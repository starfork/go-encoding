package encoding

import (
	"bytes"
	"math/big"
)

var base32Bytes = []byte("123456789abcdefghijklmnopqrstuvwxyz")

func Base35Encode(input string) ([]byte, string) {
	x := big.NewInt(0).SetBytes([]byte(input))
	base := big.NewInt(35)
	//fmt.Println(base)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	// 被除数/除数=商……余数
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base32Bytes[mod.Int64()])
		//fmt.Println(result)
	}
	ReverseBytes(result)
	return result, string(result)
}

// Base35Decode 解码
func Base35Decode(input []byte) ([]byte, string) {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base32Bytes, b)
		result.Mul(result, big.NewInt(35))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base32Bytes[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded, string(decoded)
}
