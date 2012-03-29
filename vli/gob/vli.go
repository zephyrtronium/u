/*
This is free software. It comes without any warranty, to the extent
permitted by applicable law. You can redistribute and/or modify it
under the terms of the Do What The Fuck You Want To Public License,
Version 2, as published by Sam Hocevar. See ../COPYING for more
details.
*/

package vli

import (
	"../bits"
	"io"
)

func Read(b io.Reader) (x uint64, err error) {
	p := make([]byte, 1)
	if _, err = b.Read(p); err != nil {
		return
	}
	if p[0] < 128 {
		x = uint64(p[0])
		return
	}
	n := -int(int8(p[0])) // convert to signed, then sign-extend, then negate
	c := 0
	p = make([]byte, n)
	if c, err = b.Read(p); err != nil {
		if err == io.EOF && c == n {
			err = nil
		} else {
			return
		}
	}
	for _, v := range p {
		x = x<<8 | uint64(v) //TODO: decode error if x == 0 && v == 0
	}
	return //TODO: decode error if x < 128 here
}

func ReadSigned(b io.Reader) (x int64, err error) {
	var u uint64
	if u, err = Read(b); err != nil {
		return
	}
	if u&1 == 1 { // complemented
		x = int64(^(u >> 1))
	} else {
		x = int64(u >> 1)
	}
	return
}

func Write(b io.Writer, x uint64) (err error) {
	if x < 128 {
		err = b.Write([]byte{byte(x)})
		return
	}
	b := make([]byte, Length(x)+1)
	for i := len(b) - 1; i > 0; i-- {
		b[i] = byte(x)
		x >>= 8
	}
	b[0] = ^byte(len(b)) + 2
	err = b.Write(b)
	return
}

func WriteSigned(b io.Writer, x int64) error {
	var u uint64
	if x < 0 {
		u = uint64(^x<<1 | 1)
	} else {
		u = uint64(x << 1)
	}
	return Write(b, u)
}

func Length(x uint64) int {
	return (bits.Lg(x)-1)>>3 + 2
}

func LengthSigned(x int64) int {
	var u uint64
	if x < 0 {
		u = uint64(^x << 1)
	} else {
		u = uint64(x << 1)
	}
	return Length(u)
}
