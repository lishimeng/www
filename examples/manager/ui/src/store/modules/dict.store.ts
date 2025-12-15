import { store } from "@/store";
import DictAPI, { type DictItemOption } from "@/api/system/dict.api";

export const useDictStore = defineStore("dict", () => {
  // 字典数据缓存
  const dictCache = useStorage<Record<string, DictItemOption[]>>("dict_cache", {});

  // 请求队列（防止重复请求）
  const requestQueue: Record<string, Promise<void>> = {};

  /**
   * 缓存字典数据
   * @param dictCode 字典编码
   * @param data 字典项列表
   */
  const cacheDictItems = (dictCode: string, data: DictItemOption[]) => {
    dictCache.value[dictCode] = data;
  };

  const tag_types = ["success", "warning", "info", "primary", "danger"];

  /**
   * 加载字典数据（如果缓存中没有则请求）
   * @param dictCode 字典编码
   * @param force 是否在有缓存时强制更新
   */
  const loadDictItems = async (dictCode: string, force?: boolean) => {
    if (!force && dictCache.value[dictCode]) return;
    // 防止重复请求
    if (force || !requestQueue[dictCode]) {
      requestQueue[dictCode] = DictAPI.getDictItems(dictCode).then((data) => {
        if (data === null) {
          ElMessage.error("无法加载字典: " + dictCode);
          return
        }
        const dictData = data.map((item, index) => ({
          value: item.code,
          // name 字段已经包含了别名（如果有），直接使用 name
          label: item.name,
          tagType: item.tagType || tag_types[index],
        }));
        cacheDictItems(dictCode, dictData as DictItemOption[]);
        Reflect.deleteProperty(requestQueue, dictCode);
      });
    }
    await requestQueue[dictCode];
  };

  /**
   * 获取字典项列表
   * @param dictCode 字典编码
   * @returns 字典项列表
   */
  const getDictItems = (dictCode: string): DictItemOption[] => {
    return dictCache.value[dictCode] || [];
  };


  const getDictItemsReactive = (dictCode: string) => {
    return computed(() => dictCache.value[dictCode] || []);
  };

  /**
   * 移除指定字典项
   * @param dictCode 字典编码
   */
  const removeDictItem = (dictCode: string) => {
    if (dictCache.value[dictCode]) {
      Reflect.deleteProperty(dictCache.value, dictCode);
    }
  };

  /**
   * 清空字典缓存
   */
  const clearDictCache = () => {
    dictCache.value = {};
  };

  return {
    dictCache,
    loadDictItems,
    getDictItems,
    getDictItemsReactive,
    removeDictItem,
    clearDictCache,
  };
});

export function useDictStoreHook() {
  return useDictStore(store);
}
