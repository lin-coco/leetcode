package backtrack

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dig := make([]byte, len(digits))
	for i := 0; i < len(digits); i++ {
		dig[i] = byte(digits[i])
	}
	res = nil
	paths = nil
	tracking(0, dig)
	return res
}

var res []string
var paths []byte

func tracking(i int, dig []byte) {
	if len(dig) == i {
		res = append(res, string(paths))
		return
	}
	for j := 0; j < len(list[dig[i]]); j++ {
		paths = append(paths, list[dig[i]][j])
		tracking(i+1, dig)
		paths = paths[:len(paths)-1]
	}
}

var list = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}
