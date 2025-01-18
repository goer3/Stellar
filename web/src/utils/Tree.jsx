import { DynamicIcon } from '@/common/Icon.jsx';

// 递归生成菜单项
const GenerateMenuTree = (parentId, menuList) => {
  // 参数验证
  if (!Array.isArray(menuList)) {
    return [];
  }

  const tree = [];
  for (const item of menuList) {
    // 检查必要字段，如果不正确则跳过
    if (!item || !item.key || !item.label) {
      continue;
    }

    // 如果父节点是当前节点，则生成菜单项
    if (item.parentId === parentId) {
      const node = {
        key: item.key,
        label: item.label,
        icon: item.icon ? <DynamicIcon iconName={item.icon} /> : null,
        children: GenerateMenuTree(item.id, menuList)
      };

      // 如果没有子节点，则不包含 children 字段
      if (node.children.length === 0) {
        delete node.children;
      }

      // 将生成的菜单项添加到树中
      tree.push(node);
    }
  }
  return tree;
};

export { GenerateMenuTree };
