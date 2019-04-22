package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

/**
读取数据源，并将每个JSON文档解码并返回数据源切片
*/
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 当函数返回时，关闭文件
	// defer关键字会安排随后的函数调用在函数返回时才执行，使用defer关键字可以保证Close函数一定会被调用。
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向一个Feed类型的指针
	var feeds []*Feed
	// Decode 方法使用 interface{} 类型作为参数，可以接受任何类型。这个类型在Go中很特殊，一般会配合reflect包里提供的反射功能一起使用
	err = json.NewDecoder(file).Decode(&feeds)
	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}

/**
数据源提供者类型，该类型与data.json中的数据格式相对应
*/
type Feed struct {
	// 字段名后的 `` 部分叫做tag，这个标记里描述了JSON解码的元数据，用于创建Feed类型值的切片。
	// 每个标记将结构类型字段对应到JSON文档里指定名字的字段
	Name string `json:"site"`
	URI  string `json:"uri"`
	Type string `json:"type"`
}
