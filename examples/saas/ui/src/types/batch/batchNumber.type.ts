export interface BatchNumberAttributeDto {
  paramName: string  // 属性字段名
  paramLabel: string  // 属性名称
  value: string      // 属性值
}

export interface BatchNumberDto {
  id: number
  code: string                    // 批次号
  ruleCode: string                // 规则编码
  ruleDesc: string                 // 规则描述
  desc: string                     // 记录描述
  attributes: BatchNumberAttributeDto[]  // 属性值列表
  createTime: string
  updateTime: string
  status: number
}

export interface CreateBatchNumberReq {
  ruleCode: string                 // 规则编码
  desc: string                     // 记录描述
  attributes: BatchNumberAttributeDto[]  // 属性值列表
}

