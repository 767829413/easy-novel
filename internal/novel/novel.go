package novel

import (
	"context"
	"strings"

	"github.com/767829413/easy-novel/internal/definition"
	"github.com/767829413/easy-novel/internal/functions"
	"github.com/767829413/easy-novel/pkg/utils"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Run is the main entry point for the novel downloading logic
func Run(ctx context.Context, log *logrus.Logger) error {
	log.Info("Starting novel download process")
	options := []string{"1.下载小说", "2.检查更新", "3.查看配置文件", "4.使用须知", "5.结束程序"}
	actions := map[string]functions.App{
		definition.NovelCapability_DOWNLOAD:     functions.NewDownload(log),
		definition.NovelCapability_CHECK_UPDATE: functions.NewCheckUpdate(log, 5000),
		definition.NovelCapability_PRINT_CONF:   functions.NewPrintConf(log),
		definition.NovelCapability_PRINT_HINT:   functions.NewPrintHint(log),
		definition.NovelCapability_EXIT:         functions.NewExit(log),
	}

	var completerItems []readline.PrefixCompleterInterface
	for _, option := range options {
		completerItems = append(completerItems, readline.PcItem(option))
	}

	completer := readline.NewPrefixCompleter(completerItems...)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "按 Tab 键选择功能: ",
		AutoComplete: completer,
	})

	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, err := rl.Readline()
			if err != nil {
				return err
			}
			cmd := strings.TrimSpace(line)
			index := strings.Split(cmd, ".")[0]
			action, found := actions[index]
			if !found {
				utils.GetColorIns(color.FgHiRed).Println("无效的选项，请重新选择")
				continue
			}

			err = action.Execute()
			if err != nil {
				if err.Error() == "exit" {
					return nil
				}
				utils.GetColorIns(color.FgHiRed).Printf("执行操作时发生错误: %v\n", err)
			}
		}
	}
}
