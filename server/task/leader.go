package task

import (
	"context"
	"stellar/common"
	"time"
)

// Leader 节点选举
func ElectionLeader() {
	ctx := context.Background()
	key := common.RK_LEADER
	for {
		// 查询当前是否有 Leader 节点
		leader, _ := common.Cache.Get(ctx, key).Result()
		if leader == "" {
			// 如果没有，则参与选举
			common.SystemLog.Debug("没有节点属于Leader角色，开始参与选举领导节点，节点ID：" + *common.SystemNodeId)
			if _, err := common.Cache.SetNX(ctx, key, *common.SystemNodeId, common.RKE_LEADER).Result(); err != nil {
				common.SystemLog.Error("Leader节点选举失败：" + err.Error())
			}
		} else {
			// 如果有，则判断是不是自己，如果是自己，则更新过期时间
			if leader == *common.SystemNodeId {
				common.SystemLog.Debug("当前节点已经是Leader，节点ID：" + *common.SystemNodeId)
				if _, err := common.Cache.Expire(ctx, key, common.RKE_LEADER).Result(); err != nil {
					common.SystemLog.Error("Leader节点过期时间更新失败：" + err.Error())
				}
			}
		}
		time.Sleep(common.RKE_INTERVAL_LEADER)
	}
}
