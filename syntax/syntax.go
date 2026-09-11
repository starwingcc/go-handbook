// Package syntax 演示核心语法：声明/返回值/指针/控制流/闭包/defer/panic。
package syntax

import "fmt"

// var 声明 + 零值
func decl() {
	var i int             // 0
	var s string          // ""
	var ok bool           // false
	var f float64         // 0
	var p *int            // nil
	x := 42               // 短声明 :=（仅函数内）
	_, y := divmod(17, 5) // 空白标识符丢弃某个返回值
	fmt.Printf("decl: i=%d s=%q ok=%v f=%v p=%v x=%d y=%d\n", i, s, ok, f, p, x, y)
}

// 多返回值
func divmod(a, b int) (int, int) { return a / b, a % b }

// 命名返回值：预声明，可 naked return
func namedRet() (q, r int) {
	q, r = 7, 3
	return
}

func pointer() {
	v := 10
	p := &v // *int
	*p = 20 // 解引用修改原值
	fmt.Println("pointer:", v, *p)
}

func control() {
	// if 带初始化语句
	if x := 10; x > 5 {
		fmt.Println("if-init:", x)
	}

	// for 四形态：C 风格 / 仅条件 / 无限 / range
	for range 2 { // 旧写法 for i := 0; i < 2; i++ {
	}
	n := 0
	for n < 2 {
		n++
	}
	for {
		break
	}

	// switch 默认不 fallthrough
	switch day := "Mon"; day {
	case "Sat", "Sun":
		fmt.Println("weekday: weekend")
	default:
		fmt.Println("weekday: weekday")
	}

	// 显式 fallthrough（无条件跳到下一 case，慎用）
	switch 1 {
	case 1:
		fmt.Println("switch: one")
		fallthrough
	case 2:
		fmt.Println("switch: (fallthrough) two")
	}

	// 无表达式 switch = if/else 链
	switch {
	case n > 1:
		fmt.Println("switch-no-expr: >1")
	default:
		fmt.Println("switch-no-expr: <=1")
	}
}

// 闭包：捕获其环境的函数值
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// defer：LIFO，函数返回时执行；常用于清理
func deferDemo() {
	defer fmt.Println("defer: 最后（LIFO）")
	defer fmt.Println("defer: 先出")
	fmt.Println("defer: normal flow")
}

// panic/recover：recover 仅在 defer 中有效
func safeDiv(a, b int) (r int) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover:", e)
			r = 0
		}
	}()
	return a / b // b==0 时 panic
}

func Demo() {
	decl()
	_, _ = namedRet()
	pointer()
	control()
	c := counter()
	fmt.Println("closure:", c(), c(), c())
	deferDemo()
	fmt.Println("recover result:", safeDiv(1, 0))
}
