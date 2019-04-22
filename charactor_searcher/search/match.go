package search

/**
以一致且通用的方法，来处理不同类型的匹配器值
*/

//Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的新搜索类型的行为
// Go 对接口的命名惯例： 如果接口类型只包含一个方法，那么这个类型的名字以er结尾。
// 如果接口类型内部声明了多个方法，其名字需要与其行为关联
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
