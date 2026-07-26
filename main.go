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
		fmt.Printf("dec: %25s | hex: 0x%016X | bin: %064b\n", commaize(n), n, n)
	}
}