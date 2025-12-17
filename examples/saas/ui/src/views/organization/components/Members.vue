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
      :show-operations-column="false"
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
      <template #operation="scope">
        <div class="flex">
          <el-button type="text" size="small" >停用</el-button>
          <el-button type="text" size="small" >停用</el-button>
        </div>
      </template>
    </CommonTable>

    <CreateMember
      ref="createRef"
      :organization="props.organization"
      @submit="tableRef.fetchPageData()"
    />

    <RoleDialog ref="roleDialogRef" />
  </div>
</template>

<script setup lang="ts">
import { orgMembersColumns, orgMembersUrl } from "@/types/tables/orgs.table";
import CreateMember from "@/views/organization/components/CreateMember.vue";
import {getOneOrgApi} from "@/api/org.api";
import {getRoleCategoriesApi} from "@/api/system/role.api";
import {Category, ElTagStyle, Tag} from "@/views/roles/types";
import RoleDialog from "@/views/organization/components/RoleDialog.vue";
import PermissionTags from "@/views/roles/components/PermissionTags.vue";
import RoleTags from "@/views/organization/components/RoleTags.vue";

const props = defineProps({
  organization: {
    type: String,
    required: true,
  }
});

const tableRef = ref();
const createRef = ref();
const roleDialogRef = ref();

const orgName = ref<string>("");

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
