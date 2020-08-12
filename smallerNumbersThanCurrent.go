package smaller_numbers

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	//array to keep value
	accuArray := make([]int, 101)
	valueArray := make([]int, 101)
	for _, val := range nums {
		accuArray[val] += 1
	}
	for idx, _ := range valueArray {
		if idx > 0 {
			accuArray[idx] += accuArray[idx-1]
			valueArray[idx] = accuArray[idx-1]
		}
	}
	for idx, _ := range result {
		result[idx] = valueArray[nums[idx]]
	}
	return result
}
