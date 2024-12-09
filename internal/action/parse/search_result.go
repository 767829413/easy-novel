package parse

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/767829413/easy-novel/internal/action/model"
	"github.com/767829413/easy-novel/internal/action/source"
	"github.com/767829413/easy-novel/pkg/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

const timeoutMillis = 15000

type SearchResultParser struct {
	sourceID int
	rule     *model.Rule
}

func NewSearchResultParser(sourceID int) *SearchResultParser {
	return &SearchResultParser{
		sourceID: sourceID,
		rule:     source.GetRuleBySourceID(sourceID),
	}
}

func (p *SearchResultParser) Parse(keyword string) ([]*model.SearchResult, error) {
	search := p.rule.Search
	isPaging := search.Pagination

	// Simulate search request
	doc, err := fetchDocument(
		search.URL,
		search.Method,
		p.rule.Search.Body,
		keyword,
		p.rule.Search.Cookies,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching document: %w", err)
	}

	firstPageResults, err := p.getSearchResults("", doc)
	if err != nil {
		return nil, err
	}

	if !isPaging {
		return firstPageResults, nil
	}

	urls := make(map[string]struct{})
	doc.Find(search.NextPage).Each(func(_ int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			urls[utils.NormalizeURL(href, p.rule.URL)] = struct{}{}
		}
	})

	var wg sync.WaitGroup
	resultChan := make(chan []*model.SearchResult, len(urls))
	errorChan := make(chan error, len(urls))
	semaphore := make(chan struct{}, 20) // Limit concurrency to 20

	for url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			results, err := p.getSearchResults(url, nil)
			if err != nil {
				errorChan <- err
				return
			}
			resultChan <- results
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	var additionalResults []*model.SearchResult
	for results := range resultChan {
		additionalResults = append(additionalResults, results...)
	}

	// Check for errors
	for err := range errorChan {
		return nil, err
	}

	return append(firstPageResults, additionalResults...), nil
}

func (p *SearchResultParser) getSearchResults(
	url string,
	doc *goquery.Document,
) ([]*model.SearchResult, error) {
	var err error
	if doc == nil {
		doc, err = fetchDocument(url, "GET", "", "", "")
		if err != nil {
			return nil, err
		}
	}

	var results []*model.SearchResult
	doc.Find(p.rule.Search.Result).Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Find(p.rule.Search.BookName).Attr("href")
		bookName := s.Find(p.rule.Search.BookName).Text()
		latestChapter := s.Find(p.rule.Search.LatestChapter).Text()
		author := s.Find(p.rule.Search.Author).Text()
		update := s.Find(p.rule.Search.Update).Text()

		if bookName == "" {
			return
		}

		result := &model.SearchResult{
			Url:           utils.NormalizeURL(href, p.rule.URL),
			BookName:      bookName,
			LatestChapter: latestChapter,
			Author:        author,
			LatestUpdate:  update,
		}

		results = append(results, result)
	})

	return results, nil
}

func fetchDocument(urlStr, method, body, keyword, cookies string) (*goquery.Document, error) {
	client := resty.New().
		SetTimeout(time.Duration(timeoutMillis)*time.Millisecond).
		SetHeader("User-Agent", utils.GenerateRandomUA())

	request := client.R()

	// 获取标准的请求方式
	realMethod := utils.BuildMethod(method)

	// 添加 Cookie 参数
	if len(cookies) > 0 {
		cookieParams, err := utils.BuildCookies(cookies)
		if err != nil {
			return nil, err
		}
		for k, v := range cookieParams {
			request.SetCookie(&http.Cookie{Name: k, Value: v})
		}
	}

	// 构建查询参数
	if len(body) > 0 {
		bodyParams, err := utils.BuildParams(body, keyword)
		if err != nil {
			return nil, err
		}
		if realMethod == "GET" {
			request.SetQueryParams(bodyParams)
		} else {
			request.SetBody(bodyParams)
		}
	}

	resp, err := request.Execute(realMethod, urlStr)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	return doc, nil
}
