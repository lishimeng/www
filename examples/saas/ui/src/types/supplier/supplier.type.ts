// 供应商信息类型定义
export interface SupplierInfo {
  id?: number
  code?: string
  name?: string
  shortName?: string
  isExempt?: boolean
  contactInfo?: string
  address?: string
  socialCreditCode?: string
  remark?: string
  creator?: string
  createTime?: string
  updateTime?: string
}

// 供应商列表响应类型
export interface SupplierListResponse {
  code: number
  message: string
  data: {
    items: SupplierInfo[]
    total: number
    pageNum: number
    pageSize: number
  }
}

// 供应商详情响应类型
export interface SupplierResponse {
  code: number
  message: string
  supplier: SupplierInfo
}
