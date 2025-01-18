import React from 'react';
import { useSnapshot } from 'valtio';
import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { App } from 'antd';

// 页面常量设置
const PageConfig = {
  // 默认页码
  defaultPageNumber: 1,
  // 默认每页显示的数据量
  defaultPageSize: 2,
  // 默认数据总数
  defaultTotal: 0,
  // 默认是否需要分页
  defaultIsPagination: true,
  // 页面标题
  pageTitle: '用户管理' + TitleSuffix,
  // 页面描述
  pageDescription: '用户管理',
  // 页面顶部标题
  pageHeaderTitle: '用户管理 / USER MANAGEMENT.',
  // 页面关键词
  pageKeyword: '用户管理',
  // 默认搜索展开字段数量
  defaultFilterExpandItemCount: 8
};

// 页面说明组件
const PageDescriptionComponent = () => {
  return (
    <>
      <div>用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，系统提供了以下的要求和建议，请知悉：</div>
      <ul>
        <li>系统内置的超级管理员账户（super）涉及到系统的基础逻辑判断，不允许从数据库中物理删除，用户在初始化完成之后需要对其基础信息进行修改，以此保障系统安全性。</li>
        <li>为了保障数据的安全性和可靠性，系统选择通过禁用用户来替代删除用户，禁用用户将无法登录系统，但是数据会仍然被保留，可以随时恢复。</li>
        <li>针对某些特殊的用户，例如老板、高管等，系统建议您隐藏其联系方式，做好保护隐私保护。</li>
      </ul>
    </>
  );
};

const SystemUser = () => {
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 全局配置
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 消息提示
  const { message } = App.useApp();
  // 接口列表
  const { SystemRoleApiList, SystemRoleMenuList } = useSnapshot(SystemRoleStates);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 基础数据查询
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  return (
    <>
      {/* 页面 header */}
      <Helmet>
        <title>{PageConfig.pageTitle}</title>
        <meta name="description" content={PageConfig.pageDescription} />
      </Helmet>
      {/* 页面头部介绍 */}
      <div className="admin-page-header">
        <div className="admin-page-title">{PageConfig.pageHeaderTitle}</div>
        <div className="admin-page-desc">
          <PageDescriptionComponent />
        </div>
      </div>
      {/* 页面主体 */}
      <div className="admin-page-main"></div>
    </>
  );
};

export default SystemUser;
