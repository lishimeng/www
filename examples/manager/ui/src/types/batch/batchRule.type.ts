export interface ParamDto {
  name: string
  paramType: number  // 1-序列号，2-时间，3-业务参数，4-固定字符
  config: string
  length: number
}

export interface BatchRuleDto {
  code: string
  desc: string
  template: string
  params: ParamDto[]
  createTime: string
  updateTime: string
}

