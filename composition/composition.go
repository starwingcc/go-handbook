// Package composition 演示结构体嵌入（组合）与接口嵌入、字段/方法提升。
package composition

import "fmt"

// 基础类型
type Base struct {
	ID int
}

func (b Base) Describe() string { return fmt.Sprintf("Base#%d", b.ID) }
func (b *Base) Bump()             { b.ID++ }

// 结构体嵌入（匿名字段）：Base 的字段与方法被"提升"到 Derived
type Derived struct {
	Base        // 匿名字段 → 嵌入
	Name string
}

// 接口嵌入：组合既有接口
type Stringer interface {
	String() string
}
type Describer interface {
	Stringer     // 嵌入接口
	Describe() string
}

// 实现接口（值接收者）
func (d Derived) String() string { return d.Name }

func structEmbed() {
	d := Derived{Base: Base{ID: 1}, Name: "foo"}
	// 提升的字段/方法可直接访问
	fmt.Println("embedded field ID:", d.ID)        // = d.Base.ID
	fmt.Println("embedded method:", d.Describe())  // = d.Base.Describe()
	d.Bump()                                        // 提升的指针方法（d 取址调用）
	fmt.Println("after Bump:", d.ID)

	// 满足组合接口
	var des Describer = d
	fmt.Println("Describer:", des.String(), des.Describe())
}

// 嵌入接口作为字段（依赖注入 / mock 友好）
type Service struct {
	Repo Stringer // 持有接口，而非具体类型
}

func interfaceField() {
	s := Service{Repo: Derived{Name: "injected"}}
	fmt.Println("service:", s.Repo.String())
}

func Demo() {
	structEmbed()
	interfaceField()
}
