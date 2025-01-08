package task

import (
	"context"
	"stellar/common"
	"time"
)

// 心跳上报
func HeartbeatReport() {
	ctx := context.Background()
	key := common.RKP_HEARTBEAT + common.RK_SEPARATOR + *common.SystemNodeId
	for {
		common.SystemLog.Debug("开始上报节点心跳数据，节点ID：" + *common.SystemNodeId)
		if _, err := common.Cache.Set(ctx, key, *common.SystemNodeStartTime, common.RKE_HEARTBEAT).Result(); err != nil {
			common.SystemLog.Error("节点心跳数据上报失败：" + err.Error())
		}
		time.Sleep(common.RKE_INTERVAL_HEARTBEAT)
	}
}
