// Package builtins 演示内置函数与运算符。
package builtins

import "fmt"

// make vs new
func makeNew() {
	// make：仅 slice/map/chan，返回初始化（非零）值
	s := make([]int, 3) // len=3
	m := make(map[string]int)
	ch := make(chan int, 1)
	_ = ch
	// new：任意类型，返回 *T（指向零值的指针）
	p := new(int) // *int, *p == 0
	*p = 42
	fmt.Printf("make: s=%v m=%v | new: *p=%d\n", s, m, *p)
}

// copy / clear（Go 1.21+ clear）
func copyClear() {
	src := []int{1, 2, 3}
	dst := make([]int, 3)
	n := copy(dst, src) // 返回复制元素数
	fmt.Println("copy:", dst, "n=", n)

	m := map[string]int{"a": 1, "b": 2}
	clear(m) // 清空 map（也适用于 slice，置零元素）
	fmt.Println("clear map:", m)
}

// min / max（Go 1.21+ 内置，支持 cmp.Ordered）
func minmax() {
	fmt.Println("min:", min(3, 1, 2), "max:", max(3, 1, 2))
	fmt.Println("min str:", min("b", "a", "c"))
}

// swap / 多赋值
func swap() {
	a, b := 1, 2
	a, b = b, a // 交换
	fmt.Println("swap:", a, b)

	// 多返回值直接赋给多变量
	x, y := divmod(17, 5)
	fmt.Println("multi-assign:", x, y)
}

func divmod(a, b int) (int, int) { return a / b, a % b }

// 位运算
func bitwise() {
	a, b := 0b1100, 0b1010
	fmt.Printf("AND   %b & %b = %b\n", a, b, a&b)
	fmt.Printf("OR    %b | %b = %b\n", a, b, a|b)
	fmt.Printf("XOR   %b ^ %b = %b\n", a, b, a^b)
	fmt.Printf("AND-NOT %b &^ %b = %b\n", a, b, a&^b)
	fmt.Printf("SHL   %b << 1 = %b\n", a, a<<1)
	fmt.Printf("SHR   %b >> 1 = %b\n", a, a>>1)
	fmt.Printf("NOT   ^%b = %b\n", a, ^a)
}

func Demo() {
	makeNew()
	copyClear()
	minmax()
	swap()
	bitwise()
}
