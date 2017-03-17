package main

// original code http://burtleburtle.net/bob/c/lookup3.c

import (
	"fmt"
)

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

func hashlittle(k []uint8, initval uint32) uint32 {

	var a, b, c uint32

	/* Set up the internal state */
	a = 0xdeadbeef + uint32(len(k)) + initval
	b = 0xdeadbeef + uint32(len(k)) + initval
	c = 0xdeadbeef + uint32(len(k)) + initval

	/*--------------- all but the last block: affect some 32 bits of (a,b,c) */

	for len(k) > 12 {
		a += uint32(k[0])
		a += uint32(k[1]) << 8
		a += uint32(k[2]) << 16
		a += uint32(k[3]) << 24
		b += uint32(k[4])
		b += uint32(k[5]) << 8
		b += uint32(k[6]) << 16
		b += uint32(k[7]) << 24
		c += uint32(k[8])
		c += uint32(k[9]) << 8
		c += uint32(k[10]) << 16
		c += uint32(k[11]) << 24
		a, b, c = mix(a, b, c)
		k = k[12:]
	}

	/*-------------------------------- last block: affect all 32 bits of (c) */
	switch len(k) {

	case 12:
		c += uint32(k[11]) << 24
		fallthrough
	case 11:
		c += uint32(k[10]) << 16
		fallthrough
	case 10:
		c += uint32(k[9]) << 8
		fallthrough
	case 9:
		c += uint32(k[8])
		fallthrough
	case 8:
		b += uint32(k[7]) << 24
		fallthrough
	case 7:
		b += uint32(k[6]) << 16
		fallthrough
	case 6:
		b += uint32(k[5]) << 8
		fallthrough
	case 5:
		b += uint32(k[4])
		fallthrough
	case 4:
		a += uint32(k[3]) << 24
		fallthrough
	case 3:
		a += uint32(k[2]) << 16
		fallthrough
	case 2:
		a += uint32(k[1]) << 8
		fallthrough
	case 1:
		a += uint32(k[0])
	case 0:
		return c
	}
	a, b, c = final(a, b, c)
	return c
}

func main() {
	q := []uint8{0x73, 0x69, 0x68, 0x54, 0x20, 
		0x73, 0x69, 0x20, 0x20, 0x65, 0x68, 
		0x74, 0x65, 0x6d, 0x69, 0x74, 0x72, 
		0x6f, 0x66, 0x20, 0x6c, 0x6c, 0x61, 
		0x20, 0x6f, 0x6f, 0x67, 0x20, 0x65, 
		0x6d, 0x20, 0x64, 0x6f, 0x74, 0x20, 
		0x6e, 0x6d, 0x6f, 0x63, 0x20, 0x6f, 
		0x74, 0x20, 0x65, 0x65, 0x68, 0x74, 
		0x20, 0x64, 0x69, 0x61, 0x20, 0x20, 
		0x66, 0x6f, 0x20, 0x69, 0x65, 0x68, 
		0x74, 0x6f, 0x63, 0x20, 0x72, 0x72, 
		0x74, 0x6e, 0x75, 0x2e, 0x2e, 0x2e, 
		0x79}
	fmt.Printf("%.8x \n", hashlittle(q, 47))
}
