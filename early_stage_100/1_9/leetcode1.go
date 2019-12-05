// 两数之和
func twoSum(nums []int, target int) []int {
    for i, v := range nums {
        dst := target - v
        for j := i + 1; j < len(nums); j++ {
            if dst == nums[j] {
                return []int{i, j}
            }
        }
    }
    return nil
}

