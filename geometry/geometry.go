// Package geometry 演示结构体、方法、接口、泛型在子包中的定义与导出。
package geometry

import "fmt"

// Point —— 导出类型（首字母大写）
type Point struct {
	X, Y float64
}

// 内部类型（小写）包外不可访问
type circle struct{ R float64 }

// 值接收者方法
func (p Point) Distance() float64 { return p.X*p.X + p.Y*p.Y }

// 指针接收者方法（可修改原值，*Point 与 Point 共享方法集）
func (p *Point) Move(dx, dy float64) { p.X += dx; p.Y += dy }

// Stringer —— 实现 String() string，供 fmt 打印时格式化
func (p Point) String() string { return fmt.Sprintf("(%.1f,%.1f)", p.X, p.Y) }

// Shape 接口：隐式实现（无 implements 关键字）
type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }
type Square struct{ S float64 }

func (c Circle) Area() float64 { return 3.14159 * c.R * c.R }
func (s Square) Area() float64 { return s.S * s.S }

// NewCircle 构造函数约定：NewXxx。内部类型 circle 不能包外构造，
// 但可通过构造函数返回其导出形态。
func NewCircle(r float64) Circle { return Circle{R: r} }

// 泛型结构体与方法 —— 类型参数 T
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
	var z T
	n := len(s.items)
	if n == 0 {
		return z, false
	}
	v := s.items[n-1]
	s.items = s.items[:n-1]
	return v, true
}

// 方法值：绑定接收者后的函数值，可作为一等值传递
func methodValue() {
	p := Point{X: 3, Y: 4}

	// 指针接收者方法值：p 可取址，move 绑定 &p
	move := p.Move // 类型 func(float64, float64)
	move(1, 1)      // 等价于 (&p).Move(1,1)
	fmt.Println("method value (ptr recv):", p)

	// 值接收者方法值：捕获的是 p 当时的副本
	dist := p.Distance // 类型 func() float64
	fmt.Println("method value (val recv) dist:", dist())
}

// 复合字面量取址：&T{...} 直接得到 *T
func addrOfLiteral() {
	pp := &Point{X: 5, Y: 6} // *Point，无需先声明变量再取址
	pp.Move(1, 1)
	fmt.Println("&Point{} -> *Point:", *pp)
}

func Demo() {
	p := Point{1, 2}       // 位置式字面量
	q := Point{X: 3, Y: 4} // 键值式字面量
	p.Move(1, 1)
	fmt.Println("struct:", p, q, "dist:", p.Distance())

	// 接口多态
	var shapes []Shape = []Shape{NewCircle(2), Square{S: 3}}
	for _, s := range shapes {
		fmt.Println("shape area:", s.Area())
	}

	// 泛型结构体
	st := &Stack[int]{}
	st.Push(10)
	v, _ := st.Pop()
	fmt.Println("stack:", v)

	methodValue()
	addrOfLiteral()

	// _ = circle{R: 1} // 编译错误：cannot refer to unexported name
}
