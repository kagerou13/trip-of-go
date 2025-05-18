package tripofgo

func CheckPosNeg(n int) (result string) {
	if n == 0 {
		result = "zero"
	} else if n > 0 {
		result = "positive"
	} else {
		result = "negative"
	}

	return
}
