<template>
  <el-dialog
    v-model="dialogVisible"
    title="新增组织用户"
    width="500px"
    @closed="onDialogClosed"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="120px"
    >
      <el-form-item label="登录凭证" prop="identityCode">
        <el-input v-model="formData.identityCode" placeholder="请输入登录凭证" />
      </el-form-item>
      <el-form-item label="登录凭证类型" prop="identityType">
        <el-select v-model="formData.identityType" clearable placeholder="留空以自动判断（手机号/邮箱/微信Id）">
          <el-option v-for="item in identityTypeOptions" :key="item.value" :value="item.value" :label="item.label"/>
        </el-select>
      </el-form-item>
      <el-form-item label="用户名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入用户名称" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="formData.password" placeholder="请输入密码" show-password />
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
import {createOrgApi, joinOrgApi, type ReqSaasCreate} from "@/api/org.api";
import {securityColumns, securityUrl} from "@/types/tables/identity.table";
import UserAPI from "@/api/system/user.api";

const formRef = ref();

const props = defineProps({
  organization: {
    type: String,
    required: true,
  }
});

interface UserCreate {
  identityCode: string,
  identityType: number | null,
  password: string,
  name: string,
}

const isSubmitting = ref(false);
const dialogVisible = ref(false);
const initialForm = {
  identityCode: "",
  identityType: null,
  password: "",
  name: "",
} as UserCreate;

const formData = ref<UserCreate>(initialForm);

const formRules = {
  identityCode: [
    { required: true, message: '请输入登录凭证', trigger: 'blur' },
  ],
  name: [
    { required: true, message: '请输入用户名称', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
  ],
} as FormRules;

const identityTypeOptions = [
  { value: 1, label: '登录码' },
  { value: 2, label: '手机号' },
  { value: 3, label: '邮箱' },
  { value: 4, label: '微信UnionId' },
  { value: 5, label: '微信OpenId' },
];

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

  try {
    const security = await UserAPI.register(
      formData.value.identityCode, formData.value.identityType,
      formData.value.password, 20);
    await joinOrgApi(security.code, props.organization, formData.value.name );
  } catch (e) {
    ElMessage.error("添加失败！");
    return;
  }
  ElMessage.success("添加成功！");
  dialogVisible.value = false;
  emit("submit");
}

const onDialogClosed = () => {
  isSubmitting.value = false;
}

defineExpose({
  open,
})

</script>
