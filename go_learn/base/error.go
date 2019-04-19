package main

import (
	"fmt"
)

func main() {
	// 正常情况
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		fmt.Println("100 / 10 = ", result)
	}
	// 当除数为0时会返回错误信息
	if _, errorMsg := divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

type DivideError struct {
	dividee int
	divider int
}

/**
实现 error 接口
*/
func (de *DivideError) Error() string {
	strFormat := `
		Cannot proceed , the divider is zero.
		dividee : %d
		divider : 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

/*
定义 int 类型触发运算的函数
*/
func divide(varDicidee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDicidee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDicidee / varDivider, ""
	}
}
