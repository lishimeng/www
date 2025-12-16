import {TableColumns} from "@/components/CommonTable/types";

export const roleUrl = "/api/roles/role"

export const roleColumns = [
  { prop: 'code', label: '编码', width: '150', showOverflowTooltip: true },
  { prop: 'name', label: '名称', width: '150', isName: true },
  { prop: 'roleCategoryCode', label: '权限等级', width: '120',showOverflowTooltip: true },
  // { prop: 'description', label: '描述', showOverflowTooltip: true },
  { prop: 'permissions', label: '权限列表', component: "custom" },
] as TableColumns[];
