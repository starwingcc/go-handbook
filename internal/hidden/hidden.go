// Package hidden 演示 internal 包（仅 go-handbook/* 可导入）与 blank import 副作用。
// 放在 internal/ 下：只有路径以 go-handbook 为根的包才能导入它。
package hidden

import "fmt"

func init() {
	fmt.Println("[internal/hidden] init 执行（main 用 blank import 触发的副作用）")
}

// Loaded 仅供说明：具名导入时可访问 internal 包的导出标识符。
var Loaded = true
