package strings

import (
	"testing"
	"fmt"
)

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	s := "中国helloworld中国"
	num:=lengthOfNonRepeatingSubStr(s)
	fmt.Printf("max length is %d ",num)
}
