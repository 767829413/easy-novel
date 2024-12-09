package action

import (
	"fmt"
	"os"

	// "github.com/767829413/easy-novel/internal/action/source"
	"github.com/767829413/easy-novel/internal/config"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

type printConf struct {
	log *logrus.Logger
}

func NewPrintConf(l *logrus.Logger) Action {
	return &printConf{log: l}
}

func (p *printConf) Execute() error {
	fmt.Println(config.GetConf().ToJSON())
	return nil
}

type printHint struct {
	log     *logrus.Logger
	version string
}

func NewPrintHint(l *logrus.Logger, version string) Action {
	return &printHint{log: l}
}

func (p *printHint) Execute() error {
	cfg := config.GetConf()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"使用须知"})
	table.SetBorder(true)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetAutoWrapText(false)

	// 使用color.New()创建彩色文本
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	table.Append([]string{blue(fmt.Sprintf("easy-novel %s （本项目开源且免费）", p.version))})
	table.Append([]string{fmt.Sprintf("官方地址：%s", "https://github.com/767829413/easy-novel")})
	table.Append([]string{cyan(fmt.Sprintf("当前书源：%s (ID: %d)", "未实现", cfg.Base.SourceID))})
	table.Append([]string{cyan(fmt.Sprintf("导出格式：%s", cfg.Base.Extname))})
	table.Append([]string{""}) // 空行
	table.Append([]string{yellow("请务必阅读 readme.txt")})

	table.Render()

	return nil
}
