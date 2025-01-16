import { proxy } from 'valtio';

// 角色相关状态
export const SystemRoleStates = proxy({
  SystemRoleApiList: [], // 角色Api数据
  SystemRoleMenuList: [] // 角色菜单数据
});
