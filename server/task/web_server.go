package task

import (
	"context"
	"stellar/common"
	"time"
)

// 注册 WebServer 角色
func RegisterWebServer() {
	ctx := context.Background()
	key := common.RKP_WEB_SERVER + common.RK_SEPARATOR + *common.SystemNodeId
	for {
		common.SystemLog.Debug("开始注册WebServer角色，节点ID：" + *common.SystemNodeId)
		if _, err := common.Cache.Set(ctx, key, 1, common.RKE_WEB_SERVER).Result(); err != nil {
			common.SystemLog.Error("WebServer角色注册失败：" + err.Error())
		}
		time.Sleep(common.RKE_INTERVAL_WEB_SERVER)
	}
}
