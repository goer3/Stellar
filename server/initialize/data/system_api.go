package data

import (
	"fmt"
	"stellar/common"
	"time"
)

// 迁移系统Api数据
func MigrateSystemApiData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统Api数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api")
	common.DB.Create(&publicApis)            // 公共接口初始化数据
	common.DB.Create(&systemJobPositionApis) // 职位接口初始化数据
	common.DB.Create(&systemDepartmentApis)  // 部门接口初始化数据
	common.DB.Create(&systemUserApis)        // 用户接口初始化数据
	common.DB.Create(&systemRoleApis)        // 角色接口初始化数据
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统Api数据迁移完成")
}
