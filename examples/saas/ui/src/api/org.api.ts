import { get, post } from "@/utils/request";

const ORG_BASE_URL = "/api/users/saas";

export interface ReqSaasCreate {
  userCode: string;
  orgCode: string;
  name: string;
}

export const createOrgApi = (data: ReqSaasCreate) => post(`${ORG_BASE_URL}/create`, data);

export const getOneOrgApi = (orgCode: string) => get(`${ORG_BASE_URL}/orgs/${orgCode}`);

export const joinOrgApi = (userCode: string, orgCode: string, userName: string ) => post(
  `${ORG_BASE_URL}/join`, {userCode, orgCode, userName}
);
