<template>
  <CURDTable
    ref="tableRef"
    :url="permissionUrl"
    name="权限"
    pk="code"
    :columns="permissionColumns"
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
        <el-form-item label="权限等级" prop="category">
          <UrlSelect
            v-model="formData.category"
            url="/api/roles/category"
            placeholder="请选择权限等级"
          />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="formData.description" />
        </el-form-item>
      </el-form>
    </template>
  </CURDTable>
</template>

<script setup lang="ts">
import {permissionColumns, permissionUrl} from "@/types/tables/permission.table";
import type { FormRules } from "element-plus";

// const bizIdref = ref();
const tableRef = ref();

const initialFormData = {
  code: "",
  name: "",
  category: "",
  description: "",
};

const formRules = reactive<FormRules>({
  code: [
    { required: true, message: '请填写编号', trigger: 'blur' },
  ],
  name: [
    { required: true, message: '请填写名称', trigger: 'blur' },
  ],
  category: [
    { required: true, message: '请选择权限等级', trigger: 'change' },
  ],
});

// const onDialogOpen = async (data: Record<string, any>, isCreate: boolean) => {
//   await nextTick();
//   bizIdref.value?.refresh(data, isCreate);
// };
</script>
