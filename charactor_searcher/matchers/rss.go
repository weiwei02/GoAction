package matchers

import (
	"charactor_searcher/search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// 将匹配器注册
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

type (
	// item 根据 item 字段的标签，将定义字段与rss文档关联起来
	item struct {
		XMLNAME     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`

		Link        string `xml:"link"`
		GUID        string `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}

	// image 根据image字段的标签，将定义的字段与rss文档的字段关联起来
	image struct {
		XMLNAME xml.Name `xml:"image"`
		Title   string   `xml:"title"`
		URL     string   `xml:"url"`
		Link    string   `xml:"link"`
	}

	// channel 根据channel字段的标签，将定义的字段与rss文档的字段关联起来
	channel struct {
		XMLNAME        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument 定义了与rss文档关联的字段
	rssDocument struct {
		XMLNAME xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher 实现了Matcher接口
type rssMatcher struct{}

// 在文档中搜索目标项
func (rss rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s] \n", feed.Type, feed.Name, feed.URI)

	// 获取要搜索的数据
	document, err := rss.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 检查标题部分是否包含搜索项
		mached, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 如果找到匹配项，将其作为结果保存
		if mached {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 检查描述部分是否包含被搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}

func (rss rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provoided")
	}
	// 从网络获取数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 一旦从函数返回，关闭返回的相应链接
	defer resp.Body.Close()

	// 检查状态码是不是200，这样就能知道是不是收到了正确的相应
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http Response Error %d\n", resp.StatusCode)
	}

	// 将rss数据源文档解码到我们定义的结构体类型里
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
