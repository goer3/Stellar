import { AxiosGET } from '@/handler/Request.jsx';

// 生成将 Name/Id 转换成 Label/Value 的 Select 组件数据
const GenerateNameIdSelectDataListHandler = async (api, setter, params = {}) => {
  try {
    const res = await AxiosGET(api, params);
    if (res.code === 200) {
      // 处理数据，转换字段
      const data = res?.data?.list || [];
      const result = data.map((item) => ({
        label: item.name,
        value: item.id
      }));
      setter(result);
    } else {
      console.log(res.message);
    }
  } catch (error) {
    console.log('接口请求失败：', api);
    console.log(error);
  }
};

// 生成将 Name/Id 转换成 Title/Value 的 TreeSelect 组件数据
const GenerateNameIdSelectDataTree = (dataList, parentId) => {
  const tree = [];
  for (const item of dataList) {
    if (item.parentId === parentId) {
      const node = {
        title: item.name,
        value: item.id,
        children: GenerateNameIdSelectDataTree(dataList, item.id)
      };
      // 如果没有子节点，则不包含 children 字段
      if (node.children.length === 0) {
        delete node.children;
      }
      tree.push(node);
    }
  }
  return tree;
};

// 生成将 Name/Id 转换成 Title/Value 的 TreeSelect 组件数据
const GenerateNameIdSelectDataTreeHandler = async (api, setter, params = {}) => {
  try {
    const res = await AxiosGET(api, params);
    if (res.code === 200) {
      // 处理数据，转换字段
      const data = res?.data?.list || [];
      if (data.length > 0) {
        // 生成 Tree 结构
        const tree = GenerateNameIdSelectDataTree(data, 0);
        console.log(tree);
        setter(tree);
      }
    } else {
      console.log(res.message);
    }
  } catch (error) {
    console.log('接口请求失败：', api);
    console.log(error);
  }
};

export { GenerateNameIdSelectDataListHandler, GenerateNameIdSelectDataTreeHandler };
