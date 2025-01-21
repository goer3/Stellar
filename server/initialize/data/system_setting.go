package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 系统设置数据
var systemSettings = []model.SystemSetting{
	{Key: "2fa_enabled", Value: "1", Description: "是否启用双因子验证", Updater: common.SYSTEM_DEFAULT_CREATOR},
}

// 系统设置数据迁移
func MigrateSystemSettingData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统设置数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_setting")
	common.DB.Create(&systemSettings)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统设置数据迁移完成")
}
