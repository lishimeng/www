import { post } from "@/utils/request";

export const createBizId = (
  code: string,
  data: Record<string, any>
) => post(`/api/bizId/generate/${code}`, data);
