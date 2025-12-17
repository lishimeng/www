<template>
  <el-form-item
    v-bind="$attrs"
  >
    <template #label>
      <div class="flex-y-center">
        {{ label }}
        <el-tooltip
          v-if="tooltip"
          :placement="placement"
          :effect="effect"
        >
          <template #content>
            <slot name="tooltip-content">
              {{ tooltip }}
            </slot>
          </template>
          <el-icon class="ml-1 cursor-pointer">
            <QuestionFilled />
          </el-icon>
        </el-tooltip>
      </div>
    </template>
    <slot />
  </el-form-item>
</template>

<script setup lang="ts">
import { QuestionFilled } from '@element-plus/icons-vue'

interface Props {
  // 表单项标签
  label?: string
  // 提示内容
  tooltip?: string
  // 提示框位置
  placement?: 'top' | 'top-start' | 'top-end' | 'bottom' | 'bottom-start' | 'bottom-end' | 'left' | 'left-start' | 'left-end' | 'right' | 'right-start' | 'right-end'
  // 提示框主题
  effect?: 'light' | 'dark'
}

withDefaults(defineProps<Props>(), {
  label: '',
  tooltip: '',
  placement: 'top',
  effect: 'dark',
})

defineOptions({
  inheritAttrs: false
})
</script>

<style scoped>
.flex-y-center {
  display: flex;
  align-items: center;
}
</style>
