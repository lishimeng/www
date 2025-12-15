// 客户信息类型定义
export interface CustomerInfo {
  id?: number
  code?: string
  name?: string
  shortName?: string
  contactInfo?: string
  address?: string
  socialCreditCode?: string
  remark?: string
  creator?: string
  createTime?: string
  updateTime?: string
}

// 客户列表响应类型
export interface CustomerListResponse {
  code: number
  message: string
  data: {
    items: CustomerInfo[]
    total: number
    pageNum: number
    pageSize: number
  }
}

// 客户详情响应类型
export interface CustomerResponse {
  code: number
  message: string
  customer: CustomerInfo
}
