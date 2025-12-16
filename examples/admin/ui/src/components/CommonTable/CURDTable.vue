<template>
  <div>
    <CommonTable
      ref="tableRef"
      v-bind="$attrs"
      :url="props.url"
      :columns="props.columns"
      :pk="props.pk"
      show-operations-column
    >
      <template #operation="scope">
        <slot name="all-operation" :row="scope.row" :update="updateItem" :delete="deleteItem">
          <div class="h-full flex">
            <el-button type="text" size="small" @click="updateItem(scope.row)">修改</el-button>
            <slot name="extra-operation" :row="scope.row"></slot>
            <el-button type="text" size="small" @click="deleteItem(scope.row)">删除</el-button>
          </div>
        </slot>
      </template>
      <template #head>
        <div class="flex mb-3">
          <slot v-if="props.createEnabled" name="create" :create="createItem">
            <el-button type="primary" @click="createItem(props.initialFormData)">新增</el-button>
          </slot>
          <slot name="head"></slot>
        </div>
      </template>
      <template v-for="(slot, slotName) in $slots" #[slotName]="slotProps" :key="slotName">
        <template v-if="slotName !== 'operation' && slotName !== 'head' && slotName !== 'form'">
          <slot :name="slotName" v-bind="slotProps" />
        </template>
      </template>
    </CommonTable>

    <!--  新增/编辑对话框  -->
    <el-dialog
      v-if="!props.useDrawer"
      v-model="state.dialog.visible"
      :title="state.dialog.title"
      width="500px"
      append-to-body
      @closed="onDialogClosed"
    >
      <slot
        name="form"
        :form-data="state.dialog.data"
        :on-update-form-data="(val:Record<string, any>) => (state.dialog.data = val)"
        :form-ins="getFormRef()"
        :is-create="state.dialog.isCreate"
        :visible="state.dialog.visible"
      ></slot>
      <template #footer>
        <span class="dialog-footer">
          <el-button class="w-18" :loading="isSubmitting" @click="state.dialog.visible = false">取消</el-button>
          <el-button class="w-18" type="primary" :loading="isSubmitting" @click="submitForm">确定</el-button>
        </span>
      </template>
    </el-dialog>
    <el-drawer
      v-if="props.useDrawer"
      v-model="state.dialog.visible"
      :title="state.dialog.title"
      :size="props.drawerSize"
      append-to-body
      @closed="onDialogClosed"
    >
      <slot
        name="form"
        :form-data="state.dialog.data"
        :on-update-form-data="(val:Record<string, any>) => (state.dialog.data = val)"
        :form-ins="getFormRef()"
        :is-create="state.dialog.isCreate"
        :visible="state.dialog.visible"
      ></slot>
      <template #footer>
        <span class="dialog-footer">
          <el-button class="w-18" :loading="isSubmitting" @click="state.dialog.visible = false">取消</el-button>
          <el-button class="w-18" type="primary" :loading="isSubmitting" @click="submitForm">
            <template v-if="state.dialog.isCreate">
              提交
            </template>
            <template v-else>
              保存
            </template>
          </el-button>
        </span>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">

import { TableColumns } from "@/components/CommonTable/types";
import { withConfirm } from "@/composables/useConfirmDialog";
import {createApi, deleteApi, patchUpdateApi, updateApi} from "@/api/common/tabel.api";

interface CommonTableProps {
  url: string;
  columns: TableColumns[];
  pk?: string;
  name?: string;
  useDrawer?: boolean;
  initialFormData?: Record<string, any>;
  drawerSize?: number | string;
  createEnabled?: boolean;
  updateByPatch?: boolean;
  updateWithoutPk?: boolean;
}

const props = withDefaults(defineProps<CommonTableProps>(), {
  pk: "id",
  name: "",
  useDrawer: false,
  initialFormData: () => ({}),
  drawerSize: "36%",
  createEnabled: true,
  updateByPatch: false,
  updateWithoutPk: false,
});

const emit = defineEmits(["onCreated","onUpdated","onDeleted","onChanged","onDialogOpen", "onDialogClosed"]);

const state = reactive({
  dialog: {
    visible: false,
    title: "",
    data: {} as Record<string, any>,
    isCreate: false,
  },
});

const isSubmitting = ref(false);

const tableRef = ref();
const formRef = ref();
// 使用getter函数，避免解引用，保证传递formRef对象
const getFormRef = () => formRef;

function fetchPageData() {
  tableRef.value?.fetchPageData();
}

function createItem(row: Record<string, any>) {
  state.dialog.isCreate = true;
  state.dialog.data = JSON.parse(JSON.stringify(row)); // 深拷贝的简单实现
  state.dialog.title = `新增${props.name}`;
  emit("onDialogOpen", state.dialog.data, state.dialog.isCreate);
  state.dialog.visible = true;
}

function updateItem(row: Record<string, any>) {
  state.dialog.isCreate = false;
  state.dialog.data = JSON.parse(JSON.stringify(row));
  state.dialog.title = `修改${props.name}`;
  emit("onDialogOpen", state.dialog.data, state.dialog.isCreate);
  state.dialog.visible = true;
}

async function deleteItem(row: Record<string, any>) {
  const result = await withConfirm({
    message: "删除的数据不可恢复，是否继续？",
    title: `删除${props.name}：${row[nameColumn.value] || row[props.pk]}`,
  });
  if (result) {
    deleteApi(props.url, row[props.pk])
      .then(() => {
        ElMessage.success("删除成功！");
        emit("onDeleted", row);
        emit("onChanged");
        fetchPageData();
      })
      .catch((err) => {
        console.log(err);
      });
  }
}

async function submitForm() {
  if (!formRef.value) {
    console.error("表单引用不存在，请检查插槽使用是否正确");
    ElMessage.error("表单初始化异常，请刷新页面重试");
    return;
  }

  // 防重：如果已经在提交，直接返回
  if (isSubmitting.value) return;

  isSubmitting.value = true;

  try {
    const valid = await formRef.value.validate();
    if (!valid) return;
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (validationError) {
    // 表单验证失败：静默处理，不提示，不打日志
    isSubmitting.value = false;
    return;
  }

  try {
    const valid = await formRef.value.validate();
    if (!valid) return;

    if (state.dialog.isCreate) {
      await createApi(props.url, state.dialog.data);
      ElMessage.success(`新增${props.name}成功！`);
      emit("onCreated", state.dialog.data);
    } else {
      let pk = null;
      if (!props.updateWithoutPk) {
        pk = state.dialog.data[props.pk];
      }
      if (props.updateByPatch) {
        await patchUpdateApi(
          props.url,
          pk,
          state.dialog.data
        );
      } else {
        await updateApi(
          props.url,
          pk,
          state.dialog.data
        );
      }
      ElMessage.success("修改成功！");
      emit("onUpdated", state.dialog.data);
    }

    // 成功后统一处理
    state.dialog.visible = false;
    emit("onChanged");
    fetchPageData();
    // 不应过早退出loading状态，等待对话框完全关闭

  } catch (err: any) {
    // 统一错误处理
    console.error("提交失败:", err);
    // 请求失败有统一的错误提示
    // ElMessage.error(err?.message || "操作失败，请稍后重试");
    isSubmitting.value = false;
  }
}

function onDialogClosed() {
  isSubmitting.value = false;
  emit("onDialogClosed");
}

const nameColumn = computed(() => {
  let hasNameCol = false;
  for (const col of props.columns) {
    if (col.isName) {
      return col.prop;
    } else if (col.prop === "name") {
      hasNameCol = true;
    }
  }
  if (hasNameCol) {
    return "name";
  }
  return props.columns[0].prop;
});

function getDialogData() {
  return state.dialog.data;
}

type UpdateOperation = (data: Record<string, any>) => null

function updateDialogData(h: UpdateOperation) {
  h(state.dialog.data);
}

defineExpose({
  getDialogData,
  updateDialogData,
  createItem,
  updateItem,
  deleteItem,
  fetchPageData,
})

</script>
