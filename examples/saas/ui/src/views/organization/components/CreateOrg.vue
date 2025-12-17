<template>
  <el-dialog
    v-model="dialogVisible"
    title="新增组织"
    width="500px"
    @closed="onDialogClosed"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="120px"
    >
      <el-form-item label="管理员" prop="userCode">
<!--        <el-input v-model="formData.userCode" placeholder="请选择管理员用户" />-->
        <TableSelect
          v-model="formData.userCode"
          :url="securityUrl"
          placeholder="请选择管理员用户"
          title="选择组织管理员用户"
          :columns="securityColumns"
        />
      </el-form-item>
      <el-form-item label="组织编码" prop="orgCode">
        <el-input v-model="formData.orgCode" placeholder="请输入组织编码" />
      </el-form-item>
      <el-form-item label="组织名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入组织名称" />
      </el-form-item>
    </el-form>
    <template #footer>
        <span class="dialog-footer">
          <el-button class="w-18" :loading="isSubmitting" @click="dialogVisible = false">取消</el-button>
          <el-button class="w-18" type="primary" :loading="isSubmitting" @click="submit">提交</el-button>
        </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import type {FormRules} from "element-plus";
import {createOrgApi, type ReqSaasCreate} from "@/api/org.api";
import {securityColumns, securityUrl} from "@/types/tables/identity.table";

const formRef = ref();

const isSubmitting = ref(false);
const dialogVisible = ref(false);
const initialForm = {
  userCode: "",
  orgCode: "",
  name: "",
} as ReqSaasCreate;

const formData = ref<ReqSaasCreate>(initialForm);

const formRules = {
  userCode: [
    { required: true, message: '请选择管理员用户', trigger: 'blur' },
  ],
  orgCode: [
    { required: true, message: '请输入组织编码', trigger: 'blur' },
  ],
  name: [
    { required: true, message: '请输入组织名称', trigger: 'blur' },
  ],
} as FormRules;

const emit = defineEmits(["submit"]);

const open = () => {
  formData.value = initialForm;
  dialogVisible.value = true;
}

const submit = async () => {
  try {
    const valid = await formRef.value.validate();
    if (!valid) return;
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (validationError) {
    isSubmitting.value = false;
    return;
  }
  createOrgApi(formData.value)
    .then(() => {
      ElMessage.success("添加成功！");
      dialogVisible.value = false;
      emit("submit");
    });
}

const onDialogClosed = () => {
  isSubmitting.value = false;
}

defineExpose({
  open,
})

</script>
