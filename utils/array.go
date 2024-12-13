package utils

import "slices"

func Remove[T any](arr []T, idx int) []T {
	reduced := make([]T, len(arr)-1)
	copy(reduced, arr[:idx])
	copy(reduced[idx:], arr[idx+1:])
	return reduced
}

func AppendUnique[T comparable](slice []T, elem ...T) []T {
	res := slice
	for _, el := range elem {
		if slices.Contains(res, el) {
			continue
		}
		res = append(res, el)
	}
	return res
}

func Filter[T any](arr []T, f func(T) bool) []T {
	var res []T
	for _, el := range arr {
		if f(el) {
			res = append(res, el)
		}
	}
	return res
}
