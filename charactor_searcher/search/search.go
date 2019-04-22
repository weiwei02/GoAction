package search

import (
	"log"
	"sync"
)

/* 匹配器 包级变量，变量名以小写字母开头
Go 中以小写字母开头的标识符不能被其它包直接访问
*/
var matchers = make(map[string]Matcher)

/**
Run 执行搜索逻辑
*/
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RerieveFeeds()
	if err != nil {
		// 打印致命错误，并推出程序
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	/*
		构建一个 waitGroup 以便处理所有的数据源
	*/
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务都完成了
		waitGroup.Wait()

		// 用关闭通道的方式通知 Display 函数
		close(results)
	}()

	/*
		启动函数，显示返回结果，并且在最后一个结果显示完成后返回
	*/
}
