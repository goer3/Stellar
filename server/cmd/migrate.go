package cmd

import (
	"stellar/common"
	"stellar/initialize"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateTableCmd)
	migrateTableCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SYSTEM_DEFAULT_CONFIG_FILENAME, "指定服务运行配置文件")
}

// 数据迁移
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据迁移，支持数据表迁移和基础数据迁移",
}

// 数据表迁移
var migrateTableCmd = &cobra.Command{
	Use:   "table",
	Short: "数据表迁移",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateTables()
	},
}
