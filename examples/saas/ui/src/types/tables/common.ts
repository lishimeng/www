import type { TableColumns } from "@/components/CommonTable/types";

export const createTime = { prop: 'createTime', label: '创建时间', component: 'time', width: '180' } as TableColumns;
export const updateTime = { prop: 'updateTime', label: '更新时间', component: 'time', width: '180' } as TableColumns;
export const tableChangeInfoCols = [createTime, updateTime] as TableColumns[];

export const createBy = { prop: 'createBy', label: '创建人', component: 'text', width: '140' } as TableColumns;
export const modifyBy = { prop: 'modifyBy', label: '更新人', component: 'text', width: '140' } as TableColumns;
export const operatorCols = [createBy, modifyBy] as TableColumns[];
