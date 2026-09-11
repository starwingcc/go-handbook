// Package collection 演示数组 vs 切片、rune/string 转换、map、slices/maps 标准库。
package collection

import (
	"fmt"
	"maps" // 泛型 map 工具（Go 1.21+）
	"slices"
)

// 数组 [N]T：定长，值语义（赋值即整体拷贝）
func arrayDemo() {
	var a [3]int            // 零值数组
	b := [3]int{1, 2, 3}    // 字面量
	c := [...]int{4, 5, 6}  // 由元素推断长度
	a = b                    // 数组是值类型：整体拷贝
	a[0] = 99
	fmt.Printf("array: b=%v a=%v（拷贝独立） c=%v len=%d\n", b, a, c, len(c))

	// 数组的 slice 表达式
	fmt.Println("array -> slice:", b[1:])
}

// 切片：引用语义、共享底层数组
func sliceDemo() {
	s := []int{1, 2, 3}
	t := append(s, 4) // 可能扩容到新底层数组
	fmt.Printf("slice: s=%v t=%v cap(s)=%d cap(t)=%d\n", s, t, cap(s), cap(t))
}

// rune / string 转换
func stringDemo() {
	str := "Go语言"

	// []byte：UTF-8 字节序列
	fmt.Printf("bytes: %v len=%d\n", []byte(str), len(str))

	// []rune：Unicode 码点序列
	fmt.Printf("runes: %v len=%d\n", []rune(str), len([]rune(str)))

	// range string 按 rune 遍历：每轮解出一个 rune 及其字节起始偏移
	for i, r := range str {
		fmt.Printf("  range: byte-offset=%d rune=%c (%U)\n", i, r, r)
	}

	// []rune -> string 拼回
	fmt.Println("rune->string:", string([]rune{0x8BED, 0x8A00})) // "语言"
}

// map + slices/maps 标准库
func mapDemo() {
	m := map[string]int{"a": 1, "b": 2}
	n := map[string]int{"a": 1, "b": 2}

	// maps.Equal：逐键比较
	fmt.Println("maps.Equal:", maps.Equal(m, n))

	// maps.Clone / maps.Copy
	clone := maps.Clone(m)
	maps.Copy(clone, map[string]int{"c": 3})
	fmt.Println("after Copy:", clone, "len:", len(clone))

	// slices.Sort / Contains / Compact / Reverse
	s := []int{3, 1, 2, 2, 2}
	slices.Sort(s)
	fmt.Println("sorted:", s, "contains 2:", slices.Contains(s, 2))
	comp := slices.Compact(slices.Clone(s)) // Compact 原地修改，先 Clone 保原数据
	fmt.Println("compact:", comp)
	rev := []int{1, 2, 3}
	slices.Reverse(rev) // 原地反转
	fmt.Println("reverse:", rev)
}

func Demo() {
	arrayDemo()
	sliceDemo()
	stringDemo()
	mapDemo()
}
