package tripofgo

type InvalidINput struct {
	msg string
}

func (i *InvalidINput) Error() string {
	return i.msg
}

func AreaOfSquare(a, b int) (int, error) {
	// if a and b is less than 0
	if a <= 0 || b <= 0 {
		return 0, &InvalidINput{msg: "Invalid input, can't count Area"}
	}
	return a * b, nil
}

func AroundOfSquare(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, &InvalidINput{msg: "Invalid input, can't count Around"}
	}
	//return 2 value
	return (2 * b) + (2 * b), nil
}
