export interface GroupItemInfo {
  relId: number
  itemId: number
  code: string
  name: string
  type: string
  sortNo: number
  isForce: boolean
}

export interface InspectionGroupInfo {
  id: number
  code: string  // 组编码
  groupName: string
  description?: string
  groupType?: string
  enabled: boolean  // 启用状态：true-启用，false-禁用
  creator?: string
  version?: string
  items?: GroupItemInfo[]  // 组内质检项列表（可选）
  createTime?: string
  updateTime?: string
}

export interface GroupItemInput {
  itemId: number
  sortNo: number
  isForce: boolean
}

export interface InspectionGroupCreateReq {
  id?: number
  code?: string  // 组编码（自动生成）
  groupName: string
  description?: string
  groupType?: string
  enabled?: boolean
  creator?: string
  items?: GroupItemInput[]
}

