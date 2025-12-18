<template>
  <el-dialog
    v-model="dialogVisible"
    :title="`用户资料：${userName}`"
    width="500px"
    @closed="onDialogClosed"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="120px"
    >
      <el-form-item label="名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入用户名称" show-name />
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

const formRef = ref();
const user = ref<string>("");
const userName = ref<string>("");

interface UserProfile {
  name: string,
}

const isSubmitting = ref(false);
const dialogVisible = ref(false);
const initialForm = {
  name: "",
} as UserProfile;

const formData = ref<UserProfile>(initialForm);

const formRules = {
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' },
  ],
} as FormRules;

const emit = defineEmits(["submit"]);

const open = (userCode: string, name: string) => {
  user.value = userCode;
  userName.value = name;
  formData.value.name = name;
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
    await UserAPI.updateProfile(user.value, formData.value);
  } catch (e) {
    ElMessage.error("更新失败！");
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
