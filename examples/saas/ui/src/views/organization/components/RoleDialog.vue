<template>
  <el-dialog
    v-model="status.dialogVisible"
    title="配置用户角色"
    append-to-body
    width="80%"
    @close="status.done"
  >
    <CommonTable
      v-if="status.dialogVisible"
      ref="tableRef"
      class="h-[60vh]"
      :url="roleUrl"
      pk="code"
      :columns="roleColumns"
      :full-screen="false"
      :show-operations-column="false"
      @selection-change="handleSelectionChange"
      @loaded="onTableLoaded"
    >
      <template #extra-column>
        <el-table-column type="selection" width="55" />
      </template>
      <template #permissions="scope">
        <PermissionTags :role="scope.row.code" :name="scope.row.name" readonly/>
      </template>
    </CommonTable>
  </el-dialog>
</template>

<script setup lang="ts">
import {Tag} from "@/views/roles/types";
import {
  bindUserRoleApi,
  unbindUserRoleApi
} from "@/api/system/role.api";
import {roleColumns, roleUrl} from "@/types/tables/role.table";
import PermissionTags from "@/views/roles/components/PermissionTags.vue";

const tableRef = ref();

const status = reactive({
  user: "",
  done: () => {},
  dialogVisible: false,
  current: [] as Tag[],
  initialized: false,
});

const roleModify = (user: string, done: () => void, current: Tag[]) => {
  status.user = user;
  status.done = done;
  status.current = current;
  status.dialogVisible = true;
}

const onTableLoaded = async () => {
  await nextTick();
  tableRef.value.initialized = false;
  tableRef.value.toggleSelection();
  tableRef.value.toggleSelection(status.current);
  tableRef.value.initialized = true;
}

const handleSelectionChange = async (modified: Tag[]) => {
  // find diff
  if (!tableRef.value.initialized) return;
  if (modified.length > status.current.length) {
    const currentCodeSet = new Set(
      status.current.map(item => item.code)
    );
    modified.filter(item => !currentCodeSet.has(item.code)).forEach((newRole) => {
      bindUserRoleApi(status.user, newRole.code).then(() => {
        ElMessage.success("角色绑定成功！")
      })
    });
  } else {
    const modifiedCodeSet = new Set(
      modified.map(item => item.code)
    );
    status.current.filter(item => !modifiedCodeSet.has(item.code)).forEach((oldRole) => {
      unbindUserRoleApi(status.user, oldRole.code).then(() => {
        ElMessage.primary("角色解绑成功！")
      })
    });
  }
  status.current = modified;
}

defineExpose({
  roleModify,
})

</script>
