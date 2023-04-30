package arraySum

func ArraySum(array []int) int {
	if len(array) == 1 {
		return array[0]
	}

	return array[0] + ArraySum(array[1:])
}
