import axios from 'axios';
import { GetUnexpireToken } from '@/handler/Token.jsx';

// 创建 Axios 配置实例
const instance = axios.create({
  timeout: 5000 // 请求超时时间
});

// 请求拦截器，需要在请求头中添加 Token
instance.interceptors.request.use(
  (config) => {
    config.headers.Authorization = 'Bearer ' + GetUnexpireToken();
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器，处理响应数据
instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// GET 请求
const APIGET = (url, data) => instance.get(url, { params: data }).then((res) => res.data);

// POST 请求
const APIPOST = (url, data) => instance.post(url, data).then((res) => res.data);

// PUT 请求
const APIPUT = (url, data) => instance.put(url, data).then((res) => res.data);

// PATCH 请求
const APIPATCH = (url, data) => instance.patch(url, data).then((res) => res.data);

// DELETE 请求
const APIDELETE = (url) => instance.delete(url).then((res) => res.data);

export { APIGET, APIPOST, APIPUT, APIPATCH, APIDELETE };
