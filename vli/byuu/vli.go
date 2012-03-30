/*
This is free software. It comes without any warranty, to the extent
permitted by applicable law. You can redistribute and/or modify it
under the terms of the Do What The Fuck You Want To Public License,
Version 2, as published by Sam Hocevar. See ../COPYING for more
details.
*/

package vli

import (
	"github.com/zephyrtronium/u/bits"
	"io"
)

func Read(b io.Reader) (x uint64, err error) {
	var n, shift int
	for {
		p := []byte{0}
		if n, err = b.Read(p); n != 1 {
			return
		}
		x += uint64(p[0]&0x7f) * shift
		if p[0]&0x80 == 0x80 {
			return
		}
		shift <<= 7
		x += shift
	}
	panic("unreachable")
}

func Write(b io.Writer, x uint64) (err error) {
	p := make([]byte, 0, 10)
	for x > 0x7f {
		p = append(p, byte(x&0x7f))
		x >>= 7
		x-- // what the duck is this. all it does is slow down the algorithms.
	}
	p = append(p, byte(x|0x80))
	_, err = b.Write(p)
	return
}

func Length(x uint64) int {
	b := bits.Lg(x)
	if b%7 == 0 {
		return b / 7
	}
	return b/7 + 1
}
