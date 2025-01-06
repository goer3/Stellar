// 后端地址
const BackendURL = window.CONFIG.backendUrl;

// 后端 API 前缀
const BackendAPIPrefix = '/api/v1';

// 后端 API 后缀
const BackendAPISuffix = {
  Public: {
    Login: { Path: '/login', Method: 'GET', Description: '用户登录' },
    Health: { Path: '/health', Method: 'GET', Description: '健康检查' },
    Version: { Path: '/version', Method: 'GET', Description: '版本信息' },
    Information: { Path: '/information', Method: 'GET', Description: '信息查询' }
  },
  PublicAuth: {
    Logout: { Path: '/logout', Method: 'GET', Description: '用户登出' }
  },
  System: {
    UserAuth: {},
    UserAuthAndPermission: {}
  }
};
