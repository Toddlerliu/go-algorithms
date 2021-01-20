package fibonacci

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fibonacci1() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

// 1,1,2,3,5,8,13...
//     a,b
//       a,b
func fibonacci2() func() int {
	x1, x2 := 0, 1
	return func() int {
		// 重新赋值
		x1, x2 = x2, (x1 + x2)
		return x1
	}
}

func fibonacci3() fibGen {
	x1, x2 := 0, 1
	return func() int {
		// 重新赋值
		x1, x2 = x2, (x1 + x2)
		return x1
	}
}

// 自顶向下
// 时间空间：o(n)
func fibonacci4(n int) int {
	if n <= 2 {
		return 1
	}
	s := make([]int, n+1)
	s[1], s[2] = 1, 1
	return _fibonacci4(n, s)
}

func _fibonacci4(n int, s []int) int {
	if s[n] == 0 {
		s[n] = _fibonacci4(n-1, s) + _fibonacci4(n-2, s)
	}
	return s[n]
}

// 自底向上
// 时间空间：o(n)
func fibonacci5(n int) int {
	if n <= 2 {
		return 1
	}
	s := make([]int, n+1)
	s[1], s[2] = 1, 1
	for i := 3; i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}

// 滚动数组
func fibonacci6(n int) int {
	if n <= 2 {
		return 1
	}
	s := make([]int, 2)
	s[0], s[1] = 1, 1
	for i := 0; 3 <= n; i++ {
		//s[0], s[1] = s[1], s[0]+s[1]
		s[i%2] = s[(i-1)%2] + s[(i-2)%2]
	}
	return s[1]
}

type fibGen func() int

// 函数实现接口
func (g fibGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
