<template>
  <el-dialog
    v-model="dialogVisible"
    :title="`重置密码：${userName}`"
    width="500px"
    @closed="onDialogClosed"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="120px"
    >
      <el-form-item label="新密码" prop="password">
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
const user = ref<string>("");
const userName = ref<string>("");

interface ResetPwd {
  password: string,
}

const isSubmitting = ref(false);
const dialogVisible = ref(false);
const initialForm = {
  password: "",
} as ResetPwd;

const formData = ref<ResetPwd>(initialForm);

const formRules = {
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
  ],
} as FormRules;

const emit = defineEmits(["submit"]);

const open = (userCode: string, name: string) => {
  user.value = userCode;
  userName.value = name;
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
    await UserAPI.resetPassword(user.value, formData.value.password);
  } catch (e) {
    ElMessage.error("重置密码失败！");
    return;
  }
  ElMessage.success("重置密码成功！");
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
