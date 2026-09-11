// Package iter 演示 range-over-int 与 range-over-func（Go 1.23 迭代器）。
package iter

import "fmt"

func rangeInt() {
	// range-over-int：for i := range N
	for i := range 3 {
		fmt.Print(i, " ") // 0 1 2
	}
	fmt.Println()
}

// range-over-func：对返回迭代函数 range
// 标准形式：func(yield func(V) bool) 或 func(yield func(k, v) bool)
func rangeFunc() {
	each := func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i) { // yield 返回 false 表示提前停止
				return
			}
		}
	}
	for v := range each {
		fmt.Print(v, " ") // 1 2 3
	}
	fmt.Println()
}

// 提前 break：yield 收到 false
func rangeFuncBreak() {
	each := func(yield func(int) bool) {
		for i := 1; i <= 100; i++ {
			if !yield(i) {
				return
			}
		}
	}
	count := 0
	for v := range each {
		count++
		if v >= 2 {
			break // yield 返回 false，迭代器停止
		}
	}
	fmt.Println("broke after", count, "items")
}

// 键值迭代器：func(yield func(k int, v string) bool)
func all2(s []string) func(func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// 标准库 slices.All / maps.All 即 range-over-func 风格
func rangeStd() {
	s := []string{"a", "b", "c"}
	for i, v := range all2(s) {
		fmt.Printf("(%d:%s) ", i, v)
	}
	fmt.Println()
}

func Demo() {
	rangeInt()
	rangeFunc()
	rangeFuncBreak()
	rangeStd()
}
