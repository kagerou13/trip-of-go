package tripofgo

func CheckBiggerNumber(n ...int) (max int) {
	if len(n) > 3 {
		panic("Not allowed more than 3 numbers")
	}

	for _, v := range n {
		if max > v {
			max = v
		}
	}
	return
}
