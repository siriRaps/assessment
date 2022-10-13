// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	//fmt.Println(i.GetValueAsRuneSlice())
	for pos := range i.GetValueAsRuneSlice() {
		//fmt.Println(pos)
		i.TransformRune(pos)
	}
}

type NewString struct {
	index int
	i     int
	s     string
	res   string
}

func (st *NewString) TransformRune(pos int) {
	c := st.s[pos]
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||  (c >= '0' && c <= '9') {
		st.i++
		if st.i%st.index == 0 {
			st.res += strings.ToUpper(string(c))
		} else {
			st.res += strings.ToLower(string(c))
		}
	} else {
		st.res += string(c)
	}
}

func (st NewString) GetValueAsRuneSlice() []rune {
	return []rune(st.s)
}

func NewSkipString(index int, s string) NewString {
	return NewString{
		index: index,
		s:     s,
	}
}

func (s NewString) String() string {
	return s.res
}
func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}
