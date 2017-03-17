package main

import (
	// 	"bytes"
	// 	"encoding/binary"
	// "bytes"
	// "encoding/binary"
	"fmt"
	//	"strings"
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
	switch len(k) { /* all the case statements fall through */

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
	q := []uint8{0x73, 0x69, 0x68, 0x54, 0x20, 0x73, 0x69, 0x20, 0x20, 0x65, 0x68, 0x74, 0x65, 0x6d, 0x69, 0x74, 0x72, 0x6f, 0x66, 0x20, 0x6c, 0x6c, 0x61, 0x20, 0x6f, 0x6f, 0x67, 0x20, 0x65, 0x6d, 0x20, 0x64, 0x6f, 0x74, 0x20, 0x6e, 0x6d, 0x6f, 0x63, 0x20, 0x6f, 0x74, 0x20, 0x65, 0x65, 0x68, 0x74, 0x20, 0x64, 0x69, 0x61, 0x20, 0x20, 0x66, 0x6f, 0x20, 0x69, 0x65, 0x68, 0x74, 0x6f, 0x63, 0x20, 0x72, 0x72, 0x74, 0x6e, 0x75, 0x2e, 0x2e, 0x2e, 0x79}
	// fmt.Printf("%.8x\n", hashword(q, 13))

	// str := "HTTP/1.1 204 No Content\r\n\r\n"

	// reader := strings.NewReader(str)

	// l := reader.Len()
	// n, err := reader.Read([]byte("HTTP"))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("Reader length is : ", len(q))
	fmt.Printf("%.8x \n", hashlittle(q, 47))

	//var pi float64

	// for (i=0, h=0; i<n; ++i) h = hashlittle( k[i], len[i], h);

	// var chunk uint32

	// b := []byte{0x73, 0x69, 0x68, 0x54}
	// buf := bytes.NewReader(b)
	// err := binary.Read(buf, binary.LittleEndian, &pi)
	// if err != nil {
	// 	fmt.Println("binary.Read failed:", err)
	// }
	// fmt.Print(pi)
	// const uint32_t q[18] = {0x73696854, 0x20736920, 0x20656874, 0x656d6974, 0x726f6620, 0x6c6c6120, 0x6f6f6720, 0x656d2064, 0x6f74206e, 0x6d6f6320, 0x6f742065, 0x65687420, 0x64696120, 0x20666f20, 0x69656874, 0x6f632072, 0x72746e75, 0x2e2e2e79};
	// printf("%.8x \n", hashlittle(q, sizeof(q), 47));
	//// e0919a3a
}
