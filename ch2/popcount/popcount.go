package popcount

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
	return
} ()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var c byte
	for i := uint64(0); i < 8; i++ {
		c += pc[byte(x>>(i*8))]
	}

	return int(c)
}

func PopCount3(x uint64) int {
	c := 0
	for x > 0 {
		if x & 1 == 1 {
			c++
		}

		x >>= 1
	}

	return c
}

func PopCount4(x uint64) int {
	c := 0
	for x != 0 {
		x = x & (x - 1)
		c++
	}

	return c
}