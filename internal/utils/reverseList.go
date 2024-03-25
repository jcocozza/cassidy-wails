package utils

// Reverse a list of any type
func ReverseList[T any](list []T) []T {
	length := len(list)
	for i := 0; i < length/2; i++ {
		list[i], list[length-i-1] = list[length-i-1], list[i]
	}
	return list
}