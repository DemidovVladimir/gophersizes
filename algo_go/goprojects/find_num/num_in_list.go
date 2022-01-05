package num_in_list

func NumInList(nums []int, f int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	for _, v := range nums {
		if v == f {
			return true
		}
	}
	return false
}
