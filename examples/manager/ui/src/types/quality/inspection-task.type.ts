// 质检任务类型定义
export interface InspectionTaskInfo {
  id?: number
  orderNumber?: string
  taskCode?: string // 任务编号（用于获取汇总数据）
  qualityCategory?: string
  inspectionType?: string // 检验类型
  inspectionPlan?: number
  inspectionPlanName?: string // 检验方案名称
  workOrder?: string
  workUnitCode?: string // 报工编号
  batchNo?: string // 批次号
  materialCode?: string // 物料编码
  materialName?: string // 物料名称
  process?: number
  role?: string
  taskStatus?: string
  priority?: string
  planStartTime?: string
  planEndTime?: string
  actualStartTime?: string
  actualEndTime?: string
  inspector?: string
  remark?: string
  enabled?: boolean
  // 汇总字段
  defectTotal?: number
  defectFatalCount?: number
  defectSeriousCount?: number
  defectMinorCount?: number
  qualifiedRate?: number
  totalInspectedQty?: number
  qualifiedQty?: number
  unqualifiedQty?: number
  detailCount?: number
  completedDetailCount?: number
  latestDetailUpdateTime?: string
  latestDetailOperateBy?: string
  createTime?: string
  updateTime?: string
}

// 任务汇总信息
export interface TaskSummaryInfo {
  defectTotal: number
  defectFatalCount: number
  defectSeriousCount: number
  defectMinorCount: number
  qualifiedRate?: number
  totalInspectedQty: number
  qualifiedQty: number
  unqualifiedQty: number
  detailCount: number
  completedDetailCount: number
  latestDetailUpdateTime: string
  latestDetailOperateBy: string
  latestDetail?: {
    id: number
    type: string
    documentNumber: string
    updateTime: string
    operator: string
    remark: string
  }
}

// 质检任务列表响应类型
export interface InspectionTaskListResponse {
  code: number
  message: string
  data: {
    items: InspectionTaskInfo[]
    total: number
    pageNum: number
    pageSize: number
  }
}

// 质检任务响应类型
export interface InspectionTaskResponse {
  code: number
  message: string
  task: InspectionTaskInfo
}

// 任务状态选项
export interface TaskStatusOption {
  value: string
  label: string
  type?: string
}

// 优先级选项
export interface PriorityOption {
  value: string
  label: string
  type?: string
}

// 任务分类选项
export interface TaskCategoryOption {
  value: string
  label: string
}
