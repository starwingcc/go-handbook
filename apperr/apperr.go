// Package apperr 演示自定义错误、哨兵、包装与合并。
package apperr

import (
	"errors"
	"fmt"
	"strconv"
)

// 哨兵错误
var ErrNotFound = errors.New("not found")

// 带类型的错误
type NotFound struct{ Name string }

func (e *NotFound) Error() string { return fmt.Sprintf("%s not found", e.Name) }

func Find(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	return &NotFound{Name: "item-" + strconv.Itoa(id)}
}

// Wrap 用 %w 包装，使 errors.Is/As 能沿链查找
func Wrap(err error) error { return fmt.Errorf("lookup: %w", err) }

// Join 合并多个错误为一个（Go 1.20+ errors.Join）
func Join(errs ...error) error { return errors.Join(errs...) }

func Demo() {
	// errors.Is：沿包装链匹配身份
	if err := Find(0); err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("err: 命中哨兵")
		}
	}

	// errors.As：从链中提取带类型错误
	err := Find(1)
	var nf *NotFound
	if errors.As(err, &nf) {
		fmt.Println("err: typed", nf.Name)
	}

	// %w 包装；Is/As 可遍历
	wrapped := Wrap(ErrNotFound)
	fmt.Println("wrap Is:", errors.Is(wrapped, ErrNotFound))

	// errors.Join 合并多错误
	joined := Join(ErrNotFound, Find(1))
	fmt.Println("joined Is ErrNotFound:", errors.Is(joined, ErrNotFound))
	fmt.Println("joined Error():", joined.Error())
}
