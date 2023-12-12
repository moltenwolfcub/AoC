package helpers

func IntMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func IntAbs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func LCM(nums ...int) int {
	lcm := 1
	for _, num := range nums {
		if num == 0 {
			panic("can't find LCM of zero")
		}

		lcm = lcm * num / GCD(lcm, num)
	}
	return lcm
}

func GCD(a, b int) int {
	t := 0
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}
