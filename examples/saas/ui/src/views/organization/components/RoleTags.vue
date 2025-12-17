<template>
  <div class="flex gap-2 flex-wrap">
    <el-tag
      v-for="permission in roles"
      :key="permission.code"
      closable
      :type="permission.style"
      @close="delRole(permission)"
    >
      <ForeignKey
        :code="permission.code"
        url="/api/roles/role"
      />
    </el-tag>
    <el-button size="small" @click="showSelect">
      + 角色
    </el-button>
  </div>
</template>

<script setup lang="ts">
import {getUserRolesApi, unbindRolePermissionApi, unbindUserRoleApi} from "@/api/system/role.api";
import {withConfirm} from "@/composables/useConfirmDialog";
import {
  CategoryMapFunc,
  categoryMapFuncWrapper,
  Tag,
  TagsModifyFunc,
  tagsModifyFuncWrapper
} from "@/views/roles/types";
import {useFkStore} from "@/store";

const props = defineProps({
  user: {
    type: String,
    default: "",
  },
  name: {
    type: String,
    default: "",
  },
});

const roles = ref<Tag[]>([]);
const fkStore = useFkStore();

const injectErr = () => {
  console.log("inject not found")
  ElMessage.error("系统错误!")
}

const categoryMap = inject<CategoryMapFunc>("categoryMap", categoryMapFuncWrapper(injectErr));
const roleModify = inject<TagsModifyFunc>("roleModify", tagsModifyFuncWrapper(injectErr));

const showSelect = () => {
  roleModify(props.user, () => {
    fetchRoles(props.user)
  }, roles.value);
}

const delRole = async (role: Tag) => {
  const result = await withConfirm({
    message: `是否解除用户${props.name}的角色绑定：${fkStore.get("/api/roles/role", role.code).name}？`,
    title: "确认解绑角色",
  });
  if (result) {
    unbindUserRoleApi(props.user, role.code).then(() => {
      fetchRoles(props.user);
      ElMessage.success("解绑成功！")
    })
  }
}

const fetchRoles = (user:string) => {
  if (user) {
    getUserRolesApi(user).then((data: any[]) => {
      if (data) {
        roles.value = data.map(role => {
          // role.style = categoryMap(role.category);
          if (!role.isGlobal) {
            role.style = "info";
          } else {
            role.style = "primary";
          }
          role.code = role.roleCode;
          return role;
        })
      } else {
        roles.value = [];
      }
    })
  } else {
    roles.value = [];
  }
}

watch(
  () => props.user,
  (role) => fetchRoles(role),
  {immediate: true}
)

</script>
