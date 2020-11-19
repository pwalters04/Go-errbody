package lib_strings

import (
	"strings"
)

// SimpleString... is a holder for string array
type SimpleString struct {
	BasicString string
}

func SetString(myString string) *SimpleString {
	yourString := &SimpleString{
		BasicString: myString,
	}
	return yourString
}

func (s SimpleString) GetLength() int {
	l := len(s.BasicString)
	return l
}

func (s SimpleString) SplitString_Commas() []string {
	split := strings.Split(s.BasicString, ",")
	return split
}
