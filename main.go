package main

import (
	"fmt"
)

func commaize(n uint64) string {
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
	// Print 0, then 1, then left-shift up to 63 times.
	vals := []uint64{0, 1}
	for i := uint(1); i <= 63; i++ {
		vals = append(vals, uint64(1)<<i)
	}

	for _, n := range vals {
		fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b | int64: %+25d\n",
			commaize(n), n, n, int64(n))
	}

	fmt.Println("\n--- all 64 bits set to 1 ---")
	allOnes := ^uint64(0)
	fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b | int64: %+25d\n",
		commaize(allOnes), allOnes, allOnes, int64(allOnes))

	fmt.Println("\n--- 110000...0000 (top two bits set) ---")
	topTwo := uint64(0xC000000000000000)
	fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b | int64: %+25d\n",
		commaize(topTwo), topTwo, topTwo, int64(topTwo))

	fmt.Println("\n--- all 1s (64-bit) ---")
	fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b | int64: %+25d\n",
		commaize(allOnes), allOnes, allOnes, int64(allOnes))

	fmt.Println("\n--- sign bit on (lines 1-63) ---")
	const signBit uint64 = 0x8000000000000000
	for _, n := range vals[:63] {
		m := n | signBit
		fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b | int64: %+25d\n",
			commaize(m), m, m, int64(m))
	}
}