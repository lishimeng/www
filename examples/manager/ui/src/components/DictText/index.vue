<template>
  {{ rendered }}
</template>

<script setup lang="ts">
import { useDictStore } from "@/store";

const props = defineProps({
  itemCode: {
    type: String,
    default: "",
  },
  dictCode: {
    type: String,
    required: true,
  },
});

const loading = ref(false)

const dictStore = useDictStore()

const loadDict = async (dictCode: string) => {
  loading.value = true
  try {
    await dictStore.loadDictItems(dictCode)
  } finally {
    loading.value = false
  }
}

const rendered = computed(() => {
  const items = dictStore.dictCache[props.dictCode];
  if (!items) return props.itemCode
  return items.find(item => item.value === props.itemCode)?.label || props.itemCode
})

// 响应式加载字典
watch(
  () => props.dictCode,
  async (newCode) => {
    if (newCode) await loadDict(newCode)
  },
  { immediate: true }
)
</script>
