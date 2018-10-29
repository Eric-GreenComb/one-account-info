package ethereum

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

// ErasureHexRight0 ErasureHexRight0
func ErasureHexRight0(input string) (string, error) {
	_buf, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(ErasureRight0(_buf)), nil
}

// ErasureHexLeft0 ErasureHexLeft0
func ErasureHexLeft0(input string) (string, error) {
	_buf, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(ErasureLeft0(_buf)), nil
}

// ErasureRight0 ErasureRight0
func ErasureRight0(buf []byte) []byte {
	index := bytes.IndexByte(buf, 0)
	return buf[0:index]
}

// ErasureLeft0 ErasureLeft0
func ErasureLeft0(buf []byte) []byte {
	index := bytes.LastIndexByte(buf, 0)
	return buf[index+1:]
}

// TrimLeftZeros TrimLeftZeros
func TrimLeftZeros(str string) string {
	return strings.TrimLeft(strings.TrimPrefix(str, "0x"), "0")
}

// PadLeft0 PadLeft0
func PadLeft0(addr string) string {
	return fmt.Sprintf("%064s", addr)
}
