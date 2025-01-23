import { AxiosGET } from "@/handler/Request.jsx";

// 生成将 Name/Id 转换成 Label/Value 的 Select 组件数据
const GenerateNameIdToLabelValueSelectDataList = async (api, setter, params = {}) => {
    try {
        const res = await AxiosGET(api, params);
        if (res.code === 200) {
            // 处理数据，转换字段
            const data = res?.data?.list || [];
            const result = data.map(item => ({
                label: item.name,
                value: item.id
            }));
            setter(result);
        } else {
            console.log(res.message);
        }
    } catch (error) {
        console.log("接口请求失败：", api);
        console.log(error);
    }
}

// 生成 TreeSelect 组件数据
const GenerateSelectDataTree = (api, setter, tree = false, params = {}) => {

}

export { GenerateNameIdToLabelValueSelectDataList, GenerateSelectDataTree };
