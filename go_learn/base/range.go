package main

import "fmt"

func main() {
	sliceRange()

	mapRange()

	rangeUnicode()
}

/**
range 遍历枚举unicode字符串，第一个参数是字符的索引，第二个是字符（unicode的值）本身
*/
func rangeUnicode() {
	for index, value := range "go" {
		fmt.Println(index, value)
	}
}

/**
* range 遍历map，第一个参数是key，第二个参数是value
 */
func mapRange() {
	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}
	for key, value := range kvs {
		fmt.Printf("%s -> %s \n", key, value)
	}

	delete(kvs, "delete")
}

/**
range遍历数组，第一个参数是索引，第二个参数是数组索引位置对应的元素
*/
func sliceRange() {
	// 使用range去求slice的和
	println("slice range")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum : ", sum)
}
