package redpackage

import (
	"testing"
	"fmt"
)

func TestRedPackage(t *testing.T) {
	totalMoney := 100.0
	totalSize := 10
	redpackage := NewRedPackage(totalMoney, totalSize)
	s := make([]float64,totalSize)
	sum := 0.0
	for i := 0; i < totalSize; i++ {
		money := redpackage.GetMoney()
		s[i] = money
		sum += money
	}
	fmt.Println(s)
	fmt.Println(sum)
}
