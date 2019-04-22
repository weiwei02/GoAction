package search

import "log"

// defaultMatcher 实现了默认的匹配器
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	// 空结构体在创建实例时，不会分配任何内存
	var matcher defaultMatcher
	Register("default", matcher)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
// 匹配器的注册器在main函数调用之前调用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register ", feedType, " matcher")
	matchers[feedType] = matcher
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
