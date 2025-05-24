package tripofgo

func PairSumSorted(nums []int, target int) [][]int {

	var (
		left  = 0
		right = len(nums) - 1
		sum   = 0
		pairs [][]int
	)

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
	return pairs
}
