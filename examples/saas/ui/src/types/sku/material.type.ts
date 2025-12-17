export interface MaterialInfo {
  id?: number
  code?: string
  name?: string
  spec?: string              // 规格型号（字符串，用于显示）
  specValue?: number | null  // 规格型号数值（用于计算）
  unit?: string
  warningVal?: number
  groupName?: string
  groupType?: string
  supplierCode?: string
  supplierName?: string
  route?: string
  line?: string
  process?: string
  description?: string
  creator?: string
  createTime?: string
  updateTime?: string
}