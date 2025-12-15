import {TableColumns} from "@/components/CommonTable/types";

export const permissionUrl = "/api/permissions/permission"

export const permissionColumns = [
  { prop: 'code', label: '编码', width: '150', showOverflowTooltip: true },
  { prop: 'name', label: '名称', width: '150', isName: true },
  { prop: 'category', label: '权限等级', width: '120',showOverflowTooltip: true },
  { prop: 'description', label: '描述', showOverflowTooltip: true },
] as TableColumns[];
