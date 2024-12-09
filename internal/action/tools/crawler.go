package tools

import (
	"fmt"
	"time"

	"github.com/767829413/easy-novel/internal/action/model"
	"github.com/767829413/easy-novel/internal/action/parse"
	"github.com/767829413/easy-novel/internal/config"
)

type Crawler interface {
	Search(key string) []*model.SearchResult
	Crawl(res *model.SearchResult, start, end int) *model.CrawlResult
}

type novelCrawler struct{}

func NewNovelCrawler() Crawler {
	return &novelCrawler{}
}

func (nc *novelCrawler) Search(key string) []*model.SearchResult {
	conf := config.GetConf()
	// Implement search logic using a search engine
	fmt.Println("<== 正在搜索...")
	start := time.Now()
	// 解析
	searchResults, err := parse.NewSearchResultParser(conf.Base.SourceID).Parse(key)
	if err != nil {
		fmt.Printf("<== 执行搜索索失败：%s\n", err.Error())
		return nil
	}
	duration := time.Since(start)
	fmt.Printf("<== 搜索到 %d 条记录，耗时 %f s\n", len(searchResults), duration.Seconds())
	return searchResults
}

func (nc *novelCrawler) Crawl(res *model.SearchResult, start, end int) *model.CrawlResult {
	// Implement chapter crawling logic
	return &model.CrawlResult{TakeTime: 0}
}
