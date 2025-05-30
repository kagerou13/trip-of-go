package tripofgo

func ShiftZeroToEnd(nums []int) []int {
	/*
		Example of input and output
		Input: nums = [0, 1, 0, 3, 2]
		Output: [1, 3, 2, 0, 0]
	*/

	// iterate from zero index to 2 last index
	for i := 0; i < len(nums)-1; i++ {
		//iterate from first index to last index
		for j := i + 1; j < len(nums); j++ {
			// check what current element is 0 or not
			if nums[i] == 0 {
				//swap current element with next element
				nums[i], nums[j] = nums[j], nums[i]
				continue // skip next iteration
			}
		}
	}
	return nums
}
