package ch5

func sum(nums ...int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

func myMin(nums ...int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < ans {
			ans = nums[i]
		}
	}
	return ans
}
