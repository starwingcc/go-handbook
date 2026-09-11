// Package modules 演示 module 版本管理：运行时读取 build info，
// 并以包注释形式给出 go.mod 各指令的速查。
//
// go.mod 指令速查：
//
//	module go-handbook                          // 模块路径（导入根前缀）
//	go 1.26.5                               // 要求的语言版本
//	toolchain go1.26.5                      // 显式指定工具链
//	require github.com/x/y v1.2.3           // 直接依赖
//	require github.com/x/y v1.2.3 // indirect  // 间接依赖
//	replace github.com/x/y => ../local      // 本地替换（调试常用）
//	replace github.com/x/y v1.0.0 => github.com/x/y v1.2.0  // 换版本
//	exclude github.com/x/y v1.3.0           // 禁用某版本
//	retract v1.5.0                          // 撤回自己发布的版本
//
// 版本语义：v主.次.补丁；v2+ 路径需加 /v2；伪版本 v0.0.0-20240101120000-abcdef。
// MVS（最小版本选择）：取恰好满足所有 require 的最小版本，而非最新。
// go.sum：哈希校验，防供应链篡改。
// go.work：多 module 本地联调（go work init / go work use ./...）。
// 常用命令：go mod tidy / graph / download / why / vendor；go list -m all。
package modules

import (
	"fmt"
	"runtime/debug"
)

func Demo() {
	// ReadBuildInfo：编译时嵌入的模块信息（Go 1.18+）
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("no build info")
		return
	}
	fmt.Println("main module:", info.Main.Path, info.Main.Version)
	fmt.Println("go version:", info.GoVersion)
	fmt.Println("build settings:")
	for _, s := range info.Settings {
		fmt.Printf("  %s=%s\n", s.Key, s.Value)
	}
	fmt.Println("dependencies:")
	if len(info.Deps) == 0 {
		fmt.Println("  (none — 仅依赖标准库)")
	}
	for _, dep := range info.Deps {
		fmt.Printf("  %s %s", dep.Path, dep.Version)
		if dep.Replace != nil {
			fmt.Printf(" => %s %s", dep.Replace.Path, dep.Replace.Version)
		}
		fmt.Println()
	}
}
