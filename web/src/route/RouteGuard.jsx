import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router';
import { App } from 'antd';
import { RouteRules } from '@/route/RouteRules.jsx';
import { GetToken } from '@/handler/Token.jsx';
import { AxiosGet } from '@/handler/HandlerUtilsRequest.jsx';
import { Apis } from '@/common/APIConfig.jsx';

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 从路由表中匹配路由，用于验证路由是否合法
/////////////////////////////////////////////////////////////////////////////////////////////////////////////
const RouteMatch = (path, routes) => {
  for (const route of routes) {
    if (route.path === path) {
      return route;
    }
    if (route.children) {
      const childRoute = RouteMatch(path, route.children);
      if (childRoute) {
        return childRoute;
      }
    }
  }
  // 如果没有匹配的路由，返回 null
  return null;
};

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 路由守卫，认证拦截
/////////////////////////////////////////////////////////////////////////////////////////////////////////////
const RouteGuard = ({ children }) => {
  // 获取 Ant Design 的 message 组件
  const { message } = App.useApp();
  // 获取当前路由
  const location = useLocation();
  // 初始化路由跳转
  const navigator = useNavigate();

  useEffect(() => {
    // 从路由表中匹配路由
    const route = RouteMatch(location.pathname, RouteRules);
    // 没有匹配到则返回 404 页面
    if (!route) {
      navigator('/error/404'); // 验证路由是否存在
    } else {
      if (matchRoute.auth && !GetToken()) {
        // 需要登录但是未登录
        message.error('Token 验证失败，请重新登录');
        navigator('/login');
      } else if (matchRoute.auth) {
        // 需要登录且已登录，则校验 Token 是否过期
        AxiosGet(Apis.Public.TokenVerification).then((res) => {
          if (res.code !== 200) {
            message.error('Token 过期，请重新登录');
            navigator('/login');
          }
        });
      } else if (location.pathname === '/login' && GetToken()) {
        // 登录页面还有 Token，则清理 localStorage
        localStorage.clear();
      }
      // 验证路由是否在用户权限范围内，不再则返回 403 页面
    }
  }, [location.pathname]);

  return children;
};

export default RouteGuard;
