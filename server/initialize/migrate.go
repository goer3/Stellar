package initialize

import (
	"fmt"
	"stellar/common"
	"stellar/initialize/data"
	"stellar/model"
	"time"
)

// 数据表结构迁移
func MigrateTables() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行数据表结构同步")
	common.DB.AutoMigrate(&model.SystemSetting{})
	common.DB.AutoMigrate(&model.SystemRole{})
	common.DB.AutoMigrate(&model.SystemUser{})
	common.DB.AutoMigrate(&model.SystemMenu{})
	common.DB.AutoMigrate(&model.SystemJobPosition{})
	common.DB.AutoMigrate(&model.SystemDepartment{})
	common.DB.AutoMigrate(&model.SystemApiCategory{})
	common.DB.AutoMigrate(&model.SystemApi{})
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t数据表结构同步完成")
}

// 数据表数据迁移
func MigrateTablesData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行数据表数据迁移")
	data.MigrateSystemSettingData()
	data.MigrateSystemRoleData()
	data.MigrateSystemUserData()
	data.MigrateSystemJobPositionData()
	data.MigrateSystemDepartmentData()
	data.MigrateSystemMenuData()
	data.MigrateSystemApiCategoryData()
	data.MigrateSystemApiData()
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t数据表数据迁移完成")
}
