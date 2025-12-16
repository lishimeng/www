import {TableColumns} from "@/components/CommonTable/types";
import {tableChangeInfoCols} from "@/types/tables/common";

export const identityUrl = "/api/identities"

export const identityColumns = [
  { prop: 'securityCode', label: '账号', showOverflowTooltip: true },
  { prop: 'code', label: '登录标识', showOverflowTooltip: true, isName: true },
  { prop: 'identityType', label: '登录标识类型', width: '120', component: 'enum_tag',
    enums: [
      { value: 0, label: '未知', type: "info" },
      { value: 1, label: '登录码', type: "danger" },
      { value: 2, label: '手机号', type: "success" },
      { value: 3, label: '邮箱', type: "warning" },
      { value: 4, label: '微信UnionId', type: "primary" },
      { value: 5, label: '微信OpenId', type: "primary" },
    ]
  },
  // { prop: 'status', label: '状态', width: '120', showOverflowTooltip: true },
  ...tableChangeInfoCols,
] as TableColumns[];

export const securityUrl = "/api/users"

export const securityColumns = [
  { prop: 'code', label: '账号', showOverflowTooltip: true },
  { prop: 'platform', label: '平台', width: '150', component: 'enum_tag',
    enums: [
      { value: 10, label: 'app', type: "info" },
      { value: 20, label: 'saas', type: "primary" },
      { value: 30, label: 'admin', type: "warning" },
      { value: 40, label: 'manager', type: "danger" },
    ]
  },
  { prop: 'identities', label: '登录凭证', width: '100', component: 'custom' },
  // { prop: 'status', label: '状态', width: '120', showOverflowTooltip: true },
  ...tableChangeInfoCols,
] as TableColumns[];
