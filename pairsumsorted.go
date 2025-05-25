package tripofgo

func PairSumSorted(nums []int, target int) [][]int {

	var (
		left  = 0
		right = len(nums) - 1
		sum   = 0
		pairs [][]int
	)

	// if len from nums is same with 3, iterate all elements
	if len(nums) == 3 {
		for i, v := range nums {
			for in, val := range nums {
				// if index is same, skip add to result slice
				if i == in {
					continue
				} else if v == val {
					pairs = append(pairs, []int{i, in})
				}
			}
		}
	} else {
		for left < right {
			// add left nums value with right nums value
			sum = nums[left] + nums[right]

			//if sum is same with target
			if sum == target {
				pairs = append(pairs, []int{left, right})
				left++
				right--
			}

			// if sum is less than target
			if sum < target {
				left++
			} else {
				// if sum is more than target
				right--
			}
		}
	}

	return pairs
}
