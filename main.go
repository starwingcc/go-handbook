package main

// ── 导入演示 ────────────────────────────────────────────────────────────
// 1) 别名导入：f "fmt" —— 给包起短名，此处把 fmt 别名为 f
// 2) 空白导入：_ "go-handbook/internal/hidden" —— 仅为触发 init() 副作用，不绑定标识符
// 3) 点导入（此处不用，不推荐）：. "pkg" —— 省略前缀直接用其标识符，易冲突
// 4) 本地子包：module 名(go-handbook) + 目录相对路径
// 5) internal 包：go-handbook/internal/hidden 仅能被 go-handbook 及其子包导入，外部 module 不可
import (
	f "fmt"

	"go-handbook/apperr"
	"go-handbook/builtins"
	"go-handbook/collection"
	"go-handbook/composition"
	"go-handbook/conc"
	"go-handbook/consts"
	"go-handbook/ctxdemo"
	"go-handbook/geometry"
	_ "go-handbook/internal/hidden" // blank import：触发 init 副作用
	"go-handbook/iter"
	"go-handbook/jsondemo"
	"go-handbook/modules"
	"go-handbook/nums"
	"go-handbook/reflectdemo"
	"go-handbook/syntax"
	"go-handbook/types"
)

func main() {
	// main 仅负责按序调用各包的 Demo()；具体逻辑均在子包中。
	runs := []struct {
		name string
		fn   func()
	}{
		{"syntax 声明/返回/指针/控制流/闭包/defer/panic", syntax.Demo},
		{"consts const/iota/无类型常量", consts.Demo},
		{"types 别名 vs 定义类型/可比较性", types.Demo},
		{"builtins make/new/copy/clear/min/max/swap/位运算", builtins.Demo},
		{"collection 数组/切片/rune-string/map/slices&maps", collection.Demo},
		{"composition 结构体嵌入/接口嵌入/方法提升", composition.Demo},
		{"iter range-over-int/range-over-func", iter.Demo},
		{"geometry 结构体/方法/接口/泛型", geometry.Demo},
		{"nums 泛型函数/约束/变参", nums.Demo},
		{"apperr 哨兵/typed/wrap/Join/Is/As", apperr.Demo},
		{"conc goroutine/channel/select/超时/Mutex/atomic/build-tag", conc.Demo},
		{"ctxdemo context 取消/超时/传值", ctxdemo.Demo},
		{"jsondemo marshal/unmarshal/tag", jsondemo.Demo},
		{"reflectdemo reflect 类型/字段/设值/方法", reflectdemo.Demo},
		{"modules build info / go.mod 指令", modules.Demo},
	}
	for _, r := range runs {
		f.Println("\n===", r.name, "===")
		r.fn()
	}
}
