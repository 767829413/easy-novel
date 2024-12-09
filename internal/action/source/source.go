package source

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/767829413/easy-novel/internal/action/model"
	"github.com/fatih/color"
)

var ruleCache sync.Map

func GetRuleBySourceID(sourceId int) *model.Rule {
	if v, ok := ruleCache.Load(sourceId); ok {
		return v.(*model.Rule)
	}

	// 构建文件路径
	filePath := filepath.Join("rule", fmt.Sprintf("rule%d.json", sourceId))
	// 读取文件内容
	ruleData, err := os.ReadFile(filePath)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red("书源规则初始化失败，请检查配置项 source-id"))
		fmt.Println(red("错误信息："), err.Error())
		os.Exit(1)
	}

	// 解析 JSON 到 Rule 结构体
	var rule model.Rule
	err = json.Unmarshal(ruleData, &rule)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red("内部解析错误："), err.Error())
		os.Exit(1)
	}

	// 使用 LoadOrStore 方法来确保线程安全的写入
	actual, _ := ruleCache.LoadOrStore(sourceId, &rule)
	return actual.(*model.Rule)
}
