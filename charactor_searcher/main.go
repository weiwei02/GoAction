package main

import (
	"charactor_searcher/search"
	"log"
	"os"
)

/**
init 函数在main之前调用
*/
func init() {
	// 将日志标准错误流输出到标准输出流
	log.SetOutput(os.Stdout)
}

/**
程序主执行方法
*/
func main() {
	search.Run("president")
}
