package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 定义好API分类的ID，方便后续改动
var (
	systemJobPositionApiCategoryId uint = 1
	systemDepartmentApiCategoryId  uint = 2
	systemUserApiCategoryId        uint = 3
	systemRoleApiCategoryId        uint = 4
	systemMenuApiCategoryId        uint = 5
	systemPermissionApiCategoryId  uint = 6
)

// 系统API分类数据
var systemApiCategories = []model.SystemApiCategory{
	{
		BaseModel: model.BaseModel{Id: systemJobPositionApiCategoryId},
		Name:      "职位接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{Id: systemDepartmentApiCategoryId},
		Name:      "部门接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{Id: systemUserApiCategoryId},
		Name:      "用户接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{Id: systemRoleApiCategoryId},
		Name:      "角色接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{Id: systemMenuApiCategoryId},
		Name:      "菜单接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
	{
		BaseModel: model.BaseModel{Id: systemPermissionApiCategoryId},
		Name:      "权限接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
}

// 迁移系统API分类数据
func MigrateSystemApiCategoryData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统API分类数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api_category")
	common.DB.Create(&systemApiCategories)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统API分类数据迁移完成")
}
