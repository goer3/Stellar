package task

import (
	"context"
	"stellar/common"
	"time"
)

// Worker 节点注册
func RegisterWorker() {
	ctx := context.Background()
	key := common.RKP_WORKER + common.RK_SEPARATOR + *common.SystemNodeId
	for {
		common.SystemLog.Debug("开始注册Worker节点角色，节点ID：" + *common.SystemNodeId)
		if _, err := common.Cache.Set(ctx, key, 1, common.RKE_WORKER).Result(); err != nil {
			common.SystemLog.Error("Worker节点角色注册失败：" + err.Error())
		}
		time.Sleep(common.RKE_INTERVAL_WORKER)
	}
}
