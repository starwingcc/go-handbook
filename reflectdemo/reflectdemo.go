// Package reflectdemo 演示 reflect 反射：类型/值、Kind、字段与 tag、设值、方法调用。
package reflectdemo

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	age  int // 未导出字段：反射只读不可设
}

func (u User) Greet() string  { return "hi " + u.Name }
func (u *User) Bump()          { u.ID++ }

func Demo() {
	u := User{ID: 1, Name: "Alice", age: 30}
	v := reflect.ValueOf(u) // 值的反射（不可设）
	t := reflect.TypeOf(u)
	fmt.Printf("type=%s kind=%s value=%v\n", t.Name(), t.Kind(), v.Interface())

	// 遍历结构体字段、类型、tag、是否导出
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)
		// 未导出字段不能 Interface()，否则 panic
		var val any = "(unexported)"
		if f.IsExported() {
			val = fv.Interface()
		}
		fmt.Printf("  field %-4s type=%-8v exported=%-5v tag=%-18q value=%v\n",
			f.Name, f.Type, f.IsExported(), f.Tag.Get("json"), val)
	}

	// 设值：必须经指针的 Elem()，且字段需导出
	vp := reflect.ValueOf(&u)
	vp.Elem().FieldByName("Name").SetString("Bob")
	fmt.Println("after SetString:", u.Name)
	// vp.Elem().FieldByName("age").SetInt(1) // panic：未导出字段不可设

	// 方法调用：值方法用值 v，指针方法用指针 vp
	greet := v.MethodByName("Greet") // 值方法
	fmt.Println("call Greet:", greet.Call(nil)[0].Interface())
	bump := vp.MethodByName("Bump") // 指针方法
	bump.Call(nil)
	fmt.Println("after Bump:", u.ID)

	// 零值与新指针
	fmt.Println("reflect.Zero:", reflect.Zero(t).Interface())
	newp := reflect.New(t) // 返回 *User 指向零值
	fmt.Printf("reflect.New -> type=%s zero=%v\n", newp.Type(), newp.Elem().Interface())

	// 反向取回原值
	back := v.Interface().(User)
	fmt.Println("interface back:", back.Name)
}
