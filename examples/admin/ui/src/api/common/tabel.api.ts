import { del, get, post, put, patch } from "@/utils/request";

export const fetchListApi = (
  baseUrl: string,
  p: object,
) => get(baseUrl, p);

export const fetchOneApi = (
  baseUrl: string,
  pk: any,
) => get(baseUrl + `/${pk}`);

export const createApi = (
  baseUrl: string,
  p: object,
) => post(baseUrl, p);

export const updateApi = (
  baseUrl: string,
  pk: any,
  p: object,
) => {
  if (pk) {
    return put(baseUrl + `/${pk}`, p);
  } else {
    return put(baseUrl, p);
  }
}

export const patchUpdateApi = (
  baseUrl: string,
  pk: any,
  p: object,
) => {
  if (pk) {
    return patch(baseUrl + `/${pk}`, p);
  } else {
    return patch(baseUrl, p);
  }
}

export const deleteApi = (
  baseUrl: string,
  pk: any,
) => del(baseUrl + `/${pk}`, {});
