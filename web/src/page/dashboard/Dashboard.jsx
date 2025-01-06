import React from 'react';
import { TitleSuffix } from '@/common/Text.jsx';
import { Helmet } from 'react-helmet';

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
  pageTitle: '工作空间' + TitleSuffix,
  // 页面顶部标题
  pageHeaderTitle: '工作空间 / WORKSPACE.',
  // 页面关键词
  pageKeyword: '工作空间',
  // 默认搜索展开字段数量
  defaultFilterExpandItemCount: 8
};

// 页面说明组件
const PageDescriptionComponent = () => {
  return (
    <>
      <div>接口特别说明：</div>
      <ul>
        <li>接口说明1。</li>
        <li>接口说明2。</li>
      </ul>
    </>
  );
};

const Dashboard = () => {
  return (
    <>
      {/* 页面 header */}
      <Helmet>
        <title>{PageConfig.pageTitle}</title>
        <meta name="description" content={PageConfig.pageDesc} />
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

export default Dashboard;
