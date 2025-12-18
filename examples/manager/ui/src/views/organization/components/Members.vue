<template>
  <div>
    <CommonTable
      ref="tableRef"
      :url="orgMembersUrl"
      name="组织"
      pk="code"
      :columns="orgMembersColumns"
      :query-params="{
        enterprise: props.organization
      }"
      full-screen
    >
      <template #head>
        <div class="flex mb-3">
          <span class="text-xl font-bold">组织用户 - {{ orgName }}</span>
          <span class="flex-1" />
          <el-button type="primary" @click="createRef.open()">新增用户</el-button>
        </div>
      </template>
      <template #roles="scope">
        <RoleTags :user="scope.row.code" :name="scope.row.name"/>
      </template>
      <template #avatar="scope">
        <div class="relative">
          <ForeignKey :code="scope.row.code" field="avatar" url="/api/profile" >
            <template #default="{value}">
              <el-avatar :src="value" />
            </template>
          </ForeignKey>
          <el-button
            type="info"
            class="avatar-edit-btn"
            circle
            :icon="Camera"
            size="small"
            @click="triggerFileUpload(scope.row.code)"
          />
        </div>
      </template>
      <template #operation="scope">
        <div class="flex">
          <el-button
            type="text" size="small"
            @click="profileConfigRef.open(scope.row.code, scope.row.name)"
          >修改用户资料</el-button>
          <el-button
            type="text" size="small"
            @click="identitiesRef.open(scope.row.securityCode, scope.row.name)"
          >登录凭证</el-button>
          <el-button
            type="text" size="small"
            @click="resetPwdRef.open(scope.row.securityCode, scope.row.name)"
          >重置密码</el-button>
        </div>
      </template>
    </CommonTable>

    <input ref="fileInput" type="file" style="display: none" @change="handleFileChange" />

    <CreateMember
      ref="createRef"
      :organization="props.organization"
      @submit="tableRef.fetchPageData()"
    />

    <RoleDialog ref="roleDialogRef" />

    <AdminResetPwd ref="resetPwdRef" />

    <Identities ref="identitiesRef" />

    <ProfileConfig ref="profileConfigRef" @submit="tableRef.fetchPageData()" />

  </div>
</template>

<script setup lang="ts">
import { orgMembersColumns, orgMembersUrl } from "@/types/tables/orgs.table";
import CreateMember from "@/views/organization/components/CreateMember.vue";
import {getOneOrgApi} from "@/api/org.api";
import {getRoleCategoriesApi} from "@/api/system/role.api";
import {Category, ElTagStyle, Tag} from "@/views/roles/types";
import RoleDialog from "@/views/organization/components/RoleDialog.vue";
import RoleTags from "@/views/organization/components/RoleTags.vue";
import AdminResetPwd from "@/views/security/components/AdminResetPwd.vue";
import Identities from "@/views/security/components/Identities.vue";
import ProfileConfig from "@/views/organization/components/ProfileConfig.vue";
import { Camera } from "@element-plus/icons-vue";
import { useFkStore } from "@/store";
import FileAPI from "@/api/file.api";
import UserAPI from "@/api/system/user.api";

const props = defineProps({
  organization: {
    type: String,
    required: true,
  }
});

const tableRef = ref();
const createRef = ref();
const roleDialogRef = ref();
const resetPwdRef = ref();
const identitiesRef = ref();
const profileConfigRef = ref();
const fileInput = ref<HTMLInputElement | null>(null);
const fkStore = useFkStore();

const orgName = ref<string>("");
const user = ref<string>("");

watch(
  () => props.organization,
  async () => {
    if (!props.organization) return;
    const res = await getOneOrgApi(props.organization);
    orgName.value = res.name;
  }, { immediate: true }
)

const categories = ref<Category[]>([]);
const elTagStyles = ["info", "primary", "warning", "danger"] as ElTagStyle[];

const categoryMap = (category: string) => {
  const id = categories.value.find(item => item.code === category)?.id
  return elTagStyles[(id || 1) - 1] || "info";
}

const triggerFileUpload = (userCode: string) => {
  user.value = userCode;
  fileInput.value?.click();
};

const handleFileChange = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  const file = target.files ? target.files[0] : null;
  if (file) {
    // 调用文件上传API
    try {
      const data = await FileAPI.uploadFile(file);
      // 更新用户头像
      // userProfile.value.avatar = data.url;
      // 更新用户信息
      await UserAPI.updateProfile(user.value,{
        avatar: data.url,
      });
      fkStore.refresh("/api/profile", user.value);
      ElMessage.success("头像上传成功");
    } catch (error) {
      console.error("头像上传失败：" + error);
      ElMessage.error("头像上传失败");
    }
  }
};

onMounted(() => {
  getRoleCategoriesApi().then((data) => {
    categories.value = data;
  })
})

const roleModify = (user: string, done: () => void, current: Tag[]) => {
  roleDialogRef.value?.roleModify(user, done, current);
}

provide("categoryMap", categoryMap);
provide("roleModify", roleModify);

</script>

<style scoped lang="scss">
.avatar-edit-btn {
  position: absolute;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  border: none;
  transition: all 0.3s ease;

  &:hover {
    background: rgba(0, 0, 0, 0.7);
  }
}
</style>
