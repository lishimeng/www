export interface ShipmentInspectionInfo {
  id?: number
  taskNumber?: string
  documentNumber?: string
  documentName?: string
  customerCode?: string
  customerName?: string
  batchNumber?: string
  productCode?: string
  productName?: string
  specification?: string
  unit?: string
  shipmentQuantity?: number
  inspectQuantity?: number
  defectQuantity?: number
  conclusion?: string
  result?: string
  shipmentDate?: string
  inspectDate?: string
  inspector?: string
  inspectionItems?: string
  remark?: string
  status?: number
  createTime?: string
  updateTime?: string
}