<template>
  <div>
    <CommonTable
      ref="tableRef"
      :url="securityUrl"
      name="登录凭证"
      pk="code"
      :columns="securityColumns"
      full-screen
    >
      <template #head>
        <div class="flex mb-3">
          <el-button type="primary" @click="registerUser">新增账号</el-button>
        </div>
      </template>
      <template #identities="scope">
        <el-button
          type="text" size="small"
          @click="openIdentities(scope.row.code)"
        >展开</el-button>
      </template>
      <template #operation="scope">
        <div class="flex">
<!--          <el-button-->
<!--            type="text" size="small"-->
<!--            :disabled="scope.row.platform !== 20"-->
<!--            @click="addIdentity(scope.row.code)"-->
<!--          >加入组织</el-button>-->
          <el-button
            type="text" size="small"
            @click="resetPwdRef.open(scope.row.code, scope.row.code)"
          >重置密码</el-button>
        </div>
      </template>
    </CommonTable>

    <el-dialog
      v-model="identityTableVisible"
      title="登录凭证"
      width="80%"
    >
      <CommonTable
        ref="identityTableRef"
        :url="`/api/users/identities/${currentSecurity}`"
        name="登录凭证"
        pk="code"
        :columns="identityColumns"
        :full-screen="false"
        no-border
        auto-index
      >
        <template #head>
          <div class="flex mb-3">
            <el-button type="primary" @click="addIdentity(currentSecurity)">添加凭证</el-button>
          </div>
        </template>
        <template #operation="scope">
          <el-button
            type="text" size="small"
            @click="delIdentity(scope.row as IdentityForm)"
          >删除</el-button>
        </template>
      </CommonTable>
    </el-dialog>

    <el-dialog
      v-model="identityDialogVisible"
      title="新增登录凭证"
      width="500px"
      @closed="onDialogClosed"
    >
      <el-form
          ref="identityFormRef"
          :model="identityForm"
          :rules="formRules"
          label-width="120px"
      >
        <el-form-item label="账号" prop="securityCode">
          <el-input v-model="identityForm.securityCode" :readonly="securityCodeReadOnly" />
        </el-form-item>
        <el-form-item label="登录凭证" prop="code">
          <el-input v-model="identityForm.code" />
        </el-form-item>
        <el-form-item label="登录凭证类型" prop="identityType">
          <el-select v-model="identityForm.identityType" clearable placeholder="留空以自动判断（手机号/邮箱/微信Id）">
            <el-option v-for="item in identityTypeOptions" :key="item.value" :value="item.value" :label="item.label"/>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button class="w-18" :loading="isSubmitting" @click="identityDialogVisible = false">取消</el-button>
          <el-button class="w-18" type="primary" :loading="isSubmitting" @click="submitIdentity">提交</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="registerDialogVisible"
      title="新增账号"
      width="500px"
      @closed="onDialogClosed"
    >
      <el-form
        ref="registerFormRef"
        :model="registerForm"
        :rules="formRules"
        label-width="120px"
      >
        <el-form-item label="登录凭证" prop="identityCode">
          <el-input v-model="registerForm.identityCode" />
        </el-form-item>
        <el-form-item label="登录凭证类型" prop="identityType">
          <el-select v-model="registerForm.identityType" clearable placeholder="留空以自动判断（手机号/邮箱/微信Id）">
            <el-option v-for="item in platformOptions" :key="item.value" :value="item.value" :label="item.label"/>
          </el-select>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="registerForm.password" show-password />
        </el-form-item>
        <el-form-item label="账号平台" prop="platform">
          <el-select v-model="registerForm.platform" clearable placeholder="请选择">
            <el-option v-for="item in platformOptions" :key="item.value" :value="item.value" :label="item.label"/>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button class="w-18" :loading="isSubmitting" @click="registerDialogVisible = false">取消</el-button>
          <el-button class="w-18" type="primary" :loading="isSubmitting" @click="submitRegister">提交</el-button>
        </span>
      </template>
    </el-dialog>

    <AdminResetPwd ref="resetPwdRef" />

  </div>
</template>

<script setup lang="ts">
import type { FormRules } from "element-plus";
import {identityColumns, securityColumns, securityUrl} from "@/types/tables/identity.table";
import UserAPI from "@/api/system/user.api";
import {withConfirm} from "@/composables/useConfirmDialog";
import AdminResetPwd from "@/views/security/components/AdminResetPwd.vue";
import {IdentityForm} from "@/views/security/types";

const tableRef = ref();
const identityTableRef = ref();
const identityFormRef = ref();
const registerFormRef = ref();
const resetPwdRef = ref();

const identityForm = reactive<IdentityForm>({
  code: "",
  securityCode: "",
  identityType: null,
});

interface RegisterForm {
  identityCode: string,
  identityType: number | null,
  password: string,
  platform: number,
}

const registerForm = reactive<RegisterForm>({
  identityCode: "",
  identityType: null,
  password: "",
  platform: 10,
});

const identityDialogVisible = ref(false);
const registerDialogVisible = ref(false);
const orgDialogVisible = ref(false);
const isSubmitting = ref(false);
const securityCodeReadOnly = ref(false);
const identityTableVisible = ref(false);
const currentSecurity = ref("");

const identityTypeOptions = [
  { value: 1, label: '登录码' },
  { value: 2, label: '手机号' },
  { value: 3, label: '邮箱' },
  { value: 4, label: '微信UnionId' },
  { value: 5, label: '微信OpenId' },
];

const platformOptions = [
  { value: 10, label: 'app' },
  { value: 20, label: 'saas' },
  { value: 30, label: 'admin' },
  { value: 40, label: 'manager' },
];

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

const delIdentity = async (row: IdentityForm) => {
  const result = await withConfirm({
    message: "删除的数据不可恢复，是否继续？",
    title: "确认删除登录凭证",
  });
  if (!result) return;
  UserAPI.delIdentity(row.code, row.securityCode, row.identityType)
      .then(() => {
        ElMessage.success("删除成功！");
        identityTableRef.value?.fetchPageData();
      });
}

function onDialogClosed() {
  isSubmitting.value = false;
}

const addIdentity = (securityCode?: string) => {
  identityForm.code = "";
  if (securityCode) {
    identityForm.securityCode = securityCode;
    securityCodeReadOnly.value = true;
  } else {
    identityForm.securityCode = "";
    securityCodeReadOnly.value = false;
  }
  identityForm.identityType = null;
  identityDialogVisible.value = true;
}

const submitIdentity = async () => {
  try {
    const valid = await identityFormRef.value.validate();
    if (!valid) return;
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (validationError) {
    isSubmitting.value = false;
    return;
  }
  UserAPI.addIdentity(identityForm.code, identityForm.securityCode, identityForm.identityType)
      .then(() => {
        ElMessage.success("添加成功！");
        identityDialogVisible.value = false;
        identityTableRef.value?.fetchPageData();
      });
}

const registerUser = () => {
  registerForm.identityCode = "";
  registerForm.identityType = null;
  registerForm.password = "";
  registerForm.platform = 10;
  registerDialogVisible.value = true;
}

const submitRegister = async () => {
  try {
    const valid = await registerFormRef.value.validate();
    if (!valid) return;
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (validationError) {
    isSubmitting.value = false;
    return;
  }
  UserAPI.register(registerForm.identityCode, registerForm.identityType, registerForm.password, registerForm.platform)
      .then(() => {
        ElMessage.success("注册成功！");
        registerDialogVisible.value = false;
        tableRef.value?.fetchPageData();
      });
}

const openIdentities = async (security: string) => {
  currentSecurity.value = security;
  identityTableVisible.value = true;
}

</script>
