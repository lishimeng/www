<template>
  <div v-if="hasSearchItems" class="mb-5 p-4 pb-0 bg-[var(--el-border-color-extra-light)]">
    <!-- 主筛选容器 -->
    <div ref="formContainer" class="flex flex-wrap items-end gap-x-4 gap-y-2">
      <!-- 渲染可见的筛选项 -->
      <template v-for="(field, index) in visibleSearchItems" :key="field.prop">
        <div
          v-if="!Object.hasOwn(queryParams, field.prop)"
          ref="itemRefs"
          class="min-w-0 flex items-center"
          :style="{ minWidth: 'fit-content' }"
        >
          <label class="mr-2 text-sm text-[var(--el-text-color-regular)] whitespace-nowrap">
            {{ field.label }}：
          </label>
          <!-- 根据组件类型渲染 -->
          <el-input
            v-if="field.component === 'input'"
            v-model="state.searchForm[field.prop]"
            class="w-52"
            clearable
            :placeholder="field.placeholder"
          />
          <el-select
            v-else-if="field.component === 'select'"
            v-model="state.searchForm[field.prop]"
            class="w-52"
            clearable
            :placeholder="field.placeholder"
          >
            <el-option
              v-for="option in field.options"
              :key="option.value"
              :label="option.label"
              :value="option.value"
            />
          </el-select>
          <el-date-picker
            v-else-if="field.component === 'date-picker'"
            v-model="state.searchForm[field.prop]"
            class="w-52"
            type="date"
            clearable
            :placeholder="field.placeholder"
          />
          <TableSelect
            v-else-if="field.component === 'table-select'"
            v-model="state.searchForm[field.prop]"
            class="w-52"
            :placeholder="field.placeholder"
            :url="field.table!.url"
            :title="field.table!.title"
            :columns="field.table!.columns"
            :search-items="field.table!.search"
            clearable
          />
          <DictSelect
            v-else-if="field.component === 'dict-select'"
            v-model="state.searchForm[field.prop]"
            class="w-52"
            :placeholder="field.placeholder"
            :dict-code="field.dictCode!"
            clearable
          />
        </div>
      </template>

      <!-- 折叠按钮 -->
      <div v-if="isOverflow || isFolded" class="flex items-center">
        <el-button link type="primary" size="small" @click="toggleFold">
          {{ isFolded ? '展开' : '收起' }}
          <el-icon class="ml-1">
            <component :is="isFolded ? Expand : Fold" />
          </el-icon>
        </el-button>
      </div>

      <!-- 查询 & 重置按钮（始终在最右侧） -->
      <div class="ml-auto flex gap-2 whitespace-nowrap">
        <el-button type="primary" @click="doQuery">查询</el-button>
        <el-button type="primary" plain @click="reset">重置</el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick, computed, watch } from 'vue'
import { Expand, Fold } from '@element-plus/icons-vue'
import { SearchItem } from "@/components/CommonTable/types";

// props & emits
const props = defineProps<{
  searchItems: SearchItem[]
  inDialog?: boolean
  queryParams: Record<string, any>
}>()

const emit = defineEmits(['query', 'reset'])

// state
const state = reactive({
  searchForm: {} as Record<string, any>
})
const formContainer = ref<HTMLElement | null>(null)
const itemRefs = ref<HTMLElement[]>([])
const isFolded = ref(true) // 默认折叠
const visibleCount = ref(0) // 实际能完整显示的项数（不含按钮）

// computed
const hasSearchItems = computed(() => props.searchItems.length > 0)

// 可见项：折叠时只显示前 N 项（N = visibleCount），否则全显示
const visibleSearchItems = computed(() => {
  if (isFolded.value) {
    return props.searchItems.slice(0, visibleCount.value)
  } else {
    return props.searchItems
  }
})

// 是否有溢出（即原始项数 > 可见数）
const isOverflow = computed(() => {
  return props.searchItems.length > visibleCount.value
})

// 方法
const doQuery = () => emit('query', state.searchForm)
const reset = () => {
  Object.keys(state.searchForm).forEach((k) => {
    state.searchForm[k] = ''
  })
  emit('reset')
}
const toggleFold = () => {
  isFolded.value = !isFolded.value
}

// ✅ 核心：动态计算一行最多容纳多少项（不含按钮）
const calculateVisibleCount = async () => {
  await nextTick()
  if (!formContainer.value) return

  const containerWidth = formContainer.value.clientWidth
  if (containerWidth <= 0) return

  // 按钮宽度预留（“查询”+“重置”+ gap）
  const buttonAreaWidth = 160 // 估算：两个按钮 ~ 70+70 + gap 20

  // 计算每项宽度（含 label + control + margin）
  let count = 0

  // 🔁 改进：用离屏克隆法精确测量（避免干扰布局）
  count = measureMaxItemsInOneLine(props.searchItems, containerWidth, buttonAreaWidth)

  visibleCount.value = Math.max(1, count) // 至少显示1项
}

// 📏 离屏测量：创建隐藏容器，逐个添加项直到溢出
const measureMaxItemsInOneLine = (
  items: typeof props.searchItems,
  containerWidth: number,
  reservedButtonWidth: number
): number => {
  const testContainer = document.createElement('div')
  testContainer.style.position = 'fixed'
  testContainer.style.visibility = 'hidden'
  testContainer.style.width = `${containerWidth}px`
  testContainer.style.display = 'flex'
  testContainer.style.flexWrap = 'nowrap'
  testContainer.style.gap = '16px' // 与 gap-x-4 一致（1rem=16px）
  testContainer.style.padding = '0'
  document.body.appendChild(testContainer)

  try {
    let totalWidth = reservedButtonWidth
    let count = 0

    for (const field of items) {
      if (Object.hasOwn(props.queryParams, field.prop)) continue // 跳过隐藏项

      const itemEl = document.createElement('div')
      itemEl.style.minWidth = 'fit-content'
      itemEl.style.display = 'flex'
      itemEl.style.alignItems = 'center'
      itemEl.style.whiteSpace = 'nowrap'

      // 构造 label + control（仅尺寸，不需真实组件）
      const labelSpan = document.createElement('span')
      labelSpan.textContent = field.label + '：'
      labelSpan.style.marginRight = '8px'
      labelSpan.style.fontSize = '14px'
      labelSpan.style.fontWeight = 'normal'
      labelSpan.style.color = 'var(--el-text-color-regular)'

      const controlSpan = document.createElement('span')
      controlSpan.style.width = '208px' // el-input/select 默认 220px 减去 padding/border ≈ 208
      controlSpan.style.height = '32px' // 默认 height
      controlSpan.style.display = 'inline-block'

      itemEl.appendChild(labelSpan)
      itemEl.appendChild(controlSpan)
      testContainer.appendChild(itemEl)

      // 强制 layout
      const itemWidth = itemEl.getBoundingClientRect().width
      totalWidth += itemWidth + (count > 0 ? 16 : 0) // 第一项无前 gap

      if (totalWidth > containerWidth) {
        break
      }
      count++
    }

    return count
  } finally {
    document.body.removeChild(testContainer)
  }
}

// 🔄 响应式监听
onMounted(() => {
  const observer = new ResizeObserver(() => {
    calculateVisibleCount()
  })
  if (formContainer.value) {
    observer.observe(formContainer.value)
  }
  onUnmounted(() => observer.disconnect())

  // 首次计算
  calculateVisibleCount()
})

// 当 searchItems / queryParams 变化时重新测量
watch([() => props.searchItems.length, () => props.queryParams], () => {
  nextTick(() => {
    calculateVisibleCount()
  })
})
</script>

<style scoped>
/* 确保 label 不换行 */
label {
  white-space: nowrap;
}
</style>
