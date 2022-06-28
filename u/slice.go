package u

// 将数组分割成多个数组
// SliceSplit Split a slice into n pieces.
func SliceSplit[T any](src []T, size int) [][]T {
	var (
		dst [][]T
		i   int
	)
	for i+size < len(src) {
		dst = append(dst, src[i:i+size])
		i += size
	}
	dst = append(dst, src[i:])
	return dst
}
