package nameconv

import "strings"

func Snake2Pascal(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	bs := []byte(s)
	if 'a' <= bs[0] && bs[0] <= 'z' {
		sb.WriteByte(bs[0] - 'a' + 'A')
		bs = bs[1:]
	}
	for len(bs) > 0 {
		if bs[0] == '_' {
			if len(bs) > 1 && 'a' <= bs[1] && bs[1] <= 'z' {
				sb.WriteByte(bs[1] - 'a' + 'A')
				bs = bs[2:]
			} else {
				sb.WriteByte('_')
				bs = bs[1:]
			}
		} else {
			sb.WriteByte(bs[0])
			bs = bs[1:]
		}
	}
	return sb.String()
}
