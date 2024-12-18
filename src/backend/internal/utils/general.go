package utils

func MakePointer[T any](v T) *T {
	return &v
}
