import type { InspectionItemInfo } from './inspection-item.type'

export interface QualityPlanInfo {
  id?: number
  code?: string
  name?: string
  businessType?: string
  materialCode?: string
  materialName?: string
  inspectionType?: string
  samplingMethod?: string
  samplingValue?: number | null  // 支持null值，匹配后端*float64
  samplingUnit?: string
  inspectionItems?: number[]     // 统一为ID数组，匹配后端
  uniqueItems?: number[]        // 唯一质检项ID数组
  uniqueGroups?: number[]       // 唯一质检项组ID数组
  inspectionObjectType?: string  // 质检对象类型：单件/批量
  fullInspectionStrategy?: string // 全检策略：统一计算合格率/单件分别记录结果
  samplingThreshold?: number | null // 抽检判定标准（合格率阈值，如0.95表示95%）
  workstationId?: number        // 关联工位ID
  productionPlanId?: number     // 关联生产计划单ID
  remark?: string
  enabled?: boolean
  createTime?: string
  updateTime?: string
}