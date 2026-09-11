// Package nums 演示泛型函数、类型约束、变参。
package nums

import (
	"cmp"
	"fmt"
)

// Min 类型参数 + 约束（cmp.Ordered = ~int|~float|~string）
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 用 ~ 匹配底层类型 —— 可接受具名类型
type Number interface {
	~int | ~int64 | ~float64
}

func Double[T Number](v T) T { return v * 2 }

// Sum 变参（...T）在函数内表现为切片
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Demo() {
	fmt.Println("generic:", Min(3, 7), Min("b", "a"), Double(5), Double(1.5))
	fmt.Println("variadic:", Sum(1, 2, 3), Sum([]int{4, 5}...)) // 展开切片
}
