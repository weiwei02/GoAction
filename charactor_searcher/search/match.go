package search

import "log"

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

// match 函数，为每个数据源单独启动 goroutine 来执行这个函数并发执行搜索
// 这里的通道是一个单向通道，只允许发送数据
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

//Display 从每个单独的 goroutine 接受到结果后，在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直阻塞，直到有数据写入
	// 一旦通道被关闭，for循环就会终止
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
