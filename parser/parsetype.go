package parser

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	TYPE_VARIABLE = iota
	TYPE_ARRAY
	TYPE_LIST
)

var arrayRe = regexp.MustCompile(`[aA]rray\((?P<Type>.+?), *?(?P<Size>\d+?)\)`)
var listRe = regexp.MustCompile(`[lL]ist\((?P<Type>.+?)\)`)

func TypeOf(t string) int {
	if arrayRe.MatchString(t) {
		return TYPE_ARRAY
	} else if listRe.MatchString(t) {
		return TYPE_LIST
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

var ErrInvalidListType = errors.New("invalid list type")

func ParseList(t string) (Type string, err error) {
	matchs := listRe.FindStringSubmatch(t)
	if len(matchs) != 2 {
		err = ErrInvalidListType
		return
	}
	return matchs[1], nil
}
