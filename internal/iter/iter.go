package iter

type MapFunc[V, O any] func(x V) O

// MapSlice creates a new slice from calling a function for every slice element
func MapSlice[V any, O any](val []V, fn MapFunc[V, O]) []O {
	newValues := make([]O, len(val))
	for i, v := range val {
		newValues[i] = fn(v)
	}
	return newValues
}