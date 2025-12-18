<template>
  <el-dialog
    v-model="dialogVisible"
    :title="title"
    width="80%"
  >
    <CommonTable
      ref="tableRef"
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
          <el-button type="primary" @click="addIdentityRef.open(currentSecurity, userName)">添加凭证</el-button>
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

  <AddIdentity ref="addIdentityRef" @submit="tableRef.fetchPageData()" />
</template>

<script setup lang="ts">
import {identityColumns} from "@/types/tables/identity.table";
import AddIdentity from "@/views/security/components/AddIdentity.vue";
import {withConfirm} from "@/composables/useConfirmDialog";
import {IdentityForm} from "@/views/security/types";
import UserAPI from "@/api/system/user.api";

const tableRef = ref();
const addIdentityRef = ref();
const currentSecurity = ref<string>("");
const userName = ref<string>("");

const dialogVisible = ref(false);

const open = (userCode: string, name?: string) => {
  currentSecurity.value = userCode;
  if (name) userName.value = name;
  dialogVisible.value = true;
}

const delIdentity = async (row: IdentityForm) => {
  const result = await withConfirm({
    message: "删除的数据不可恢复，是否继续？",
    title: "确认删除登录凭证",
  });
  if (!result) return;
  UserAPI.delIdentity(row.code, row.securityCode, row.identityType)
    .then(() => {
      ElMessage.success("删除成功！");
      tableRef.value?.fetchPageData();
    });
}

const title = computed(() => {
  if (userName.value) {
    return `登录凭证：${userName.value}`;
  } else {
    return "登录凭证";
  }
})

defineExpose({
  open,
})

</script>
