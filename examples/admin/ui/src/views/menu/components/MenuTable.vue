<template>
  <div>
    <CommonTable
      ref="tableRef"
      :url="url"
      :columns="columns"
      :search-items="searchItems"
      full-screen
      default-expand-all
      show-operations-column
      :query-params="{menuGroup: props.project}"
    >
      <template #name="scope">
        <!--    自定义列    -->
        <template v-if="scope.row.icon && scope.row.icon.startsWith('el-icon')">
          <el-icon style="vertical-align: -0.15em">
            <component :is="scope.row.icon.replace('el-icon-', '')" />
          </el-icon>
        </template>
        <template v-else-if="scope.row.icon">
          <div :class="`i-svg:${scope.row.icon}`" />
        </template>
        {{ scope.row.name }}
      </template>
      <template #head>
        <div class="flex mb-3">
          <el-button type="primary" @click="createItem">新增</el-button>
        </div>
      </template>
      <template #routeName="scope">
        <template v-if="scope.row.type !== MenuTypeEnum.MENU">
          <span class="flex justify-center text-[var(&#45;&#45;el-text-color-secondary)]">-</span>
        </template>
      </template>
      <template #component="scope">
        <template v-if="scope.row.type !== MenuTypeEnum.MENU">
          <span class="flex justify-center text-[var(&#45;&#45;el-text-color-secondary)]">-</span>
        </template>
      </template>
      <template #perm="scope">
          <ForeignKey
            v-if="scope.row.hasPerm"
            :url="permissionUrl"
            :code="scope.row.perm"
          ></ForeignKey>
          <span v-else class="text-[var(&#45;&#45;el-text-color-secondary)]">
            公开
          </span>
      </template>
      <template #operation="scope">
        <div class="flex">
          <el-button type="text" size="small" @click="updateItem(scope.row)">修改</el-button>
          <el-button
            v-if="scope.row.type === MenuTypeEnum.CATALOG || scope.row.isMeasure"
            type="text"
            size="small"
            @click="createChild(scope.row)"
          >新增子项</el-button>
          <el-button type="text" size="small" @click="deleteItem(scope.row)">删除</el-button>
        </div>
      </template>
    </CommonTable>
    <el-drawer
      v-model="drawerState.visible"
      :title="drawerState.title"
      :size="drawerSize"
    >
      <el-form
        ref="formRef"
        :model="drawerState.data"
        :rules="rules"
        label-width="100px"
      >
        <el-form-item label="父级菜单" prop="parentId">
          <el-tree-select
            v-model="drawerState.data.parentId"
            placeholder="选择上级菜单"
            :data="menuOptions"
            filterable
            check-strictly
            default-expand-all
            :render-after-expand="false"
          />
        </el-form-item>

        <el-form-item label="菜单名称" prop="name">
          <el-input v-model="drawerState.data.name" placeholder="请输入菜单名称" />
        </el-form-item>

        <el-form-item label="菜单类型" prop="type">
          <el-radio-group v-model="drawerState.data.type">
            <el-radio :value="MenuTypeEnum.CATALOG">目录</el-radio>
            <el-radio :value="MenuTypeEnum.MENU">菜单</el-radio>
            <!-- <el-radio :value="MenuTypeEnum.BUTTON">按钮</el-radio> -->
            <el-radio :value="MenuTypeEnum.EXTLINK">外链</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="drawerState.data.type == MenuTypeEnum.EXTLINK" label="外链地址" prop="path">
          <el-input v-model="drawerState.data.routePath" placeholder="请输入外链完整路径" />
        </el-form-item>

        <TooltipFormItem
          v-if="drawerState.data.type == MenuTypeEnum.CATALOG || drawerState.data.type == MenuTypeEnum.MENU"
          prop="routePath"
          label="路由路径"
          tooltip="相对于上级目录的 URL 相对路径，以 “/”开头。"
        >
          <el-input
            v-if="drawerState.data.type == MenuTypeEnum.CATALOG"
            v-model="drawerState.data.routePath"
            placeholder="/system"
          />
          <el-input v-else v-model="drawerState.data.routePath" placeholder="/user" />
        </TooltipFormItem>

        <TooltipFormItem
          v-if="drawerState.data.type == MenuTypeEnum.MENU"
          prop="routeName"
          label="路由名称"
          tooltip="菜单的唯一名称，可用于页面缓存与 router.push 跳转。"
        >
          <el-input v-model="drawerState.data.routeName" placeholder="User" />
        </TooltipFormItem>

        <TooltipFormItem
          v-if="drawerState.data.type == MenuTypeEnum.MENU"
          prop="component"
          label="组件路径"
          tooltip="组件页面文件路径，相对于 src/views/，如 system/user/index，省略后缀 .vue"
        >
          <el-input
            v-model="drawerState.data.component"
            placeholder="system/user/index"
            style="width: 95%"
          >
            <template v-if="drawerState.data.type == MenuTypeEnum.MENU" #prepend>src/views/</template>
            <template v-if="drawerState.data.type == MenuTypeEnum.MENU" #append>.vue</template>
          </el-input>
        </TooltipFormItem>

        <TooltipFormItem
          v-if="drawerState.data.type == MenuTypeEnum.MENU"
          label="路由参数"
          tooltip="组件页面使用 `useRoute().query.参数名` 获取路由参数值。"
        >

          <div v-if="!drawerState.data.params || drawerState.data.params.length === 0">
            <el-button type="success" plain @click="drawerState.data.params = [{ key: '', value: '' }]">
              添加路由参数
            </el-button>
          </div>

          <div v-else>
            <div v-for="(item, index) in drawerState.data.params" :key="index">
              <el-input v-model="item.key" placeholder="参数名" style="width: 100px" />

              <span class="mx-1">=</span>

              <el-input v-model="item.value" placeholder="参数值" style="width: 100px" />

              <el-icon
                v-if="drawerState.data.params.indexOf(item) === drawerState.data.params.length - 1"
                class="ml-2 cursor-pointer color-[var(--el-color-success)]"
                style="vertical-align: -0.15em"
                @click="drawerState.data.params.push({ key: '', value: '' })"
              >
                <CirclePlusFilled />
              </el-icon>
              <el-icon
                class="ml-2 cursor-pointer color-[var(--el-color-danger)]"
                style="vertical-align: -0.15em"
                @click="drawerState.data.params.splice(drawerState.data.params.indexOf(item), 1)"
              >
                <DeleteFilled />
              </el-icon>
            </div>
          </div>
        </TooltipFormItem>

        <TooltipFormItem
          label="权限"
          prop="hasPerm"
          tooltip="子菜单会自动继承父级菜单的权限，无需逐一配置"
        >
          <el-radio-group v-model="drawerState.data.hasPerm">
            <el-radio :value="false">无需权限</el-radio>
            <el-radio :value="true">要求权限</el-radio>
          </el-radio-group>
        </TooltipFormItem>

        <el-form-item
          v-if="drawerState.data.hasPerm"
          label="权限分组"
          prop="perm"
        >
          <TableSelect
            v-model="drawerState.data.perm"
            :url="permissionUrl"
            placeholder="请选择目录权限"
            title="选择目录权限"
            :columns="permissionColumns"
          />
        </el-form-item>

        <el-form-item label="启用状态" prop="status">
          <el-radio-group v-model="drawerState.data.status">
            <el-radio :value="1">启用</el-radio>
            <el-radio :value="0">停用</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item
          v-if="drawerState.data.type !== MenuTypeEnum.BUTTON"
          prop="visible"
          label="显示状态"
        >
          <el-radio-group v-model="drawerState.data.visible">
            <el-radio :value="1">显示</el-radio>
            <el-radio :value="0">隐藏</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="drawerState.data.type === MenuTypeEnum.MENU" label="缓存页面" prop="keepAlive">
          <el-radio-group v-model="drawerState.data.keepAlive">
            <el-radio :value="1">开启</el-radio>
            <el-radio :value="0">关闭</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="排序" prop="sort">
          <el-input-number
            v-model="drawerState.data.sort"
            style="width: 100px"
            controls-position="right"
            :min="0"
          />
        </el-form-item>

        <!-- 权限标识 -->
        <el-form-item v-if="drawerState.data.type == MenuTypeEnum.BUTTON" label="权限标识" prop="perm">
          <el-input v-model="drawerState.data.perm" placeholder="sys:user:add" />
        </el-form-item>

        <el-form-item v-if="drawerState.data.type !== MenuTypeEnum.BUTTON" label="图标" prop="icon">
          <!-- 图标选择器 -->
          <icon-select v-model="drawerState.data.icon" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="drawerState.visible = false">取消</el-button>
          <el-button type="primary" @click="submitForm">确定</el-button>
        </span>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import {SearchItem, TableColumns} from "@/components/CommonTable/types";
import { useAppStore } from "@/store/modules/app.store";
import { DeviceEnum } from "@/enums/settings/device.enum";
import MenuAPI from "@/api/system/menu.api";
import { withConfirm } from "@/composables/useConfirmDialog";
import { createApi, deleteApi, updateApi } from "@/api/common/tabel.api";
import {permissionColumns, permissionUrl} from "@/types/tables/permission.table";

const url = "/api/menus"

const appStore = useAppStore();
const tableRef = ref();
const formRef = ref();

const drawerState = reactive({
  visible: false,
  title: "",
  data: {} as Record<string, any>,
  isCreate: false,
});

enum MenuTypeEnum {
  MENU = 1, // 菜单
  CATALOG = 2, // 目录
  EXTLINK = 3, // 外链
  BUTTON = 4, // 按钮
}

const props = defineProps({
  project: {
    type: String,
    required: true,
  }
});

const columns: TableColumns[] = [
  {
    prop: "name",
    label: "菜单名称",
    component: "custom",
    fixed: true,
    width: "200px",
  },
  {
    prop: "type",
    label: "类型",
    component: "enum_tag",
    enums: [
      { value: 0, label: "Api", type: "danger" },
      { value: MenuTypeEnum.CATALOG, label: "目录", type: "success" },
      { value: MenuTypeEnum.MENU, label: "菜单", type: "warning" },
      { value: MenuTypeEnum.EXTLINK, label: "外链", type: "primary" },
    ],
  },
  {
    prop: "routeName",
    label: "路由名称",
    component: "custom",
  },
  {
    prop: "routePath",
    label: "路由路径",
    component: "text",
  },
  {
    prop: "component",
    label: "组件路径",
    component: "custom",
  },
  {
    prop: "perm",
    label: "权限分组",
    component: "custom",
  },
  {
    prop: "status",
    label: "状态",
    component: "enum_tag",
    enums: [
      { value: 0, label: "停用", type: "info" },
      { value: 1, label: "启用", type: "success" },
    ],
  },
  {
    prop: "visible",
    label: "显示状态",
    component: "enum_tag",
    enums: [
      { value: 0, label: "隐藏", type: "primary" },
      { value: 1, label: "显示", type: "success" },
    ],
  },
  {
    prop: "sort",
    label: "排序",
    component: "text",
  },
  {
    prop: "createTime",
    label: "创建时间",
    component: "time",
  },
  {
    prop: "updateTime",
    label: "更新时间",
    component: "time",
  },
];
const searchItems: SearchItem[] = [
  { prop: "keywords", label: "关键字", placeholder: "名称或路由名称（模糊查询）", component: "input" },
];

const rules = reactive({
  parentId: [{ required: true, message: "请选择上级菜单", trigger: "blur" }],
  name: [{ required: true, message: "请输入菜单名称", trigger: "blur" }],
  routeName: [{ required: true, message: "请输入路由名称", trigger: "blur" }],
  perm: [{ required: true, message: "请选择目录权限", trigger: [] }],
  type: [{ required: true, message: "请选择菜单类型", trigger: "blur" }],
  routePath: [{ required: true, message: "请输入路由路径", trigger: "blur" }],
  component: [{ required: true, message: "请输入组件路径", trigger: "blur" }],
});

const drawerSize = computed(() => (appStore.device === DeviceEnum.DESKTOP ? "600px" : "90%"));

const menuOptions = ref<OptionType[]>([]);

const initialFormData = {
  id: undefined,
  parentId: "0",
  visible: 1,
  sort: 1,
  type: MenuTypeEnum.CATALOG, // 默认菜单
  group: props.project,
  status: 1,
  keepAlive: 1,
  hasPerm: false,
  perm: "",
  params: [],
};


async function getOptions(id?: number) {
  const data = await MenuAPI.getOptions(id, true, props.project);
  menuOptions.value = [{ value: "0", label: "顶级菜单", children: data }];
}


async function createItem() {
  drawerState.isCreate = true;
  await getOptions();
  drawerState.data = { ...initialFormData };
  drawerState.title = "新增菜单";
  drawerState.visible = true;
}

async function createChild(row: Record<string, any>) {
  drawerState.isCreate = true;
  await getOptions();
  drawerState.data = { ...initialFormData };
  drawerState.data.parentId = row.id;
  drawerState.title = "新增菜单";
  drawerState.visible = true;
}

async function updateItem(row: Record<string, any>) {
  drawerState.isCreate = false;
  drawerState.data = { ...row };
  drawerState.title = "修改菜单";
  await getOptions(row.id);
  drawerState.visible = true;
}

async function deleteItem(row: Record<string, any>) {
  console.log();
  const result = await withConfirm({
    message: "删除的数据不可恢复，是否继续？",
    title: `删除菜单：${row.name}`,
  });
  if (result) {
    deleteApi(url, row.id)
      .then(() => {
        ElMessage.success("删除成功！");
        tableRef.value?.fetchPageData();
      })
      .catch((err) => {
        console.log(err);
      });
  }
}

function submitForm() {
  formRef.value
    .validate()
    .then((valid: boolean) => {
      if (valid) {
        if (drawerState.isCreate) {
          createApi(url, drawerState.data)
            .then(() => {
              ElMessage.success("新增菜单成功！");
              drawerState.visible = false;
              tableRef.value?.fetchPageData();
            })
            .catch((err) => {
              console.log(err);
            });
        } else {
          updateApi(url, drawerState.data.id, drawerState.data)
            .then(() => {
              ElMessage.success("修改成功！");
              drawerState.visible = false;
              tableRef.value?.fetchPageData();
            })
            .catch((err) => {
              console.log(err);
            });
        }
      }
    })
    .catch((err: any) => {
      console.log(err);
    });
}

</script>
