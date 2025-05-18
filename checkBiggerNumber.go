package tripofgo

func CheckBiggerNumber(n ...int) int {
	if len(n) > 3 {
		panic("Not allowed more than 3 numbers")
	}

	max := 0
	for _, v := range n {
		max = v
		if max > v {
			max = v
		}
	}
	return max
}
