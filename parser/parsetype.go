package parser

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	TYPE_VARIABLE = iota
	TYPE_ARRAY
)

var arrayRe = regexp.MustCompile(`[aA]rray\((?P<Type>.*?), *?(?P<Size>\d+?)\)`)

func TypeOf(t string) int {
	if arrayRe.MatchString(t) {
		return TYPE_ARRAY
	}
	return TYPE_VARIABLE
}

var ErrInvalidArrayType = errors.New("invalid array type")

func ParseArray(t string) (Type string, Size int, err error) {
	matchs := arrayRe.FindStringSubmatch(t)
	if len(matchs) != 3 {
		err = ErrInvalidArrayType
		return
	}
	size, err := strconv.Atoi(matchs[2])
	return matchs[1], size, err
}
