import request from "@/utils/request";

const DICT_BASE_URL = "/api/dict";

const DictAPI = {
  //---------------------------------------------------
  // 字典相关接口
  //---------------------------------------------------

  /**
   * 获取字典根节点列表
   *
   * @param category 字典分类
   * @param name 字典名称（模糊查询）
   * @returns 字典根节点列表
   */
  getRoots(category?: string, name?: string) {
    return request<any, DictListResponse>({
      url: `${DICT_BASE_URL}/roots`,
      method: "get",
      params: { category, name },
    });
  },

  /**
   * 获取字典列表
   *
   * @param category 字典分类
   * @param name 字典名称（模糊查询）
   * @returns 字典列表
   */
  getList(category?: string, name?: string) {
    return request<any, DictListResponse>({
      url: `${DICT_BASE_URL}/`,
      method: "get",
      params: { category, name },
    });
  },

  /**
   * 获取字典详情
   *
   * @param code 字典编码
   * @param name 字典名称（模糊查询）
   * @returns 字典详情
   */
  getDetail(code: string, name?: string) {
    return request<any, DictDetailResponse>({
      url: `${DICT_BASE_URL}/${code}/list`,
      method: "get",
      params: { name },
    });
  },

  /**
   * 创建字典
   *
   * @param data 字典数据
   */
  create(data: DictForm) {
    return request<any, DictDetailResponse>({
      url: `${DICT_BASE_URL}/`,
      method: "post",
      data,
    });
  },

  /**
   * 更新字典
   *
   * @param code 字典编码
   * @param data 字典数据
   */
  update(code: string, data: DictForm) {
    return request<any, DictDetailResponse>({
      url: `${DICT_BASE_URL}/${code}`,
      method: "put",
      data,
    });
  },

  /**
   * 删除字典
   *
   * @param code 字典编码
   */
  delete(code: string) {
    return request<any, DictResponse>({
      url: `${DICT_BASE_URL}/${code}`,
      method: "delete",
    });
  },

  /**
   * 获取字典项列表（业务接口，别名覆盖name，不返回alias）
   *
   * @param dictCode 字典编码
   * @returns 字典项列表
   */
  getDictItems(dictCode: string) {
    return request<any, DictResponse[]>({
      url: `${DICT_BASE_URL}`,
      method: "get",
      params: { category: dictCode },
    });
  },
};

export default DictAPI;

/**
 * 字典项数据结构
 */
export interface DictItem {
  /**
   * 排序索引
   */
  index: number;
  /**
   * 分类
   */
  category: string;
  /**
   * 是否根节点
   */
  isRoot: boolean;
  /**
   * 字典编码
   */
  code: string;
  /**
   * 字典名称
   */
  name: string;
  /**
   * 别名，不为空时覆盖name字段
   */
  alias?: string;
  /**
   * 字典描述
   */
  desc: string;
}

/**
 * 字典列表响应
 */
export interface DictListResponse {
  code: number;
  message: string;
  data: DictItem[];
}

/**
 * 字典详情响应
 */
export interface DictDetailResponse {
  code: number;
  message: string;
  data: DictItem;
}

/**
 * 字典响应
 */
export interface DictResponse {
  code: number;
  name: string;
  message: string;
  tagType: string;
}

/**
 * 字典表单数据
 */
export interface DictForm {
  /**
   * 排序索引
   */
  index?: number;
  /**
   * 分类
   */
  category: string;
  /**
   * 是否根节点
   */
  isRoot: boolean;
  /**
   * 字典编码
   */
  code: string;
  /**
   * 字典名称
   */
  name: string;
  /**
   * 别名，不为空时覆盖name字段
   */
  alias?: string;
  /**
   * 字典描述
   */
  desc?: string;
}

/**
 * 字典查询参数
 */
export interface DictItemPageQuery extends PageQuery {
  /** 关键字(字典数据值/标签) */
  keywords?: string;

  /** 字典编码 */
  dictCode?: string;
}

/**
 * 字典分页对象
 */
export interface DictItemPageVO {
  /**
   * 字典ID
   */
  id: string;
  /**
   * 字典编码
   */
  dictCode: string;
  /**
   * 字典数据值
   */
  value: string;
  /**
   * 字典数据标签
   */
  label: string;
  /**
   * 状态（1:启用，0:禁用)
   */
  status: number;
  /**
   * 字典排序
   */
  sort?: number;
}

/**
 * 字典
 */
export interface DictItemForm {
  /**
   * 字典ID
   */
  id?: string;
  /**
   * 字典编码
   */
  dictCode?: string;
  /**
   * 字典数据值
   */
  value?: string;
  /**
   * 字典数据标签
   */
  label?: string;
  /**
   * 状态（1:启用，0:禁用)
   */
  status?: number;
  /**
   * 字典排序
   */
  sort?: number;

  /**
   * 标签类型
   */
  tagType?: "success" | "warning" | "info" | "primary" | "danger" | "";
}

/**
 * 字典项下拉选项
 */
export interface DictItemOption {
  value: number | string;
  label: string;
  tagType?: "" | "success" | "info" | "warning" | "danger";
  [key: string]: any;
}
