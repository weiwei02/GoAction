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
1. 获取数据源的feeds列表。这些数据源从互联网上抓取数据，之后对数据使用特定的搜索项进行匹配
*/
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	/*
		函数返回两个值，第一组返回值是一组Feed类型的切片,
		第二个返回值是一个错误值。
		针对这种具有两个返回值的函数（一个正常值，一个错误值），如果程序发生了错误，永远不应该使用另一个返回值。
		否则程序可能产生更多的错误，甚至崩溃
	*/
	feeds, err := RetrieveFeeds()
	if err != nil {
		// 打印致命错误，并推出程序
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	/**
	何时使用 := 符号声明变量？
	1. 如果需要声明初始值为0的变量，应该使用var关键字声明变量。如果提供确切的非零值初始化变量，
	或者使用函数返回值创建变量，应该使用简化变量声明运算符。
	*/
	results := make(chan *Result)

	/*
			构建一个 waitGroup 以便处理所有的数据源，waitGroup是一个计数信号量，我们可以用它来统计所有的 goroutine 是不是都完成了工作。
		注意 go 声明类型时，默认值并非 nil ，声明变量时已经为变量分配了内存地址
	*/
	var waitGroup sync.WaitGroup
	// 将waitGroup的值设置为要启动的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找.
		// map 可以使用精确查找： 将查找结果赋值给两个变量。第二个变量时一个布尔值
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// 每个 goroutine 完成工作后，都会递减变量的计数值，当这个计数值递减到0时，就代表所有的工作都做完了
			waitGroup.Done()
			// 匿名函数的参数是在这里传递的。没有传递searchTerm和results也能访问到的原因是因为闭包，匿名函数并没有拿到这些变量的副本
			// 而是直接访问外层函数作用域中声明的这些变量本身。因为matcher和feed变量每次调用时值不相同，所以没有使用闭包的方式访问这
			// 两个变量。如果使用闭包访问这两个变量，随着外层函数里变量值得改变，内层的匿名函数也会感知到这些改变。所有的goroutine都会
			// 因为闭包共享同样的变量
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
