package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
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
	var sum byte
	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}

	return int(sum)
}

func PopCount3(x uint64) int {
	var sum int
	for i := uint(0); i < 64; i++ {
		if (x>>i)&1 == 1 {
			sum++
		}
	}

	return sum
}

func PopCount4(x uint64) int {
	var count int
	for {
		tmp := x & (x - 1)
		if tmp == x {
			break
		}
		x = tmp
		count++
	}
	return count
}