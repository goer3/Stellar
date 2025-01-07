package cmd

import (
	"stellar/common"
	"strings"

	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   strings.ToLower(common.SYSTEM_NAME),
	Short: common.SYSTEM_DESCRIPTION,
}

// 执行根命令
func Execute() error {
	return rootCmd.Execute()
}
