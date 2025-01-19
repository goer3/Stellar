import React, { useState } from 'react';
import { useSnapshot } from 'valtio';
import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { App, Form, Col, Row, Space, Button } from 'antd';
import { SearchOutlined, ClearOutlined, DownOutlined } from '@ant-design/icons';
import { SYSTEM_USER_STATUS_MAP, SYSTEM_USER_GENDER_MAP } from '@/common/Map.jsx';
import { GenerateFormItem } from '@/utils/Form.jsx';

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

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 基础数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 全局数据
  const { SystemRoleApiList, SystemRoleMenuList } = useSnapshot(SystemRoleStates);
  // 角色数据
  const [systemRoleList, setSystemRoleList] = useState([]);
  // 部门数据
  const [systemDepartmentList, setSystemDepartmentList] = useState([]);
  // 职位数据
  const [systemJobPositionList, setSystemJobPositionList] = useState([]);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 数据筛选
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 表单字段定义说明：
  // 1. label：表单字段名称
  // 2. name：表单字段名称
  // 3. placeholder：表单字段提示信息
  // 4. type：表单字段类型，可选值：input、passwordInput、select、treeSelect
  // 5. rules：表单字段验证规则
  // 6. search：是否允许搜索
  // 7. tree：是否是树形结构
  // 8. multiple：是否允许多选
  // 9. data：select 类型字段数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 数据筛选 form 表单定义
  const systemUserFilterFields = [
    { label: '工号', name: 'jobNumber', placeholder: '使用工号进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选工号长度不能超过 30 个字符' }] },
    { label: '用户名', name: 'username', placeholder: '使用用户名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选用户名长度不能超过 30 个字符' }] },
    { label: '中文名', name: 'cnName', placeholder: '使用中文名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选中文名长度不能超过 30 个字符' }] },
    { label: '英文名', name: 'enName', placeholder: '使用英文名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选英文名长度不能超过 30 个字符' }] },
    { label: '邮箱', name: 'email', placeholder: '使用邮箱进行筛选过滤', type: 'input', rules: [{ max: 50, message: '筛选邮箱长度不能超过 50 个字符' }] },
    { label: '手机号', name: 'phone', placeholder: '使用手机号进行筛选过滤', type: 'input', rules: [{ max: 11, message: '筛选手机号长度不能超过 11 个字符' }] },
    { label: '状态', name: 'status', placeholder: '使用状态进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, data: SYSTEM_USER_STATUS_MAP, rules: [] },
    { label: '性别', name: 'gender', placeholder: '使用性别进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, data: SYSTEM_USER_GENDER_MAP, rules: [] },
    { label: '角色', name: 'systemRole', placeholder: '使用角色进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, data: systemRoleList, rules: [] },
    { label: '部门', name: 'systemDepartment', placeholder: '使用部门进行筛选过滤', type: 'select', search: true, tree: true, multiple: false, data: systemDepartmentList, rules: [] },
    { label: '职位', name: 'systemJobPosition', placeholder: '使用职位进行筛选过滤', type: 'select', search: true, tree: true, multiple: false, data: systemJobPositionList, rules: [] }
  ];

  // 生成筛选表单
  const [systemUserFilterForm] = Form.useForm();

  // 是否展开更多筛选
  const [expandSystemUserFilterForm, setExpandSystemUserFilterForm] = useState(false);

  // 生成筛选表单 Form.Item
  const generateSystemUserFilterFormItem = () => {
    return systemUserFilterFields.slice(0, expandSystemUserFilterForm ? systemUserFilterFields.length : PageConfig.defaultFilterExpandItemCount).map((field) => (
      <Col span={6} key={field?.name}>
        {GenerateFormItem(field)}
      </Col>
    ));
  };

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
      <div className="admin-page-main">
        {/* 搜索栏 */}
        <div className="admin-page-search">
          <Form form={systemUserFilterForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="systemUserFilterForm" onFinish={() => {}} autoComplete="off">
            <Row gutter={24}>
              {generateSystemUserFilterFormItem()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">条件筛选</Button>
                  <Button icon={<ClearOutlined />} onClick={() => systemUserFilterForm.resetFields()}>清理条件</Button>
                  {systemUserFilterFields.length > PageConfig.defaultFilterExpandItemCount && (
                    <a onClick={() => setExpandSystemUserFilterForm(!expandSystemUserFilterForm)}>
                      <DownOutlined rotate={expandSystemUserFilterForm ? 180 : 0} /> {expandSystemUserFilterForm ? '收起条件' : '展开更多'}
                    </a>
                  )}
                </Space>
              </Col>
            </Row>
          </Form>
        </div>
      </div>
    </>
  );
};

export default SystemUser;
