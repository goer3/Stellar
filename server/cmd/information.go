package cmd

import (
	"fmt"
	"stellar/common"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(informationCmd)
}

// 系统信息
var informationCmd = &cobra.Command{
	Use:   "info",
	Short: "打印系统信息",
	Run: func(cmd *cobra.Command, args []string) {
		tb := table.NewWriter()
		header := table.Row{"配置项名称", "配置项值"}
		tb.AppendHeader(header)
		rows := []table.Row{}
		rows = append(rows, table.Row{"项目名称", common.SYSTEM_NAME})
		rows = append(rows, table.Row{"项目描述", common.SYSTEM_DESCRIPTION})
		rows = append(rows, table.Row{"系统 Go 版本", common.SYSTEM_GO_VERSION})
		rows = append(rows, table.Row{"系统版本", common.SystemVersion})
		rows = append(rows, table.Row{"开发者信息", common.SYSTEM_DEVELOPER_NAME + " <" + common.SYSTEM_DEVELOPER_EMAIL + ">"})
		tb.AppendRows(rows)
		fmt.Println(tb.Render())
	},
}
