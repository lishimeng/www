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
