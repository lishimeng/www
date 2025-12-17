<template>
  <el-select
    v-bind="$attrs"
    :model-value="props.modelValue"
    :options="options"
    :empty-text="emptyText"
    @update:model-value="handleInput"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useDictStore } from '@/store'

const props = defineProps<{
  modelValue: string | number | null,
  dictCode: string,
}>()

const loading = ref(false)

const emit = defineEmits(['update:modelValue'])

const dictStore = useDictStore()

const loadDict = async (dictCode: string) => {
  loading.value = true
  try {
    await dictStore.loadDictItems(dictCode)
  } finally {
    loading.value = false
  }
}

// 响应式加载字典
watch(
  () => props.dictCode,
  async (newCode) => {
    if (newCode) await loadDict(newCode)
  },
  { immediate: true }
)

const emptyText = computed(() => {
  if (loading.value) return '加载中...'
  const items = dictStore.dictCache[props.dictCode] || []
  return items.length === 0 ? '暂无数据，请先配置字典项' : ''
})

const options = computed(() => {
  return dictStore.dictCache[props.dictCode] || []
})

// 透传 v-model：将内部 el-select 的值变化 emit 出去
const handleInput = (value: string | number | null) => {
  emit('update:modelValue', value)
}
</script>
