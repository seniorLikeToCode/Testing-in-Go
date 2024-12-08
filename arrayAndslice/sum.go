package array_slice

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsArr ...[]int) []int {
	sum := make([]int, 0)

	for _, nums := range numsArr {
		sum = append(sum, Sum(nums))
	}

	return sum
}

func SumAllTails(numsArr ...[]int) []int {
	var sums []int
	for _, nums := range numsArr {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
