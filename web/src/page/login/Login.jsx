import React from 'react';
import { Helmet } from 'react-helmet';
import { useNavigate } from 'react-router-dom';
import { App, Button, Checkbox, Divider, Form, Input } from 'antd';
import { DingtalkOutlined, InsuranceOutlined, LockOutlined, UserOutlined } from '@ant-design/icons';
import { TitleSuffix } from '@/common/Text.jsx';
import { BackendApiPrefix, BackendApiSuffix } from '@/common/Api.jsx';
import { AxiosPOST } from '@/handler/Request.jsx';
import { SetLocalTokenAndExpireTime } from '@/handler/Token.jsx';

// 页面常量设置
const PageConfig = {
  // 页面标题
  pageTitle: '登录' + TitleSuffix,
  // 页面描述
  pageDescription: '登录'
};

const Login = () => {
  // 消息提示
  const { message } = App.useApp();
  // 路由跳转
  const navigate = useNavigate();
  // 登录处理方法
  const loginHandler = async (data) => {
    try {
      const loginApi = BackendApiPrefix + BackendApiSuffix.Public.Common.Login.Path;
      const res = await AxiosPOST(loginApi, data);
      if (res.code === 200) {
        message.success('登录成功');
        // 保存 Token 相关数据，然后跳转到主页
        SetLocalTokenAndExpireTime(res.data.token, res.data.expireTime);
        navigate('/dashboard');
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('系统异常，请稍后再试');
    }
  };

  return (
    <>
      {/* 页面 header */}
      <Helmet>
        <title>{PageConfig.pageTitle}</title>
        <meta name="description" content={PageConfig.pageDescription} />
      </Helmet>
      <div className="admin-login-body">
        <div className="admin-login-title">登录 / Sign in</div>
        <Divider className="admin-welcome">欢迎回来</Divider>
        <div className="admin-login-form">
          {/* 账户 */}
          <Form name="login" initialValues={{ remember: true }} onFinish={loginHandler}>
            <Form.Item
              name="account"
              rules={[
                { min: 4, max: 30, message: '账户长度范围为 4-30 个字符' },
                { required: true, message: '请输入您的用户名' }
              ]}
            >
              <Input autoComplete="off" className="admin-login-input" prefix={<UserOutlined />} placeholder="工号 / 用户名 / 手机号 / Email" />
            </Form.Item>
            {/* 密码 */}
            <Form.Item
              className="admin-login-form-item"
              name="password"
              rules={[
                { min: 8, max: 30, message: '密码长度范围为 8-30 个字符' },
                { required: true, message: '请输入您的密码' }
              ]}
            >
              <Input.Password autoComplete="off" className="admin-login-input" prefix={<LockOutlined />} type="password" placeholder="密码" />
            </Form.Item>
            {/*手机令牌*/}
            <Form.Item
              className="admin-login-form-item"
              name="code"
              rules={[
                { min: 4, max: 4, message: '手机令牌验证码应该是 4 个数字' },
                { required: true, message: '请输入您的验证码' }
              ]}
            >
              <Input autoComplete="off" className="admin-login-input" prefix={<InsuranceOutlined />} placeholder="手机令牌验证码" />
            </Form.Item>

            {/*邮件短信获取验证码方式*/}
            {/*<Form.Item style={{ marginBottom: 15 }}>*/}
            {/*  <Space direction="horizontal">*/}
            {/*    <Input*/}
            {/*      autoComplete="off"*/}
            {/*      className="admin-login-input"*/}
            {/*      prefix={<MailOutlined className="site-forms-item-icon" />}*/}
            {/*      placeholder="邮件 / 短信验证码"*/}
            {/*      style={{*/}
            {/*        width: '200px'*/}
            {/*      }}*/}
            {/*    />*/}
            {/*    <Button type="primary" className="admin-login-code-button">获取验证码</Button>*/}
            {/*  </Space>*/}
            {/*</Form.Item>*/}

            {/* 记住密码 */}
            <Form.Item className="admin-login-remember-item">
              <Form.Item name="remember" valuePropName="checked" noStyle>
                <Checkbox>记住密码 |</Checkbox>
              </Form.Item>
              <span className="login-form-forgot">
                忘记密码？<a href="">找回密码</a>
              </span>
            </Form.Item>
            {/* 登录按钮 */}
            <Form.Item style={{ margin: 0 }}>
              <Button block type="primary" htmlType="submit" className="admin-login-form-button">
                登录
              </Button>
            </Form.Item>
            <Divider className="admin-login-change">或者使用钉钉扫码直接登录</Divider>
            {/* 钉钉扫码登录 */}
            <Button className="admin-login-form-button" block icon={<DingtalkOutlined />}>
              切换到钉钉扫码登录
            </Button>
          </Form>
        </div>
      </div>
    </>
  );
};

export default Login;
