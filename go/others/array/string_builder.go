package array

import (
	"fmt"
	"strings"
)

type StringBuilder struct {
	strings []string
}

func (sb *StringBuilder) Append(str string) {
	sb.strings = append(sb.strings, str)
}

func (sb *StringBuilder) ToString() string {
	return strings.Join(sb.strings, "")
}

func StringBuild() {
	sb := &StringBuilder{}
	sb.Append("hello")
	sb.Append(" world")
	fmt.Println(sb.ToString())
}
