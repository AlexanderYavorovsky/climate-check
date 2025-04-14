package histter

import (
	"slices"
	"strings"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func MakeHistogram[T Number](data []T, symbol byte, height ...T) []string {
	res := []string{}

	var scale T = 1
	cur_height := slices.Max(data)
	if len(height) != 0 {
		scale = height[0] / cur_height
	}

	cur_height *= scale

	offset := 1
	for cur_height >= 1 {
		str := make([]byte, len(data)+offset)
		str[0] = '|'
		for i := range data {
			if data[i]*scale >= cur_height {
				str[i+offset] = symbol
			} else {
				str[i+offset] = ' '
			}
		}
		cur_height--
		res = append(res, string(str))
	}
	res = append(res, strings.Repeat("=", len(data)+1))

	return res
}
