// 排班信息接口
export interface ScheduleInfo {
  id?: number;
  code?: string;
  name: string;
  startTime: string;
  endTime: string;
  workDay: string;
  description?: string;
  creator?: string;
  createTime?: string;
  updateTime?: string;
}

// 假期信息接口
export interface HolidayInfo {
  id?: number;
  code?: string;
  date: string;
  type: string; // holiday: 假期, workday: 工作日
  name: string;
  description?: string;
  creator?: string;
  createTime?: string;
  updateTime?: string;
}
