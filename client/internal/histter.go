package histter

import (
	"fmt"
	"slices"
	"strings"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func MakeHistogram[T Number](data []T, symbol byte, height int, max ...T) []string {
	var res []string
	divisions := int(slices.Max(data)) / height
	if len(max) != 0 {
		divisions = int(max[0]) / height
	}

	offset := 1
	curHeight := height
	for curHeight >= 1 {
		str := make([]byte, len(data)+offset)
		str[0] = '|'
		for i := range data {
			if int(data[i])/divisions >= curHeight {
				str[i+offset] = symbol
			} else {
				str[i+offset] = ' '
			}
		}
		curHeight--
		res = append(res, string(str))
	}
	res = append(res, strings.Repeat("=", len(data)+1))

	return res
}

func PrintHistogram(histogram []string) {
	for _, line := range histogram {
		fmt.Println(line)
	}
}
