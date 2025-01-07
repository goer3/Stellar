package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stellar/common"
	"stellar/initialize"
	"stellar/pkg/trans"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&common.SystemListenHost, "host", "", common.SystemListenHost, "指定服务运行监听地址")
	startCmd.Flags().StringVarP(&common.SystemListenPort, "port", "", common.SystemListenPort, "指定服务运行监听端口")
	startCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SYSTEM_DEFAULT_CONFIG_FILENAME, "指定服务运行配置文件")
	startCmd.Flags().StringVarP(&common.SystemRoleWebServer, "web-server", "", common.SystemRoleWebServer, "指定服务是否启动 Web 后端服务，1 表示是，0 表示否")
	startCmd.Flags().StringVarP(&common.SystemRoleLeaderElection, "leader-election", "", common.SystemRoleLeaderElection, "指定服务是否参与 Leader 选举，1 表示是，0 表示否")
	startCmd.Flags().StringVarP(&common.SystemRoleWorker, "worker", "", common.SystemRoleWorker, "指定服务是否运行 Worker 角色，1 表示是，0 表示否")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 Stellar 后端服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(common.LOGO)                                      // 打印系统 Logo
		systemStartTime := fmt.Sprintf("%d", time.Now().UnixNano()) // 系统启动时间
		common.SystemNodeStartTime = trans.String(systemStartTime)  // 保存系统启动时间
		initialize.Config()                                         // 配置文件初始化
		initialize.NodeId()                                         // 生成节点唯一标识ID

		// 打印节点配置启动信息
		tb := table.NewWriter()
		header := table.Row{"配置项名称", "配置项值"}
		tb.AppendHeader(header)
		rows := []table.Row{}
		rows = append(rows, table.Row{"项目名称", common.SYSTEM_NAME})
		rows = append(rows, table.Row{"项目描述", common.SYSTEM_DESCRIPTION})
		rows = append(rows, table.Row{"节点唯一标识ID", *common.SystemNodeId})
		rows = append(rows, table.Row{"系统 Go 版本", common.SYSTEM_GO_VERSION})
		rows = append(rows, table.Row{"系统版本", common.SystemVersion})
		rows = append(rows, table.Row{"开发者名称", common.SYSTEM_DEVELOPER_NAME})
		rows = append(rows, table.Row{"开发者邮箱", common.SYSTEM_DEVELOPER_EMAIL})
		rows = append(rows, table.Row{"配置文件", common.SystemConfigFilename})
		rows = append(rows, table.Row{"监听地址", common.Config.System.Listen.Host})
		rows = append(rows, table.Row{"监听端口", common.Config.System.Listen.Port})
		rows = append(rows, table.Row{"节点角色 Web Server", common.Config.System.Role.WebServer})
		rows = append(rows, table.Row{"节点角色 Leader Election", common.Config.System.Role.LeaderElection})
		rows = append(rows, table.Row{"节点角色 Worker", common.Config.System.Role.Worker})
		tb.AppendRows(rows)
		fmt.Println(tb.Render())

		// 如果启动 Web 后端服务，则需要初始化路由
		if common.Config.System.Role.WebServer {
			// 初始化路由
			r := initialize.Router()

			// 创建 HTTP 服务器
			server := &http.Server{
				Addr:    fmt.Sprintf("%s:%s", common.Config.System.Listen.Host, common.Config.System.Listen.Port),
				Handler: r,
			}

			// 启动 Web 后端服务
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic("Web 后端服务启动失败：" + err.Error())
				}
			}()

			// 监听信号，用于关闭 Web 后端服务
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit

			// 接收到关闭信号后，优雅关闭 Web 后端服务
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				panic("Web 后端服务关闭失败：" + err.Error())
			}
			fmt.Println("Web 后端服务关闭成功")
		} else {
			select {} // 没有配置 Web 后端服务的时候需要一个保活进程
		}
	},
}
