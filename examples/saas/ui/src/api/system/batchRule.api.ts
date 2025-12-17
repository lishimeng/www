import request from "@/utils/request";
import { BatchRuleDto } from "@/types/batch/batchRule.type";

// 获取批次规则列表
export function getBatchRuleList(params?: any) {
  return request({
    url: '/api/batch/rule',
    method: 'get',
    params
  })
}

// 创建批次规则
export function createBatchRule(data: BatchRuleDto) {
  return request({
    url: '/api/batch/rule',
    method: 'post',
    data
  })
}

// 更新批次规则
export function updateBatchRule(code: string, data: BatchRuleDto) {
  return request({
    url: `/api/batch/rule/${code}`,
    method: 'put',
    data
  })
}

// 删除批次规则
export function deleteBatchRule(code: string) {
  return request({
    url: `/api/batch/rule/${code}`,
    method: 'delete'
  })
}

// 生成批次号（测试用）
export const createBatchNumber = (
  ruleCode: string,
  data: Record<string, any>
) => request({
  url: `/api/batch/rule/generate/${ruleCode}`,
  method: 'post',
  data
});

