import {TableColumns} from "@/components/CommonTable/types";

export const orgUrl = "/api/users/saas/orgs";

export const orgColumns = [
  { prop: 'code', label: '编码', showOverflowTooltip: true },
  { prop: 'name', label: '名称', isName: true },
  { prop: 'admin', label: '超级管理员', showOverflowTooltip: true },
] as TableColumns[];

export const orgMembersUrl = "/api/users/saas/list"

export const orgMembersColumns = [
  { prop: 'code', label: '编码', width: '100', showOverflowTooltip: true },
  { prop: 'avatar', label: '头像', width: '80', component: "custom" },
  { prop: 'name', label: '名称', width: '120', isName: true },
  { prop: 'status', label: '状态', width: '100', component: 'enum_tag',
    enums: [
      { value: 10, label: '正常', type: "success" },
      { value: 20, label: '停用', type: "warning" },
    ]
  },
  { prop: 'roles', label: '用户角色', component: "custom" },
] as TableColumns[];
