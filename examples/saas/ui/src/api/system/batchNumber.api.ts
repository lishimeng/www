import request from '@/utils/request'
import { BatchNumberDto, CreateBatchNumberReq } from '@/types/batch/batchNumber.type'

// 获取批次号列表
export function getBatchNumberList(params?: any) {
  return request({
    url: '/api/batch',
    method: 'get',
    params
  })
}

// 创建批次号
export function createBatchNumber(data: CreateBatchNumberReq) {
  return request({
    url: '/api/batch',
    method: 'post',
    data
  })
}

// 删除批次号
export function deleteBatchNumber(code: string) {
  return request({
    url: `/api/batch/${code}`,
    method: 'delete'
  })
}

// 预览批次号生成值
export function previewBatchNumber(data: any) {
  return request({
    url: '/api/batch/preview',
    method: 'post',
    data
  })
}

