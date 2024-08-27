package slice

func Contains[T string | int](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func ArrayIndexOf[T string | int](slice []T, element T) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
