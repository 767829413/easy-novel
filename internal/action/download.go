package action

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/767829413/easy-novel/internal/action/model"
	"github.com/767829413/easy-novel/internal/action/tools"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

type download struct {
	log *logrus.Logger
}

func NewDownload(l *logrus.Logger) Action {
	return &download{log: l}
}

func (d *download) Execute() error {
	d.log.Info("Starting novel download")
	color.Green("开始下载小说...")
	// 实现下载小说的逻辑
	rl, err := readline.New("")
	if err != nil {
		return err
	}
	defer rl.Close()

	// 1. Query
	blue := color.New(color.FgBlue).SprintFunc()
	prompt := blue("==> 请输入书名或作者（宁少字别错字）: ")
	rl.SetPrompt(prompt)

	keyword, err := rl.Readline()
	if err != nil {
		return err
	}
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return nil
	}

	crawler := tools.NewNovelCrawler()
	results := crawler.Search(keyword)

	if len(results) == 0 {
		return nil
	}

	// 2. Print search results
	d.printSearchResult(results)

	// 3. Select download chapter
	for {
		rl.SetPrompt("==> 请输入下载序号（首列的数字，或输入 0 返回）：")
		input, err := rl.Readline()
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		if num == 0 {
			return nil
		}
		if num < 0 || num > len(results) {
			continue
		}

		sr := results[num-1]
		fmt.Printf("<== 你选择了《%s》(%s)\n", sr.BookName, sr.Author)
		fmt.Println("==> 0: 重新选择功能")
		fmt.Println("==> 1: 下载全本")
		fmt.Println("==> 2: 下载指定章节")
		fmt.Println("==> 3: 重新输入序号")

		rl.SetPrompt("==> 请输入数字：")
		actionInput, err := rl.Readline()
		if err != nil {
			return err
		}

		action, err := strconv.Atoi(strings.TrimSpace(actionInput))
		if err != nil {
			continue
		}

		if action != 3 {
			if action == 0 {
				return nil
			}

			start, end := 1, int(^uint(0)>>1) // Max int
			if action == 2 {
				rl.SetPrompt("==> 请输起始章(最小为1)和结束章，用空格隔开：")
				rangeInput, err := rl.Readline()
				if err != nil {
					return err
				}
				rangeInput = strings.TrimSpace(rangeInput)
				parts := strings.Split(rangeInput, " ")
				if len(parts) != 2 {
					return fmt.Errorf("invalid input")
				}
				start, err = strconv.Atoi(parts[0])
				if err != nil {
					return err
				}
				end, err = strconv.Atoi(parts[1])
				if err != nil {
					return err
				}
			}

			res := crawler.Crawl(sr, start, end)
			fmt.Printf("<== 完成！总耗时 %d s\n", res)
			return nil
		}
	}
}

func (d *download) printSearchResult(results []*model.SearchResult) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "书名", "作者", "最新章节", "最后更新时间"})

	for i, r := range results {
		table.Append([]string{
			strconv.Itoa(i + 1),
			r.BookName,
			r.Author,
			r.LatestChapter,
			r.LatestUpdate,
		})
	}

	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetTablePadding("\t") // tab-separated columns

	table.Render()
}
