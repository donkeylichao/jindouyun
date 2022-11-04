package jdyError

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrCompanyNotExists = errors.New("不支持当前公司配置")
	ErrNotSupportHandle = errors.New("不支持的操作类型")
)

func CheckError(err error, exit bool) bool {
	if err == nil {
		return false
	}
	fmt.Printf("%s\n", err)
	if exit {
		os.Exit(0)
	}
	return true
}