func twoSum(nums []int, target int) []int {

	for i, v1 := range nums {
		for j, v2 := range nums {
			if (i < j && v1 + v2 == target) {
				return []int{i, j};
			}
		}
	}
	return []int{};
}