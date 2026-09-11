package testingdemo

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// 表驱动测试 + 子测试 t.Run + t.Parallel
func TestReverse(t *testing.T) {
	t.Parallel() // 标记可并行执行
	cases := []struct {
		in, want string
	}{
		{"abc", "cba"},
		{"", ""},
		{"a", "a"},
		{"中文", "文中"},
	}
	for _, c := range cases {
		// Go 1.22+ 每轮 range 自动新建变量，无需 c := c 捕获
		t.Run(c.in, func(t *testing.T) {
			t.Parallel()
			if got := reverse(c.in); got != c.want {
				t.Errorf("reverse(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

// t.Helper：失败时堆栈定位到调用处而非辅助函数内部
func TestWithHelper(t *testing.T) {
	checkLen := func(t *testing.T, s string, want int) {
		t.Helper()
		if len(s) != want {
			t.Fatalf("len(%q)=%d, want %d", s, len(s), want)
		}
	}
	checkLen(t, "hello", 5)
}

// t.Skip 条件跳过；t.Cleanup 返回时清理（LIFO，后注册先执行）
func TestSkipAndCleanup(t *testing.T) {
	if strings.TrimSpace(os.Getenv("SKIP_DEMO")) != "" {
		t.Skip("skipped via SKIP_DEMO env")
	}
	t.Cleanup(func() { /* 清理 1，后执行 */ })
	t.Cleanup(func() { /* 清理 2，先执行（LIFO） */ })
}

// Example：// Output: 注释匹配则通过。Example_xxx 为独立示例（不绑定某标识符）
func Example_reverse() {
	fmt.Println(reverse("hello"))
	// Output: olleh
}

// Benchmark：go test -bench=. ./testingdemo
func BenchmarkReverse(b *testing.B) {
	for range b.N {
		reverse("hello")
	}
}

// 被测函数
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
