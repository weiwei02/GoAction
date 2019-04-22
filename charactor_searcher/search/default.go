package search

// defaultMatcher 实现了默认的匹配器
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	// 空结构体在创建实例时，不会分配任何内存
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
