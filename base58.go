package encoding

import (
	"bytes"
	"fmt"
	"math/big"
)

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) ([]byte, string) {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	// 被除数/除数=商……余数
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		fmt.Println("mod=", mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	ReverseBytes(result)
	return result, string(result)
}

// Base58Decode 解码
func Base58Decode(input []byte) ([]byte, string) {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded, string(decoded)
}
