package sumnums

func Sum(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	var c int
	for _, v := range nums {
		c = c + v
	}
	return c
}
