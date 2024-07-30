package slices

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var r []T
	for _, v := range slice {
		if predicate(v) {
			r = append(r, v)
		}
	}
	return r
}
