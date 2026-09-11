// Package jsondemo 演示 encoding/json：序列化、反序列化、struct tag。
package jsondemo

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"` // 零值时省略
	Email string `json:"-"`              // 永不导出
	Role  string `json:"role,omitempty"`
}

func Demo() {
	a := Account{ID: 1, Name: "Alice", Role: "admin"}
	b, _ := json.Marshal(a)
	fmt.Println("marshal:", string(b))

	// omitempty：Name 为空则不出现
	a2 := Account{ID: 2}
	b2, _ := json.Marshal(a2)
	fmt.Println("marshal omitempty:", string(b2))

	// 反序列化
	var got Account
	json.Unmarshal([]byte(`{"id":7,"name":"Bob","role":"user"}`), &got)
	fmt.Printf("unmarshal: %+v\n", got)

	// MarshalIndent 缩进
	pretty, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println("indent:\n" + string(pretty))
}
