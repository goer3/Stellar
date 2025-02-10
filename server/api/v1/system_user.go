package v1

import (
	"fmt"
	"stellar/common"
	"stellar/dto"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetSystemUserListHandler(ctx *gin.Context) {
	// 获取用户的角色关键字
	systemRoleKeyword, err := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色信息失败")
		return
	}

	// 获取筛选条件
	filter := dto.SystemUserFilterRequest{}
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.FailedWithMessage("获取筛选条件失败")
		return
	}

	// 查询模板和条件
	query := common.DB.Model(&model.SystemUser{})

	// 工号
	if filter.JobNumber != nil {
		query = query.Where("job_number LIKE ?", fmt.Sprintf("%%%s%%", *filter.JobNumber))
	}

	// 用户名
	if filter.Username != nil {
		query = query.Where("username LIKE ?", fmt.Sprintf("%%%s%%", *filter.Username))
	}

	// 中文名
	if filter.CnName != nil {
		query = query.Where("cn_name LIKE ?", fmt.Sprintf("%%%s%%", *filter.CnName))
	}

	// 英文名
	if filter.EnName != nil {
		query = query.Where("en_name LIKE ?", fmt.Sprintf("%%%s%%", *filter.EnName))
	}

	// 邮箱
	if filter.Email != nil {
		query = query.Where("email LIKE ?", fmt.Sprintf("%%%s%%", *filter.Email))
	}

	// 手机号
	if filter.Phone != nil {
		query = query.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", *filter.Phone))
	}

	// 状态
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	// 性别
	if filter.Gender != nil {
		query = query.Where("gender = ?", *filter.Gender)
	}

	// 系统角色
	if filter.SystemRole != nil {
		query = query.Where("system_role_id = ?", *filter.SystemRole)
	}

	// 系统部门
	if filter.SystemDepartment != nil {
		query = query.Joins("INNER JOIN system_user_department_relation ON system_user.id = system_user_department_relation.system_user_id").
			Where("system_user_department_relation.system_department_id = ?", *filter.SystemDepartment)
	}

	// 系统岗位
	if filter.SystemJobPosition != nil {
		query = query.Joins("INNER JOIN system_user_job_position_relation ON system_user.id = system_user_job_position_relation.system_user_id").
			Where("system_user_job_position_relation.system_job_position_id = ?", *filter.SystemJobPosition)
	}

	// 分页信息
	resp := dto.PaginationResponse{}

	// 判断是否需要分页
	if filter.Paginable != nil && *filter.Paginable {
		var pc dto.PaginationConfig
		// 修改分页信息
		pc.Paginable = true
		// 获取数据总数
		query.Count(&pc.Total)
		// 获取分页请求参数
		if filter.PageNumber != nil {
			pc.PageNumber = *filter.PageNumber
		}
		if filter.PageSize != nil {
			pc.PageSize = *filter.PageSize
		}
		// 获取 limit 和 offset
		limit, offset := pc.GetPaginationLimitAndOffset()
		query = query.Offset(offset).Limit(limit)
		// 回填响应数据
		resp.Pagination = pc
	}

	// 查询出最终结果
	var systemUsers []model.SystemUser
	if err := query.Preload("SystemRole").Preload("SystemDepartments").Preload("SystemJobPositions").Find(&systemUsers).Error; err != nil {
		response.FailedWithMessage("获取用户列表失败")
		return
	}

	// 判断用户角色，如果不是超级管理员，则需要隐藏掉手机号
	if systemRoleKeyword != common.SystemSuperAdministratorRoleKeyword {
		for i := range systemUsers {
			if *systemUsers[i].HidePhone == common.BOOLEAN_TRUE {
				systemUsers[i].Phone = utils.HidePhoneNumber(systemUsers[i].Phone)
			}
		}
	}

	// 回填数据
	resp.List = systemUsers

	// 返回结果
	response.SuccessWithData(resp)
}

// 获取指定用户详情
func GetSystemUserDetailHandler(ctx *gin.Context) {

}

// 创建用户
func CreateSystemUserHandler(ctx *gin.Context) {

}

// 批量导入用户
func BatchImportSystemUserHandler(ctx *gin.Context) {

}

// 更新用户详情
func UpdateSystemUserDetailHandler(ctx *gin.Context) {

}

// 批量修改用户
func BatchModifySystemUserHandler(ctx *gin.Context) {
	// 用户状态：operationType(Status)
	// 用户角色：operationType(SystemRole)
	// 用户部门：operationType(SystemDepartment)
	// 用户岗位：operationType(SystemJobPosition)
}
