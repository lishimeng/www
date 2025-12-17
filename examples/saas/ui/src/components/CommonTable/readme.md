# 通用表格组件
封装了增删改查、分页器、条件筛选的通用ui组件
## Usage
### props
- `base-url` - 数据源URL，组件会自动从此URL请求数据。需要实现以下Api接口：
  - GET `base-url/list`
  - POST `base-url/`
  - DELETE `base-url/{pk}`
  - PUT `base-url/{pk}`
- `name` - 表格数据类型的名称，用于对话框显示
- `columns` - 表格的数据列，`TableColumns[]`类型
  ```ts
  export interface TableColumns {
    prop: string; // 属性名
    label: string; // 表头标签显示
    component?: 'text' | 'enum' | 'time';
    width?: string;
    showOverflowTooltip?: boolean;
    enums?: { label: string; value: any }[];
  }
  ```
- `search-items` - 筛选搜索的选项，`SearchItem[]`类型
  ```ts
  export interface SearchItem {
    prop: string; // 属性名，发送请求时作为url查询参数的键
    label: string;
    placeholder?: string; 
    component: 'input' | 'select' | 'date-picker';
    options?: { label: string; value: any }[];  // select 类型需要使用的选项
  }
  ```

### 插槽
- `form` - 插入表单，对单个表格元素进行编辑
  ```html
  <template #form="{data, ins, create}">
    <el-form
      :ref="(el) => { ins.value = el }"  // 绑定ref，以便CommonTable组件中访问表单
      :model="data"
      :rules="formRules"
      label-width="100px"
    >
  </template>
  ```
