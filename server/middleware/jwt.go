package middleware

import (
	"context"
	"errors"
	"fmt"
	"stellar/common"
	"stellar/dto"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           common.Config.Auth.JWT.Realm,                                // 定义一个域，用于展示给用户。一般作为 HTTP 认证时的提示信息
		Key:             []byte(common.Config.Auth.JWT.Key),                          // 签名 JWT 的密钥，用于加密和解密令牌的字符串
		Timeout:         time.Duration(common.Config.Auth.JWT.Timeout) * time.Second, // Token 有效期
		Authenticator:   authenticator,                                               // 用户登录校验
		PayloadFunc:     payloadFunc,                                                 // Token 封装
		LoginResponse:   loginResponse,                                               // 登录成功响应
		Unauthorized:    unauthorized,                                                // 登录，认证失败响应
		IdentityHandler: identityHandler,                                             // 解析 Token
		Authorizator:    authorizator,                                                // 验证 Token
		LogoutResponse:  logoutResponse,                                              // 注销登录
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",          // Token 查找的字段
		TokenHeadName:   "Bearer",                                                    // Token 请求头名称
	})
}

// 隶属 Login 中间件
// 当调用 LoginHandler 就会触发：通过从 ctx 中检索出数据，进行用户登录认证
// 返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
	// 1. 从请求体中获取登录请求
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, errors.New("登录信息数据解析失败")
	}

	// 2. 获取登录用户的客户端 IP 地址
	ip := ctx.ClientIP()
	if ip == "" {
		ip = "none"
	}

	// 3. 判断该用户是否已经达到登录次数限制
	// 登录错误次数 Key
	loginErrorTimesKey := fmt.Sprintf("%s%s%s%s%s", common.RKP_LOGIN_ERROR_TIMES, common.RK_SEPARATOR, req.Account, common.RK_SEPARATOR, ip)

	// 获取登录错误次数
	loginErrorTimes, err := common.Cache.Get(context.Background(), loginErrorTimesKey).Int()
	if err != nil {
		return nil, errors.New("用户登录错误次数获取失败")
	}

	// 判断登录错误次数是否达到上限
	if loginErrorTimes >= common.Config.Auth.Login.ErrorTimesLimit {
		return nil, errors.New("用户登录错误次数已经达到上限")
	}

	// 4. 优化查询逻辑
	var systemUser model.SystemUser
	query := common.DB.Model(&model.SystemUser{})

	// 4.1 根据不同的登录方式，查询不同的用户
	switch {
	// 手机号登录
	case utils.IsPhone(req.Account):
		query = query.Where("phone = ?", req.Account)
	// 邮箱登录
	case utils.IsEmail(req.Account):
		query = query.Where("email = ?", req.Account)
	// 用户名登录
	case utils.IsUsername(req.Account):
		query = query.Where("username = ?", req.Account)
	// 工号登录
	default:
		query = query.Where("job_number = ?", req.Account)
	}

	// 4.2 查询用户
	if err := query.First(&systemUser).Error; err != nil {
		common.Cache.Set(context.Background(), loginErrorTimesKey, loginErrorTimes+1,
			time.Duration(common.Config.Auth.Login.ErrorLockTime)*time.Second)
		return nil, errors.New("用户名或密码错误")
	}

	// 5. 校验用户
	// 5.1 校验密码
	if !utils.ComparePassword(systemUser.Password, req.Password) {
		common.Cache.Set(context.Background(), loginErrorTimesKey, loginErrorTimes+1,
			time.Duration(common.Config.Auth.Login.ErrorLockTime)*time.Second)
		if loginErrorTimes+1 >= common.Config.Auth.Login.ErrorTimesLimit {
			return nil, errors.New("用户名或密码错误，且用户登录错误次数已经达到上限")
		}
		return nil, errors.New("用户名或密码错误")
	}

	// 5.2 验证用户状态
	if *systemUser.Status == common.STATUS_NOTOK {
		return nil, errors.New("用户已禁用，请联系管理员")
	}

	// 6. 使用事务处理登录成功后的操作
	tx := common.DB.Begin()
	// 设置 panic 恢复机制，确保发生异常时事务会回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 6.1 更新最后登录信息
	if err := tx.Model(&systemUser).Updates(map[string]interface{}{
		"last_login_ip": ip,
		"last_login_at": carbon.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("用户最后登录信息更新失败")
	}

	// 6.2 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("用户最后登录信息提交失败")
	}

	// 7. 清理登录错误次数
	common.Cache.Del(context.Background(), loginErrorTimesKey)

	// 8. 设置用户名并返回，方便后续直接解析出用户名然后根据用户名查询 Redis 中保存的 Token
	ctx.Set("username", systemUser.Username)
	return &systemUser, nil
}

// 隶属 Login 中间件
// 接收 Authenticator 验证成功后传递过来的数据，然后将想要的字段 MapClaims 封装到 Token 中
// 后续可以通过 ExtractClaims 对 Token 进行解析
func payloadFunc(data interface{}) jwt.MapClaims {
	// 断言判断获取传递过来数据是不是用户数据
	if systemUser, ok := data.(*model.SystemUser); ok {
		// 封装一些常用的字段，方便直接使用前端和后端都能直接使用
		return jwt.MapClaims{
			// MapClaims 必须包含 IdentityKey
			jwt.IdentityKey:     systemUser.Id,
			"id":                systemUser.Id,
			"jobNumber":         systemUser.JobNumber,
			"username":          systemUser.Username,
			"cnName":            systemUser.CNName,
			"enName":            systemUser.ENName,
			"phone":             systemUser.Phone,
			"email":             systemUser.Email,
			"gender":            systemUser.Gender,
			"avatar":            systemUser.Avatar,
			"systemRoleId":      systemUser.SystemRoleId,
			"systemRoleKeyword": systemUser.SystemRole.Keyword,
		}
	}
	return jwt.MapClaims{}
}

// 隶属 Login 中间件
// 响应用户请求：接收 PayloadFunc 传递过来的 Token 信息，返回登录成功
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	// 登录响应数据
	var resp dto.LoginResponse
	resp.Token = token
	resp.ExpireTime = expire.Format(common.TIME_FORMAT_SECOND)

	// 如果不允许多设备登录，则需要存放到 Redis 中，后续登录状态校验也验证是否和 Redis 中保存的 Token 一致
	if !common.Config.Auth.Login.MultiDevices {
		// 获取登录用户的用户名
		username := ctx.GetString("username")
		if username == "" {
			response.FailedWithMessage("用户登录信息获取失败")
			return
		}
		// 将 Token 保存到 Redis 中，替换掉旧的 Token
		key := common.RKP_LOGIN_TOKEN + common.RK_SEPARATOR + username
		common.Cache.Set(context.Background(), key, token, time.Duration(common.Config.Auth.JWT.Timeout)*time.Second)
	}
	response.SuccessWithData(resp)
}

// 所有登录失败，验证失败的响应
func unauthorized(ctx *gin.Context, code int, message string) {
	response.FailedWithCodeAndMessage(response.RequestUnauthorized, message)
}

// 隶属用户登录后的中间件
// 解析 Token
func identityHandler(ctx *gin.Context) interface{} {
	// 从 Context 中获取用户名
	claims := jwt.ExtractClaims(ctx)
	username, _ := claims["username"].(string)
	return &model.SystemUser{
		Username: username,
	}
}

// 隶属登录后的中间件
// 验证 Token
func authorizator(data interface{}, ctx *gin.Context) bool {
	// 断言判断获取传递过来数据是不是用户数据且用户名不为空
	systemUser, ok := data.(*model.SystemUser)
	if !ok || systemUser.Username == "" {
		return false
	}

	// 当不允许多设备登录配置，需要验证 Redis 中的数据是否一致
	if !common.Config.Auth.Login.MultiDevices {
		// 获取请求中的 Token
		token := jwt.GetToken(ctx)
		// 判断 Redis 中的 Token 和请求中的 Token 是否一致
		key := common.RKP_LOGIN_TOKEN + common.RK_SEPARATOR + systemUser.Username
		if common.Cache.Get(context.Background(), key).Val() != token {
			return false
		}
	}
	return true
}

// 注销登录
func logoutResponse(ctx *gin.Context, code int) {
	// 解析请求中的 Token
	claims := jwt.ExtractClaims(ctx)
	username, ok := claims["username"].(string)
	if !ok || username == "" {
		response.FailedWithMessage("清理登录用户的Token失败")
		return
	}
	// 删除 Redis 中的 Token
	key := common.RKP_LOGIN_TOKEN + common.RK_SEPARATOR + username
	common.Cache.Del(context.Background(), key)
	response.Success()
}
