<template>
  <div>
    <CURDTable
      ref="tableRef"
      :url="roleUrl"
      name="角色"
      pk="code"
      :columns="roleColumns"
      full-screen
      table-layout="fixed"
      :initial-form-data="initialFormData"
      update-by-patch
      update-without-pk
    >
      <template #form="{ formData, formIns, isCreate }">
        <el-form
          :ref="(el) => {formIns.value = el}"
          :model="formData"
          :rules="formRules"
          label-width="100px"
        >
          <el-form-item label="编码" prop="code">
            <!--          <BizId ref="bizIdref" v-model="formData.code" id-code="route" :is-create="isCreate" />-->
            <el-input v-model="formData.code" :disabled="!isCreate" />
          </el-form-item>
          <el-form-item label="名称" prop="name">
            <el-input v-model="formData.name" />
          </el-form-item>
          <el-form-item label="权限等级" prop="roleCategoryCode">
            <el-select
              v-model="formData.roleCategoryCode"
              placeholder="请选择权限等级"
            >
              <el-option
                v-for="item in categories"
                :key="item.code"
                :value="item.code"
                :label="item.name"
              />
            </el-select>
          </el-form-item>
  <!--        <el-form-item label="描述" prop="description">-->
  <!--          <el-input v-model="formData.description" />-->
  <!--        </el-form-item>-->
        </el-form>
      </template>
      <template #permissions="scope">
        <PermissionTags :role="scope.row.code" :name="scope.row.name"/>
      </template>
    </CURDTable>
    <el-dialog
      v-model="permissionStatus.dialogVisible"
      title="编辑角色权限"
      append-to-body
      width="80%"
      @close="permissionStatus.done"
    >
      <CommonTable
        v-if="permissionStatus.dialogVisible"
        ref="permissionTableRef"
        class="h-[60vh]"
        :url="permissionUrl"
        pk="code"
        :columns="permissionColumns"
        :full-screen="false"
        :show-operations-column="false"
        @selection-change="handleSelectionChange"
        @loaded="onTableLoaded"
      >
        <template #extra-column>
          <el-table-column type="selection" width="55" />
        </template>
      </CommonTable>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import type { FormRules } from "element-plus";
import {roleColumns, roleUrl} from "@/types/tables/role.table";
import PermissionTags from "@/views/roles/components/PermissionTags.vue";
import {getRoleCategoriesApi, bindRolePermissionApi, unbindRolePermissionApi} from "@/api/system/role.api";
import {permissionColumns, permissionUrl} from "@/types/tables/permission.table";
import {Category, ElTagStyle, Tag} from "@/views/roles/types";

// const bizIdref = ref();
const tableRef = ref();
const permissionTableRef = ref();

const initialFormData = {
  code: "",
  name: "",
  roleCategoryCode: "",
  description: "",
};

const formRules = {
  code: [
    { required: true, message: '请填写编号', trigger: 'blur' },
  ],
  name: [
    { required: true, message: '请填写名称', trigger: 'blur' },
  ],
  roleCategoryCode: [
    { required: true, message: '请选择权限等级', trigger: 'change' },
  ],
} as FormRules;


const permissionStatus = reactive({
  role: "",
  done: () => {},
  dialogVisible: false,
  current: [] as Tag[],
  initialized: false,
});

const permissionModify = (role: string, done: () => void, current: Tag[]) => {
  permissionStatus.role = role;
  permissionStatus.done = done;
  permissionStatus.current = current;
  permissionStatus.dialogVisible = true;
}

provide("permissionModify", permissionModify)

const onTableLoaded = async () => {
  await nextTick();
  permissionTableRef.value.initialized = false;
  permissionTableRef.value.toggleSelection();
  permissionTableRef.value.toggleSelection(permissionStatus.current);
  permissionTableRef.value.initialized = true;
}

const handleSelectionChange = async (modified: Tag[]) => {
  // find diff
  if (!permissionTableRef.value.initialized) return;
  if (modified.length > permissionStatus.current.length) {
    const currentCodeSet = new Set(
      permissionStatus.current.map(item => item.code)
    );
    modified.filter(item => !currentCodeSet.has(item.code)).forEach((newPerm) => {
      bindRolePermissionApi(permissionStatus.role, newPerm.code).then(() => {
        ElMessage.success("权限绑定成功！")
      })
    });
  } else {
    const modifiedCodeSet = new Set(
      modified.map(item => item.code)
    );
    permissionStatus.current.filter(item => !modifiedCodeSet.has(item.code)).forEach((oldPerm) => {
      unbindRolePermissionApi(permissionStatus.role, oldPerm.code).then(() => {
        ElMessage.primary("权限解绑成功！")
      })
    });
  }
  permissionStatus.current = modified;
}


const categories = ref<Category[]>([]);

const elTagStyles = ["info", "primary", "warning", "danger"] as ElTagStyle[];

const categoryMap = (category: string) => {
  const id = categories.value.find(item => item.code === category)?.id
  return elTagStyles[(id || 1) - 1] || "info";
}

onMounted(() => {
  getRoleCategoriesApi().then((data) => {
    categories.value = data;
  })
})

provide("categoryMap", categoryMap);

</script>
