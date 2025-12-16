<template>
  <div class="flex gap-2 flex-wrap">
    <el-tag
      v-for="permission in permissions"
      :key="permission.code"
      closable
      :type="permission.style"
      @close="delPermission(permission)"
    >
      {{ permission.name }}
    </el-tag>
    <el-button size="small" @click="showSelect">
      + 权限
    </el-button>
  </div>
</template>

<script setup lang="ts">
import {getRolePermissionsApi, unbindRolePermissionApi} from "@/api/system/role.api";
import {withConfirm} from "@/composables/useConfirmDialog";
import {
  CategoryMapFunc,
  categoryMapFuncWrapper,
  Permission,
  PermissionModifyFunc,
  permissionModifyFuncWrapper
} from "@/views/roles/types";

const props = defineProps({
  role: {
    type: String,
    default: "",
  },
  name: {
    type: String,
    default: "",
  },
});

const permissions = ref<Permission[]>([]);

const injectErr = () => {
  console.log("inject not found")
  ElMessage.error("系统错误!")
}

const categoryMap = inject<CategoryMapFunc>("categoryMap", categoryMapFuncWrapper(injectErr))
const permissionModify = inject<PermissionModifyFunc>("permissionModify", permissionModifyFuncWrapper(injectErr));

const showSelect = () => {
  permissionModify(props.role, () => {
    fetchPermissions(props.role)
  }, permissions.value);
}

const delPermission = async (permission: Permission) => {
  const result = await withConfirm({
    message: `是否删除角色${props.name}的绑定权限：${permission.name}？`,
    title: "确认删除权限",
  });
  if (result) {
    unbindRolePermissionApi(props.role, permission.code).then(() => {
      fetchPermissions(props.role);
      ElMessage.success("删除成功！")
    })
  }
}

const fetchPermissions = (role:string) => {
  if (role) {
    getRolePermissionsApi(role).then((data: Permission[]) => {
      if (data) {
        permissions.value = data.map(permission => {
          permission.style = categoryMap(permission.category)
          return permission
        })
      } else {
        permissions.value = []
      }
    })
  } else {
    permissions.value = []
  }
}

watch(
  () => props.role,
  (role) => fetchPermissions(role),
  {immediate: true}
)

</script>
