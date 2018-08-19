package fibonacci

import (
	"testing"
	"fmt"
)

func TestFib(t *testing.T) {
	x := fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	y := fibonacci2()
	fmt.Println("====")
	for i := 0; i < 10; i++ {
		fmt.Println(y())
	}
	fmt.Println("====")
	z := fibonacci3()
	printFileContents(z)
}
