package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 定义好Api分类的ID，方便后续改动
var (
	publicApiCategoryId            uint = 1
	systemJobPositionApiCategoryId uint = 2
	systemDepartmentApiCategoryId  uint = 3
	systemUserApiCategoryId        uint = 4
	systemRoleApiCategoryId        uint = 5
	systemMenuApiCategoryId        uint = 6
	systemPermissionApiCategoryId  uint = 7
	systemSettingApiCategoryId     uint = 8
)

// 系统Api分类数据
var systemApiCategories = []model.SystemApiCategory{
	{
		BaseModel: model.BaseModel{Id: publicApiCategoryId},
		Name:      "公共接口",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
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
	{
		BaseModel: model.BaseModel{Id: systemSettingApiCategoryId},
		Name:      "系统配置",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
	},
}

// 迁移系统Api分类数据
func MigrateSystemApiCategoryData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统Api分类数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api_category")
	common.DB.Create(&systemApiCategories)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统Api分类数据迁移完成")
}
