package data

import (
	"fmt"
	"stellar/common"
	"time"
)

// 迁移系统API数据
func MigrateSystemApiData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统API数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api")
	common.DB.Create(&systemUserApis) // 用户接口初始化数据
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统API数据迁移完成")
}
