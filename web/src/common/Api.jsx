// 后端地址
const BackendURL = window.CONFIG.backendUrl;

// 后端 Api 版本前缀
const BackendApiURIPrefix = '/api/v1';

// 后端 Api 前缀
const BackendApiPrefix = BackendURL + BackendApiURIPrefix;

// 后端 Api 后缀
const BackendApiSuffix = {
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
      AuthAndPermission: {
        List: { Path: '/system/department/list', Method: 'GET', Description: '获取系统部门列表' }
      }
    },
    JobPosition: {
      Auth: {},
      AuthAndPermission: {
        List: { Path: '/system/job-position/list', Method: 'GET', Description: '获取系统职位列表' }
      }
    },
    User: {
      Auth: {},
      AuthAndPermission: {
        List: { Path: '/system/user/list', Method: 'GET', Description: '获取用户列表' }
      }
    },
    Role: {
      Auth: {
        ApiList: { Path: '/system/role/api-list', Method: 'GET', Description: '获取角色授权接口列表' },
        MenuList: { Path: '/system/role/menu-list', Method: 'GET', Description: '获取角色授权菜单列表' }
      },
      AuthAndPermission: {
        List: { Path: '/system/role/list', Method: 'GET', Description: '获取系统角色列表' }
      }
    },
    Menu: {
      Auth: {},
      AuthAndPermission: {}
    },
    Api: {
      Auth: {},
      AuthAndPermission: {}
    },
    Permission: {
      Auth: {},
      AuthAndPermission: {}
    }
  }
};

export { BackendURL, BackendApiURIPrefix, BackendApiPrefix, BackendApiSuffix };
