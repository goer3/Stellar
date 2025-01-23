import React, { useState, useEffect } from 'react';
import { useSnapshot } from 'valtio';
import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { App, Form, Col, Row, Space, Button, Avatar, Table, Descriptions, Dropdown, Popconfirm } from 'antd';
import {
  SearchOutlined,
  ClearOutlined,
  DownOutlined,
  UserAddOutlined,
  UploadOutlined,
  DownloadOutlined,
  ClockCircleOutlined,
  EditOutlined,
  DeleteOutlined,
  RestOutlined,
  LockOutlined
} from '@ant-design/icons';
import { SYSTEM_USER_STATUS_MAP, SYSTEM_USER_GENDER_MAP } from '@/common/Map.jsx';
import { GenerateFormItem } from '@/utils/Form.jsx';
import { AxiosGET } from '@/handler/Request.jsx';
import { BackendApiPrefix, BackendApiSuffix } from '@/common/Api.jsx';
import { GenerateGenderIcon, GenerateStatusTag, GenerateRoleTag } from '@/common/Tag.jsx';
import { GenerateNameIdSelectDataListHandler, GenerateNameIdSelectDataTreeHandler } from '@/utils/Select.jsx';

// 页面常量设置
const PageConfig = {
  // 默认页码
  defaultPageNumber: 1,
  // 默认每页显示的数据量
  defaultPageSize: 2,
  // 默认数据总数
  defaultTotal: 0,
  // 默认是否需要分页
  defaultPaginable: true,
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
  // 全局数据
  const { SystemRoleApiList, SystemRoleMenuList } = useSnapshot(SystemRoleStates);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 数据筛选
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 职位接口地址
  const SystemJobPositionListApi = BackendApiPrefix + BackendApiSuffix.System.JobPosition.AuthAndPermission.List.Path;
  // 部门接口地址
  const SystemDepartmentListApi = BackendApiPrefix + BackendApiSuffix.System.Department.AuthAndPermission.List.Path;
  // 角色接口地址
  const SystemRoleListApi = BackendApiPrefix + BackendApiSuffix.System.Role.AuthAndPermission.List.Path;

  // 职位数据
  const [systemJobPositionList, setSystemJobPositionList] = useState([]);
  // 部门数据
  const [systemDepartmentList, setSystemDepartmentList] = useState([]);
  // 角色数据
  const [systemRoleList, setSystemRoleList] = useState([]);

  // 获取系统数据
  useEffect(() => {
    // 获取职位数据
    GenerateNameIdSelectDataListHandler(SystemJobPositionListApi, setSystemJobPositionList);
    // 获取部门数据，生成树形结构
    GenerateNameIdSelectDataTreeHandler(SystemDepartmentListApi, setSystemDepartmentList);
    // 获取角色数据
    GenerateNameIdSelectDataListHandler(SystemRoleListApi, setSystemRoleList);
  }, []);

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
  // 9. allowClear：是否允许清空
  // 10. data：select 类型字段数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 数据筛选 form 表单定义
  const systemUserFilterFields = [
    { label: '工号', name: 'jobNumber', placeholder: '使用工号进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选工号长度不能超过 30 个字符' }] },
    { label: '用户名', name: 'username', placeholder: '使用用户名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选用户名长度不能超过 30 个字符' }] },
    { label: '中文名', name: 'cnName', placeholder: '使用中文名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选中文名长度不能超过 30 个字符' }] },
    { label: '英文名', name: 'enName', placeholder: '使用英文名进行筛选过滤', type: 'input', rules: [{ max: 30, message: '筛选英文名长度不能超过 30 个字符' }] },
    { label: '邮箱', name: 'email', placeholder: '使用邮箱进行筛选过滤', type: 'input', rules: [{ max: 50, message: '筛选邮箱长度不能超过 50 个字符' }] },
    { label: '手机号', name: 'phone', placeholder: '使用手机号进行筛选过滤', type: 'input', rules: [{ max: 11, message: '筛选手机号长度不能超过 11 个字符' }] },
    { label: '状态', name: 'status', placeholder: '通过选择状态进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, allowClear: true, data: SYSTEM_USER_STATUS_MAP, rules: [] },
    { label: '性别', name: 'gender', placeholder: '通过选择性别进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, allowClear: true, data: SYSTEM_USER_GENDER_MAP, rules: [] },
    { label: '角色', name: 'systemRole', placeholder: '通过选择角色进行筛选过滤', type: 'select', search: true, tree: false, multiple: false, allowClear: true, data: systemRoleList, rules: [] },
    { label: '部门', name: 'systemDepartment', placeholder: '通过选择部门进行筛选过滤', type: 'treeSelect', search: true, tree: true, multiple: false, allowClear: true, data: systemDepartmentList, rules: [] },
    { label: '职位', name: 'systemJobPosition', placeholder: '通过选择职位进行筛选过滤', type: 'select', search: true, tree: true, multiple: false, allowClear: true, data: systemJobPositionList, rules: [] }
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
  // 表格数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 用户操作按钮组
  const systemUserActionButtons = (record) => {
    return (
      <>
        <Space>
          <Button color="primary" variant="link" icon={<EditOutlined />} onClick={() => {}}>
            编辑
          </Button>
          {record.status === 1 ? (
            <Popconfirm
              placement="topRight"
              title="确定要禁用该用户吗？"
              okText="确定"
              cancelText="取消"
              okButtonProps={{ style: { backgroundColor: '#ff4d4f', borderColor: '#ff4d4f' } }}
              onConfirm={() => {}}
            >
              <Button color="danger" variant="link" icon={<DeleteOutlined />}>
                禁用
              </Button>
            </Popconfirm>
          ) : (
            <Popconfirm
              placement="topRight"
              title="确定要启用该用户吗？"
              okText="确定"
              cancelText="取消"
              okButtonProps={{ style: { backgroundColor: '#52c41a', borderColor: '#52c41a' } }}
              onConfirm={() => {}}
            >
              <Button color="success" variant="link" icon={<RestOutlined />}>
                启用
              </Button>
            </Popconfirm>
          )}
          <Button color="primary" variant="link" icon={<LockOutlined />} onClick={() => {}}>
            重置
          </Button>
        </Space>
      </>
    );
  };

  // 表格列定义
  const systemUserTableColumns = [
    { title: '头像', dataIndex: 'avatar', key: 'avatar', render: (avatar) => <Avatar src={avatar} /> },
    { title: '性别', dataIndex: 'gender', key: 'gender', render: (gender) => GenerateGenderIcon(gender) },
    { title: '工号', dataIndex: 'jobNumber', key: 'jobNumber' },
    { title: '用户名', dataIndex: 'username', key: 'username' },
    { title: '中文名', dataIndex: 'cnName', key: 'cnName' },
    { title: '英文名', dataIndex: 'enName', key: 'enName' },
    { title: '邮箱', dataIndex: 'email', key: 'email' },
    { title: '手机号', dataIndex: 'phone', key: 'phone' },
    { title: '状态', dataIndex: 'status', key: 'status', render: (status) => GenerateStatusTag(status) },
    { title: '角色', dataIndex: ['systemRole', 'name'], key: 'systemRole', render: (role) => GenerateRoleTag(role) },
    { title: '部门', dataIndex: 'systemDepartments', key: 'systemDepartments', render: (departments) => departments?.map((item) => item.name).join(',') },
    { title: '职位', dataIndex: 'systemJobPositions', key: 'systemJobPositions', render: (positions) => positions?.map((item) => item.name).join(',') },
    { title: '操作', key: 'action', fixed: 'right', width: 150, render: (_, record) => systemUserActionButtons(record) }
  ];

  // 每页显示的数据条数
  const [pageSize, setPageSize] = useState(PageConfig.defaultPageSize);
  // 当前页码
  const [pageNumber, setPageNumber] = useState(PageConfig.defaultPageNumber);
  // 数据总数
  const [total, setTotal] = useState(PageConfig.defaultTotal);
  // 是否需要分页
  const [paginable, setPaginable] = useState(PageConfig.defaultPaginable);
  // 筛选参数
  const [systemUserListFilterParams, setSystemUserListFilterParams] = useState({});
  // 表格数据
  const [systemUserTableDataSource, setSystemUserTableDataSource] = useState([]);

  // 查询和筛选用户列表方法封装
  const getSystemUserListRequest = async (params) => {
    try {
      const systemUserListApi = BackendApiPrefix + BackendApiSuffix.System.User.AuthAndPermission.List.Path;
      const res = await AxiosGET(systemUserListApi, params);
      if (res.code === 200) {
        setSystemUserTableDataSource(res.data?.list);
        setPageSize(res.data?.pagination?.pageSize);
        setPageNumber(res.data?.pagination?.pageNumber);
        setTotal(res.data?.pagination?.total);
        setPaginable(res.data?.pagination?.paginable);
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.error(error);
      message.error('系统异常，请稍后再试');
    }
  };

  // 手动筛选方法，需要先初始化页码，避免在结果溢出数据的页码的时候，导致请求参数中带了页码，无法请求到数据
  // 比如在第三页的时候筛选，但是结果只有两页，则会因为页码问题，显示没有数据
  const getSystemUserListHandler = (params) => {
    setPageNumber(PageConfig.defaultPageNumber);
    setPageSize(PageConfig.defaultPageSize);
    setTotal(PageConfig.defaultTotal);
    setPaginable(PageConfig.defaultPaginable);
    setSystemUserListFilterParams(params);
  };

  // 监听展示条件变化，然后请求数据
  useEffect(() => {
    const params = { ...systemUserListFilterParams, pageSize, pageNumber, paginable };
    getSystemUserListRequest(params);
  }, [pageSize, pageNumber, paginable, systemUserListFilterParams]);

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
          <Form form={systemUserFilterForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="systemUserFilterForm" onFinish={getSystemUserListHandler} autoComplete="off">
            <Row gutter={24}>
              {generateSystemUserFilterFormItem()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">
                    条件筛选
                  </Button>
                  <Button
                    icon={<ClearOutlined />}
                    onClick={() => {
                      systemUserFilterForm.resetFields();
                      setSystemUserListFilterParams({});
                    }}
                  >
                    清理条件
                  </Button>
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
        {/* 表格 */}
        <div className="admin-page-list">
          <div className="admin-page-btn-group">
            <Space>
              <Button type="primary" icon={<UserAddOutlined />} onClick={() => {}}>
                添加用户
              </Button>
              <Button icon={<UploadOutlined />} onClick={() => {}}>
                批量导入
              </Button>
              <Dropdown menu="">
                <Button icon={<DownOutlined />}>批量操作</Button>
              </Dropdown>
            </Space>
            <Space style={{ float: 'right' }}>
              <Button icon={<DownloadOutlined />} onClick={() => {}}>
                模板下载
              </Button>
              <Button icon={<ClockCircleOutlined />} onClick={() => {}}>
                导入记录
              </Button>
            </Space>
          </div>
          <Table
            // 表格布局大小
            size="small"
            // 表格布局方式，支持 fixed、auto
            tableLayout="fixed"
            // 表格行选择
            rowSelection={{ type: 'checkbox' }}
            // 表格列
            columns={systemUserTableColumns}
            // 表格展开信息
            expandable={{
              expandedRowRender: (record) => {
                const [username, cnName, enName, uid] = record.creator.split(',');
                const items = [
                  { label: '用户创建者', children: `${cnName} / ${enName} (${username} / ${uid})` },
                  { label: '创建时间', children: record.createdAt || '-' },
                  { label: '更新时间', children: record.updatedAt || '-' },
                  { label: '最后登录 IP', children: record.lastLoginIP || '-' },
                  { label: '最后登录时间', children: record.lastLoginAt || '-' }
                ];
                return <Descriptions column={1} items={items} />;
              },
              // 用于限制是否可以展开
              rowExpandable: (record) => record.name !== 'Not Expandable'
            }}
            dataSource={systemUserTableDataSource}
            // 行唯一标识
            rowKey="id"
            // 表格分页
            pagination={{
              pageSize: pageSize,
              current: pageNumber,
              total: total,
              showTotal: (total) => `总共 ${total} 条记录`,
              // 是否隐藏分页器，当只有一页时隐藏
              // hideOnSinglePage: true,
              // 是否允许修改显示数量
              showSizeChanger: true,
              // 是否显示快速跳转
              showQuickJumper: true,
              // 页码变化时触发
              onChange: (page, pageSize) => {
                setPageNumber(page);
                setPageSize(pageSize);
              }
            }}
            // 表格滚动，目的是为了最后一列固定
            scroll={{ x: 'max-content' }}
          />
        </div>
      </div>
    </>
  );
};

export default SystemUser;
