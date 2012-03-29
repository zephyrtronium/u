// This is free software. It comes without any warranty, to the extent
// permitted by applicable law. You can redistribute and/or modify it
// under the terms of the Do What The Fuck You Want To Public License,
// Version 2, as published by Sam Hocevar. See ../COPYING for more
// details.

// It's alright, though; I have no idea whether I'm doing this right.

// func Lg(x uint64) int
TEXT ·Lg(SB),7,$0
	BSRQ	x+0(FP),AX
	JZ	waszero
	MOVQ	AX,noname+8(FP)
	RET
waszero:
	MOVQ	$-1,noname+8(FP)
	RET
