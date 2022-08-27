package protocol

import (
	"strconv"
	"strings"
)

func ToIpv4(bs []byte) string {
	s := make([]string, 0, 4)
	for _, b := range bs {
		s = append(s, strconv.Itoa(int(b)))
	}
	return strings.Join(s, ".")
}

func ToPort(bs []byte) string {
	return strconv.Itoa((int(bs[0]) << 8) + int(bs[1]))
}
