package tripofgo

// import "fmt"

func CheckBiggerNumber(n ...int) (max int) {
	if len(n) > 3 {
		return -1
	}

	for _, v := range n {
		if max < v {
			max = v
		}
	}

	return
}
