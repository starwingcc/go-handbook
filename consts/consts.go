// Package consts 演示常量、iota、无类型常量。
package consts

import "fmt"

// 无类型常量：按上下文隐式转类型
const Pi = 3.14159
const Lang = "Go" // 无类型字符串常量

// 批量声明：iota 在每个 const 块从 0 起，每行递增
const (
	Sun = iota // 0
	Mon        // 1（省略表达式，重复上一行模式）
	Tue        // 2
	Wed
	Thu
	Fri
	Sat
)

// 位标志枚举：iota 配合位运算
type Permission uint8

const (
	Read Permission = 1 << iota // 1<<0 = 1
	Write                       // 1<<1 = 2
	Execute                     // 1<<2 = 4
)

// 用 _ 占位跳过值
const (
	_  = iota           // 0 忽略
	KB = 1 << (10 * iota) // 1<<10
	MB                     // 1<<20
	GB                     // 1<<30
)

func Demo() {
	fmt.Println("Pi:", Pi, "Lang:", Lang)
	fmt.Println("days Sun..Sat:", Sun, Mon, Tue, Wed, Thu, Fri, Sat)
	fmt.Printf("perm: Read=%d Write=%d Execute=%d | all=%d\n", Read, Write, Execute, Read|Write|Execute)
	fmt.Println("sizes KB/MB/GB:", KB, MB, GB)

	// 无类型常量高精度，可赋给不同具体类型
	var f64 float64 = Pi
	var f32 float32 = Pi
	fmt.Printf("untyped Pi -> float64=%v float32=%v\n", f64, f32)
}
