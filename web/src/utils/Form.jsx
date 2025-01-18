import { Form, Input, Select, TreeSelect } from 'antd';

// 组件类型对应的组件配置
const ComponentTypeMap = {
  // 输入框
  input: (props) => <Input {...props} autoComplete="off" />,
  // 密码输入框
  passwordInput: (props) => <Input.Password {...props} autoComplete="off" />,
  // 文本输入框
  textarea: (props) => <Input.TextArea {...props} autoComplete="off" />,
  // 下拉框
  select: (props) => <Select {...props} autoComplete="off" />,
  // 树形下拉框
  treeSelect: (props) => <TreeSelect {...props} autoComplete="off" treeDefaultExpandAll={true} />
};

// 传入字段定义信息，返回 Form.Item
const GenerateFormItem = (field) => {
  // 通用配置，所有都会用到的公共属性
  const commonConfig = {
    allowClear: true, // 是否允许显示清空图标，都允许显示
    placeholder: field?.placeholder, // 输入框占位符
    hidden: field?.hidden, // 是否隐藏该字段
    disabled: field?.disabled, // 是否禁用该字段
    value: field?.value // 回填的值
  };

  // 定有组件特有的属性
  const componentConfig = {
    input: {},
    passwordInput: {},
    textarea: {},
    select: {
      options: field?.data, // 下拉框数据
      showSearch: field?.search, // 是否允许搜索
      optionFilterProp: 'label', // 选择器过滤字段
      allowClear: false, // 是否允许清空，覆盖公共配置
      multiple: field?.multiple // 是否允许多选
    },
    treeSelect: {
      options: field?.data, // 下拉框数据
      showSearch: field?.search, // 是否允许搜索
      treeNodeFilterProp: 'label', // 树形选择器过滤字段
      allowClear: false, // 是否允许清空，覆盖公共配置
      mode: field?.multiple ? 'multiple' : 'default' // 是否允许多选
    }
  };

  // 根据字段类型，生成对应的 Form.Item
  return (
    <Form.Item key={field?.name} name={field?.name} label={field?.label} hidden={field?.hidden} rules={field?.rules}>
      {ComponentTypeMap[field?.type]?.({
        ...commonConfig,
        ...componentConfig[field?.type]
      })}
    </Form.Item>
  );
};

export { ComponentTypeMap, GenerateFormItem };
