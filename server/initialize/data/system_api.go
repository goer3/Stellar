package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 系统API数据
var systemApis = []model.SystemApi{}

// 迁移系统API数据
func MigrateSystemApiData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统API数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api")
	common.DB.Create(&systemApis)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统API数据迁移完成")
}
