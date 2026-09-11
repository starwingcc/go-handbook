# go-handbook

一份可运行的 Go 语言速查表。每个语言特性都拆成独立子包,既能被根 `main.go` 聚合起来一次跑完,也能单独 `go run` 只看某一主题。

适合已经懂编程、想快速过一遍 Go 的读者:读代码 + 改代码 + 跑输出,对照学。

## 前置

- Go 1.25+(用了 `wg.Go`);1.23+(用了 range-over-func)。
- VS Code + Go 扩展(或任意支持 gopls 的编辑器)。

## 如何运行

```bash
go run .                  # 跑全部,按主题顺序输出
go run ./consts/run       # 只跑 consts 一个包(任意包同理)
go test -v -bench=. ./testingdemo  # testingdemo 用 go test 跑(不是 go run)
go build ./...            # 全量编译检查
```

每个 demo 包目录下有个 `run/main.go`,内容仅两行——调用该包的 `Demo()`。这是让"库包"也能单独 `go run` 的标准做法(库包本身不能 run,需经一个 `package main` 入口)。

## 目录与每个包的含义

| 包 | 主题 | 学到什么 |
|----|------|----------|
| `syntax` | 声明 / 返回值 / 指针 / 控制流 / 闭包 / defer / panic | `var`/`:=`、零值、多返回值、命名返回与 naked return、指针解引用、if-init、for 四形态、switch 与 fallthrough、闭包捕获、defer 的 LIFO、panic/recover |
| `consts` | const / iota / 无类型常量 | 批量常量、`iota` 递增、位标志枚举、`_` 跳值、无类型常量按上下文隐式转类型 |
| `types` | 类型别名 vs 定义类型 / 可比较性 | `type T = int`(别名,可隐式互转) vs `type T int`(定义类型,需显式转换);struct 可比较、slice/map/func 不可比较 |
| `builtins` | 内置函数 / 运算符 | `make` vs `new`、`copy`、`clear`、`min`/`max`、swap 多赋值、位运算 `& | ^ &^ << >>` |
| `collection` | 数组 / 切片 / rune-string / map / 标准库 | 数组值语义 vs 切片引用语义、`[]byte`/`[]rune` 与 `range string`、`maps`/`slices` 标准库 |
| `composition` | 结构体嵌入 / 接口嵌入 / 方法提升 | 匿名字段组合、字段与方法提升、接口组合(`Stringer` 嵌入 `Describer`)、接口字段作依赖注入 |
| `iter` | range-over-int / range-over-func | Go 1.23 迭代器:`func(yield func(V) bool)`,提前 break,标准库 `slices.All` 同形态 |
| `geometry` | 结构体 / 方法 / 接口 / 泛型 | 值接收者 vs 指针接收者、Stringer、隐式接口实现、泛型结构体与方法、方法值、`&Point{}` 复合字面量取址、导出 vs 未导出 |
| `nums` | 泛型函数 / 约束 / 变参 | 类型参数、`cmp.Ordered`、`~` 底层类型约束、变参 `...T` 与切片展开 |
| `apperr` | 错误处理 | 哨兵错误、带类型错误、`%w` 包装、`errors.Join`、`errors.Is`/`As` 沿链查找 |
| `conc` | 并发 | goroutine、`wg.Go`、channel(缓冲/方向/select/超时/nil)、`sync.Mutex`、`sync.Once`、`atomic`、build tag(`conc_linux.go`/`conc_other.go`) |
| `ctxdemo` | context | `WithTimeout`、`WithCancel`、`WithValue`,取消/超时/传值 |
| `jsondemo` | encoding/json | marshal/unmarshal、struct tag、`omitempty`、`json:"-"`、`MarshalIndent` |
| `reflectdemo` | reflect 反射 | `TypeOf`/`ValueOf`、`Kind`、遍历字段与 tag、经指针 `Elem()` 设值、方法调用 `MethodByName`、`reflect.New`/`Zero` |
| `testingdemo` | testing 测试 | 表驱动 + `t.Run` 子测试、`t.Parallel`、`t.Helper`、`t.Cleanup`、`t.Skip`、`Example`、`Benchmark`;只能 `go test` 跑 |
| `modules` | module 版本管理 | `debug.ReadBuildInfo` 读模块/依赖/构建设置;go.mod 指令速查(`module`/`require`/`replace`/`exclude`/`retract`/`toolchain`)、MVS、go.sum、go.work |
| `internal/hidden` | internal 包 + blank import | `internal/` 的工具链访问约束、`_` 空白导入触发 `init()` 副作用 |
| `main.go` | 聚合入口 | 只负责按序调用各包 `Demo()`;演示别名导入(`f "fmt"`)、空白导入、internal 导入 |

## 学习路径建议

1. **`syntax` → `consts` → `builtins`**:语言核心,先过一遍。
2. **`types` → `collection` → `composition`**:类型系统与组合。
3. **`geometry` → `nums` → `iter`**:方法、接口、泛型、迭代器。
4. **`apperr` → `conc` → `ctxdemo`**:错误、并发、context——服务端基础。
5. **`jsondemo` → `reflectdemo`**:数据编码、反射。
6. **`testingdemo`**:测试体系,`go test -v -bench=.` 跑。
7. **`modules` + `internal/hidden` + `main.go` 的导入块**:module 版本管理与包机制。

每包都建议:先 `go run ./<pkg>/run` 看输出 → 打开源文件读注释 → 改一两个值再跑,观察输出变化。