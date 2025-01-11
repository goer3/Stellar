import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router';
import { App } from 'antd';
import { APIGET } from '@/handler/Request.jsx';
import { GetUnexpireToken } from '@/handler/Token.jsx';
import { BackendAPIPrefix, BackendAPISuffix } from '@/common/Api.jsx';
import { RouteRules } from '@/route/RouteRules.jsx';

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
  // 清理缓存并返回登录页
  const ClearLocalStorageAndNavigateToLoginPath = () => {
    localStorage.clear();
    navigator('/login');
  };

  useEffect(() => {
    // 从路由表中匹配路由
    const route = RouteMatch(location.pathname, RouteRules);
    // 没有匹配到则返回 404 页面
    if (!route) {
      navigator('/error/404');
    } else {
      // 如果路由存在，则需要多方面判断
      // 1. 如果路由需要登录
      if (route.auth) {
        if (!GetUnexpireToken()) {
          // 1.1 如果本地记录的 Token 为空或者已过期，则返回登录页面
          message.error('Token已过期，请重新登录');
          ClearLocalStorageAndNavigateToLoginPath();
        } else {
          // 1.2 如果本地记录的 Token 没过期，则需要校验 Token 在后端缓存中是否过期
          APIGET(BackendAPIPrefix + BackendAPISuffix.PublicAuth.TokenVerification).then((res) => {
            if (res.code !== 200) {
              // 1.2.1 如果 Token 过期，则返回登录页面
              message.error('Token过期，请重新登录');
              ClearLocalStorageAndNavigateToLoginPath();
            } else {
              // 1.2.2 如果 Token 没过期，但是访问 /login 页面，则不允许访问，直接返回 /dashboard 页面
              if (location.pathname === '/login') {
                navigator('/dashboard');
              }
            }
          });
        }
      }
      // TODO: 验证路由是否在用户权限范围内，不再则返回 403 页面
    }
  }, [location.pathname]);

  return children;
};

export default RouteGuard;
