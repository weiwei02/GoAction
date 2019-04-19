package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Books
	var book2 Books

	// book1 描述
	book1.title = "Go 语言"
	book1.author = "golang.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 1024214

	// book2 描述
	book2.title = "在线DEBUG技术"
	book2.author = "基本"
	book2.subject = "主题"
	book2.book_id = 21321

	fmt.Printf("book 1 titile : %s \n", book1.title)

	// slice
	var array []int
	array = append(array, 0)
	fmt.Println(array)

	var array2 = make([]int, 1)
	array2 = append(array2, 0)
	array2 = append(array2, 1)
	fmt.Println(array2)
}
