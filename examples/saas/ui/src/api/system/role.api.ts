import httpRequest, { get, del } from "@/utils/request";

const ROLE_PERM_BASE_URL = "/api/permissions/role-permission";

export const getRolePermissionsApi = (roleId: string) => get(`${ROLE_PERM_BASE_URL}/${roleId}`);

export const bindRolePermissionApi = (roleCode: string, permissionCode: string) => {
  return httpRequest.request<any, any>({
    url:ROLE_PERM_BASE_URL,
    method: 'post',
    data: {
      roleCode,
      permissionCode
    },
  });
}

export const unbindRolePermissionApi = (roleCode: string, permissionCode: string) => {
  return httpRequest.request<any, any>({
    url:ROLE_PERM_BASE_URL,
    method: 'delete',
    data: {
      roleCode,
      permissionCode
    },
  });
}

export const getRoleCategoriesApi = () => get("/api/roles/category");


const USER_ROLE_BASE_URL = "/api/roles/bind";

export const getUserRolesApi = (saasUser: string) => get(`${USER_ROLE_BASE_URL}/user/${saasUser}`);

export const getRoleUsersApi = (roleCode: string) => get(`${USER_ROLE_BASE_URL}/${roleCode}`);

export const getRoleOneApi = (roleCode: string) => get(`/api/roles/role/${roleCode}`);

export const bindUserRoleApi = (saasUserCode: string, roleCode: string) => {
  return httpRequest.request<any, any>({
    url:USER_ROLE_BASE_URL,
    method: 'post',
    data: {
      saasUserCode,
      roleCode,
    },
  });
}

export const unbindUserRoleApi = (saasUserCode: string, roleCode: string) => {
  return httpRequest.request<any, any>({
    url:USER_ROLE_BASE_URL,
    method: 'delete',
    data: {
      saasUserCode,
      roleCode,
    },
  });
}
