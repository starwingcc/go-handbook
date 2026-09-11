// Package types 演示类型别名 vs 定义类型、可比较性。
package types

import "fmt"

// 类型别名（=）：与原类型完全相同，可隐式互转。
type Celsius = float64

// 定义类型（无 =）：底层类型相同但已是新类型，需显式转换。
type Kelvin float64

func aliasVsDefined() {
	var c Celsius = 36.5
	var raw float64 = c // 别名：可直接赋值
	_ = raw

	var k Kelvin = 273.15
	// var bad float64 = k // 编译错误：不能把 Kelvin 隐式赋给 float64
	var ok float64 = float64(k) // 需显式转换
	fmt.Printf("alias Celsius=%v float64=%v | defined Kelvin=%v 转后=%v\n", c, raw, k, ok)
}

// 可比较性：struct 全字段可比较则可比较；slice/map/func 不可比较
type Point struct{ X, Y int }

func comparableDemo() {
	a := Point{1, 2}
	b := Point{1, 2}
	fmt.Println("struct == :", a == b) // true

	// slice 不可比较，只能与 nil 比
	s := []int{1, 2}
	// _ = s == s // 编译错误
	fmt.Println("slice == nil:", s == nil)

	// map/func 同理不可比较
	m := map[string]int{}
	fmt.Println("map == nil:", m == nil)
}

func Demo() {
	aliasVsDefined()
	comparableDemo()
}
