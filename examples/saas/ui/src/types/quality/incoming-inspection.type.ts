export interface IncomingInspectionInfo {
  id?: number
  taskNumber?: string      // 质检任务单号
  documentNumber?: string  // 检验单编号
  materialCode?: string    // 物料编码
  materialName?: string    // 物料名称
  supplierCode?: string    // 供应商编码
  supplierName?: string    // 供应商名称
  incomingDate?: string    // 来料日期
  inspectDate?: string     // 检测日期
  inspector?: string       // 检测人员
  inspectionStatus?: string  // 单据状态
  inspectionItems?: string   // 检验项目结果（JSON）
  conclusion?: string        // 检验结论
  remark?: string            // 备注
  createTime?: string
  updateTime?: string
}