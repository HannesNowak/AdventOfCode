package utils

func Remove[T comparable](arr []T, idx int) []T {
	reduced := make([]T, len(arr)-1)
	copy(reduced, arr[:idx])
	copy(reduced[idx:], arr[idx+1:])
	return reduced
}
