package utils

import (
	"cmp"
	"slices"
)

func Remove[T any](arr []T, idx int) []T {
	if idx < 0 || idx >= len(arr) || len(arr) == 0 {
		return arr
	}
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

// From stringslice package
func Intersection[T cmp.Ordered](a, b []T) []T {
	var (
		i, j int
		r    = []T{}
	)
	slices.Sort(a)
	slices.Sort(b)
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			r = append(r, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return r
}

func Union[T cmp.Ordered](a, b []T) []T {
	u := append(a, b...)
	slices.Sort(u)
	v := []T{}
	for i := 0; i < len(u); i++ {
		if i < len(u)-1 && u[i] == u[i+1] {
			continue
		}
		v = append(v, u[i])
	}
	return v
}

// Maximum cliques of a graph using Bron-Kerbosch algorithm
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func MaximumCliques[T cmp.Ordered](r, p, x []T, conns map[T][]T, cliques *[][]T) {
	if len(p) == 0 && len(x) == 0 {
		*cliques = append(*cliques, r)
		return
	}
	for i := 0; i < len(p); i++ {
		node := p[i]
		recR := Union(r, []T{node})
		recP := Intersection(p, conns[node])
		recX := Intersection(x, conns[node])
		MaximumCliques(recR, recP, recX, conns, cliques)
		p = Remove(p, i)
		x = Union(x, []T{node})
		i--
	}
}
