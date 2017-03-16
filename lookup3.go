package main

import (
	// 	"bytes"
	// 	"encoding/binary"
	"fmt"
)

// #define hashsize(n) ((uint32_t)1<<(n))
// #define hashmask(n) (hashsize(n)-1)
func rot(x, k uint32) uint32 {
	return (((x) << (k)) | ((x) >> (32 - (k))))
}

func mix(a, b, c uint32) (uint32, uint32, uint32) {

	a -= c
	a ^= rot(c, 4)
	c += b

	b -= a
	b ^= rot(a, 6)
	a += c

	c -= b
	c ^= rot(b, 8)
	b += a

	a -= c
	a ^= rot(c, 16)
	c += b

	b -= a
	b ^= rot(a, 19)
	a += c

	c -= b
	c ^= rot(b, 4)
	b += a

	return a, b, c
}

func final(a, b, c uint32) (uint32, uint32, uint32) {

	c ^= b
	c -= rot(b, 14)

	a ^= c
	a -= rot(c, 11)

	b ^= a
	b -= rot(a, 25)

	c ^= b
	c -= rot(b, 16)

	a ^= c
	a -= rot(c, 4)

	b ^= a
	b -= rot(a, 14)

	c ^= b
	c -= rot(b, 24)

	return a, b, c
}

func hashword(k []uint32, initval uint32) uint32 {
	/* Set up the internal state */
	var a, b, c uint32
	a = 0xdeadbeef + ((uint32(len(k))) << 2) + initval
	b = a
	c = b

	/*------------------------------------------------- handle most of the key */

	for len(k) > 3 {
		fmt.Printf("1: % x % x % x \n", k[0], k[1], k[2])
		a += k[0]
		b += k[1]
		c += k[2]
		a, b, c = mix(a, b, c)
		k = k[3:]
	}
	fmt.Printf(">>: % x % x % x \n", a, b, c)

	/*------------------------------------------- handle the last 3 uint32_t's */

	/* all the case statements fall through */
	switch k_len := len(k); k_len {
	case 3:
		c += k[2]
		fmt.Printf("k[2] >>: % x\n", k[2])
		fallthrough
	case 2:
		b += k[1]
		fmt.Printf("k[1] >>: % x\n", k[1])
		fallthrough
	case 1:
		a += k[0]
		fmt.Printf("k[0] >>: % x\n", k[0])
		a, b, c = final(a, b, c)
	case 0:
		break
	}

	/*------------------------------------------------------ report the result */
	return c
}

func main() {
	q := []uint32{0x73696854, 0x20736920, 0x20656874, 0x656d6974, 0x726f6620, 0x6c6c6120, 0x6f6f6720, 0x656d2064, 0x6f74206e, 0x6d6f6320, 0x6f742065, 0x65687420, 0x64696120, 0x20666f20, 0x69656874, 0x6f632072, 0x72746e75, 0x2e2e2e79}
	fmt.Printf("%.8x\n", hashword(q, 13))
}
