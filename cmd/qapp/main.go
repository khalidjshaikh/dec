package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/bits"
)

func commaize32(n uint32) string {
	s := fmt.Sprintf("%d", n)
	out := make([]byte, 0, len(s)+len(s)/3)
	for i, ch := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, byte(ch))
	}
	return string(out)
}

func main() {
	var seed uint32
	binary.Read(rand.Reader, binary.LittleEndian, &seed)
	var v uint32 = seed

	for range 63 {
		// xorshift32 — deterministic from the seed, so output is reproducible.
		v ^= v << 13
		v ^= v >> 17
		v ^= v << 5

		fmt.Printf("dec: %14s | hex: %08X | bin: %032b\n",
			commaize32(v), v, v)
	}

	fmt.Println("\n--- max uint32 ---")
	allOnes := ^uint32(0)
	fmt.Printf("dec: %14s | hex: %08X | bin: %032b\n",
		commaize32(allOnes), allOnes, allOnes)
	_ = bits.UintSize // ensure "math/bits" import is used
}
