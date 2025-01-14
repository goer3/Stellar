// 后端地址
const BackendURL = window.CONFIG.backendUrl;

// 后端 API 版本前缀
const BackendAPIURIPrefix = '/api/v1';

// 后端 API 前缀
const BackendAPIPrefix = BackendURL + BackendAPIURIPrefix;

// 后端 API 后缀
const BackendAPISuffix = {
  Public: {
    Common: {
      Health: { Path: '/health', Method: 'GET', Description: '健康检查，用于检查服务是否正常' },
      Information: { Path: '/information', Method: 'GET', Description: '获取系统信息，用于获取系统基础信息' },
      Login: { Path: '/login', Method: 'POST', Description: '用户登录，用于用户登录系统认证' }
    },
    Auth: {
      Logout: { Path: '/logout', Method: 'GET', Description: '注销登录，用于用户注销系统登录状态' },
      TokenVerification: { Path: '/token-verification', Method: 'GET', Description: '校验Token是否过期，用于校验Token是否过期' }
    }
  },
  System: {
    Department: {
      Auth: {},
      AuthAndPermission: {}
    },
    JobPosition: {
      Auth: {},
      AuthAndPermission: {}
    },
    User: {
      Auth: {
        APIList: { Path: '/user/api-list', Method: 'GET', Description: '获取用户授权接口列表，用于获取用户授权接口列表' }
      },
      AuthAndPermission: {}
    },
    Role: {
      Auth: {},
      AuthAndPermission: {}
    },
    Menu: {
      Auth: {},
      AuthAndPermission: {}
    },
    API: {
      Auth: {},
      AuthAndPermission: {}
    },
    Permission: {
      Auth: {},
      AuthAndPermission: {}
    }
  }
};

export { BackendURL, BackendAPIURIPrefix, BackendAPIPrefix, BackendAPISuffix };
