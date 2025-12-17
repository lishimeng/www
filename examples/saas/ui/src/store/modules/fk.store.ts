import { fetchOneApi } from "@/api/common/tabel.api";

type obj = Record<string, any>
type cls = Record<string, obj>
type subscriber = Record<string, number>
type pendingRequest = string[]

export const useFkStore = defineStore("fk", () => {
  const cache = ref<Record<string, cls>>({})
  const subscribers = ref<Record<string, subscriber>>({})
  const pendingRequests = ref<Record<string, pendingRequest>>({})
  const cleanupScheduled = ref(false)

  const get = computed(() => {
    return (classPath: string, code: string) => {
      return cache.value[classPath]?.[code];
    };
  })

  async function requestData(classPath: string, code: string) {
    if (cache.value[classPath]?.[code]) {
      return
    }
    if (!pendingRequests.value[classPath]) {
      pendingRequests.value[classPath] = [code]
    } else {
      if (pendingRequests.value[classPath].includes(code)) {
        return
      }
      pendingRequests.value[classPath].push(code)
    }
    try {
      const data = await fetchOneApi(classPath, code)
      if (data) {
        if (!cache.value[classPath]) {
          cache.value[classPath] = {}
        }
        cache.value[classPath][code] = data
      }
      pendingRequests.value[classPath] = pendingRequests.value[classPath].filter(c => c !== code)
    } catch (err: any) {
      console.error(err)
    }
  }

  function subscribe(classPath: string, code: string) {
    if (!classPath || !code) {
      return
    }
    if (!subscribers.value[classPath]) {
      subscribers.value[classPath] = {}
    }
    subscribers.value[classPath][code] = (subscribers.value[classPath][code]??0) + 1;
    requestData(classPath, code)
  }

  function unsubscribe(classPath: string, code: string) {
    // console.log('unsubscribe', classPath, code)
    if (!subscribers.value[classPath]?.[code]) {
      return
    }
    subscribers.value[classPath][code] -= 1

    // 如果计数为0，安排清理
    if (subscribers.value[classPath][code] <= 0) {
      delete subscribers.value[classPath][code]
      scheduleCleanup()
    }
  }

  function scheduleCleanup() {
    if (cleanupScheduled.value) {
      return // 已经安排了清理，避免重复
    }

    cleanupScheduled.value = true
    nextTick(() => {
      performCleanup()
      cleanupScheduled.value = false
    })
  }

  function performCleanup() {
    // 清理没有订阅者的缓存
    Object.keys(cache.value).forEach(classPath => {
      Object.keys(cache.value[classPath] || {}).forEach(code => {
        if (!subscribers.value[classPath]?.[code]) {
          delete cache.value[classPath][code]
        }
      })

      // 如果该 classPath 没有缓存项了，删除空对象
      if (Object.keys(cache.value[classPath] || {}).length === 0) {
        delete cache.value[classPath]
      }
    })
    // console.log('cleanup', cache.value)
  }

  return {
    get,
    subscribe,
    unsubscribe,
  }
})
