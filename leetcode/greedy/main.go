func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int
	for i := 0; i < len(nums1); i++ {
		check := -1
		value := -1
		answer := -1
		for j := 0; j < len(nums2); j++ {
			if check != -1 {
				if value < nums2[j] {
					answer = nums2[j]
					break
				}
			}
			if nums2[j] == nums1[i] {
				check = j
				value = nums2[j]
			}
		}

		if answer == 0 {
			result = append(result, -1)
		} else {
			result = append(result, answer)
		}
	}

	return result
}