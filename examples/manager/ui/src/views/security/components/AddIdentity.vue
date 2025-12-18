<template>
  <el-dialog
    v-model="dialogVisible"
    :title="title"
    width="500px"
    @closed="onDialogClosed"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="120px"
    >
      <el-form-item label="账号" prop="securityCode">
        <el-input v-model="formData.securityCode" :readonly="securityCodeReadOnly" />
      </el-form-item>
      <el-form-item label="登录凭证" prop="code">
        <el-input v-model="formData.code" />
      </el-form-item>
      <el-form-item label="登录凭证类型" prop="identityType">
        <el-select v-model="formData.identityType" clearable placeholder="留空以自动判断（手机号/邮箱/微信Id）">
          <el-option v-for="item in identityTypeOptions" :key="item.value" :value="item.value" :label="item.label"/>
        </el-select>
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
import UserAPI from "@/api/system/user.api";
import {IdentityForm} from "@/views/security/types";

const formRef = ref();
const userName = ref<string>("");

const isSubmitting = ref(false);
const dialogVisible = ref(false);
const securityCodeReadOnly = ref(false);
const initialForm = {
  code: "",
  securityCode: "",
  identityType: null,
} as IdentityForm;

const formData = ref<IdentityForm>(initialForm);

const formRules = {
  code: [
    { required: true, message: '请输入登录凭证', trigger: 'blur' },
  ],
  identityCode: [
    { required: true, message: '请输入登录凭证', trigger: 'blur' },
  ],
  platform: [
    { required: true, message: '请选择用户平台', trigger: 'blur' },
  ],
  securityCode: [
    { required: true, message: '请输入用户码', trigger: 'blur' },
  ],
} as FormRules;

const emit = defineEmits(["submit"]);

const open = (securityCode?: string, name?: string) => {
  formData.value.code = "";
  if (securityCode) {
    formData.value.securityCode = securityCode;
    securityCodeReadOnly.value = true;
  } else {
    formData.value.securityCode = "";
    securityCodeReadOnly.value = false;
  }
  formData.value.identityType = null;
  dialogVisible.value = true;
  if (name) userName.value = name;
  formData.value = initialForm;
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
  UserAPI.addIdentity(formData.value.code, formData.value.securityCode, formData.value.identityType)
    .then(() => {
      ElMessage.success("添加成功！");
      dialogVisible.value = false;
    });
  emit("submit");
}

const title = computed(() => {
  if (userName.value) {
    return `登录凭证：${userName.value}`;
  } else {
    return "登录凭证";
  }
})

const identityTypeOptions = [
  { value: 1, label: '登录码' },
  { value: 2, label: '手机号' },
  { value: 3, label: '邮箱' },
  { value: 4, label: '微信UnionId' },
  { value: 5, label: '微信OpenId' },
];

const onDialogClosed = () => {
  isSubmitting.value = false;
}

defineExpose({
  open,
})

</script>
