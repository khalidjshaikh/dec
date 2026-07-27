package main

import (
	"fmt"
	"math/rand"
	"time"
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
	for {
		n := rand.Uint64()
		fmt.Printf("dec: %26s | hex: 0x%016X | bin: %064b\n",
			commaize(n), n, n)
		time.Sleep(time.Second)
	}
}
