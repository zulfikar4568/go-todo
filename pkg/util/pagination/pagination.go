package pagination

import "math"

func GetOffset(page int, limit int) int {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}

func ParseMeta[R any](total int, limit int, result []R) (count int, totalPage int) {
	return len(result), int(math.Ceil(float64(total) / float64(limit)))
}
