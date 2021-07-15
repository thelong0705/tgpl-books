package popcount

var pc [256]byte

func init() {
	for i := 0; i < 256; i++ {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// CountByLoop exercise 2.4
func CountByLoop(x uint64) (count int) {
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x = x >> 1
	}
	return count
}

// CountByTable example
func CountByTable(x uint64) (count int) {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))],
	)
}

// CountByClearRightMost exercise 2.5
func CountByClearRightMost(x uint64) (count int) {
	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}


7: 0111
4: 0100

7 & 4 = 0100
7 & ans = 0