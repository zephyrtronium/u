/*
This is free software. It comes without any warranty, to the extent
permitted by applicable law. You can redistribute and/or modify it
under the terms of the Do What The Fuck You Want To Public License,
Version 2, as published by Sam Hocevar. See ../COPYING for more
details.
*/

package bits

// Find the position of the highest set bit of x.
func Lg(x uint64) int

func lg(x uint64) int {
	var i int
	for i = 0; x != 0; i++ {
		x >>= 1
	}
	return i - 1
}
