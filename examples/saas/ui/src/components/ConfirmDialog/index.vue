<template>
  <Transition name="dialog-fade">
    <div v-if="visible" class="fixed inset-0 z-3025">
      <!-- 遮罩层 -->
      <div
        class="absolute inset-0 bg-black bg-opacity-50"
        @click="handleCancel"
      ></div>

      <!-- 对话框 -->
      <div
        ref="dialogRef"
        class="dialog-container absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-128 bg-white rounded-xl shadow-xl overflow-hidden min-h-45"
        tabindex="-1"
      >
        <!-- 头部 -->
        <div class="flex items-center justify-between border-b border-gray-200">
          <h3 class="text-lg font-semibold text-gray-800 ml-6">{{ title }}</h3>
          <button
            class="w-8 h-8 flex items-center justify-center rounded-full hover:bg-gray-100 transition-colors mr-4"
            @click="handleCancel"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
          </button>
        </div>

        <!-- 内容 -->
        <div class="p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-10 h-10 rounded-full bg-red-100 flex items-center justify-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
            </div>
            <div class="ml-4">
              <p class="text-gray-600">{{ message }}</p>
            </div>
          </div>
        </div>

        <!-- 底部按钮 -->
        <div class="flex justify-end space-x-3 p-4 bg-gray-50">
          <button
            class="px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
            @click="handleCancel"
          >
            {{ cancelButtonText }}
          </button>
          <button
            class="px-4 py-2 text-white bg-red-600 rounded-lg hover:bg-red-700 transition-colors"
            @click="handleConfirm"
          >
            {{ confirmButtonText }}
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted, nextTick } from 'vue'

interface Props {
  modelValue: boolean
  title?: string
  message?: string
  confirmButtonText?: string
  cancelButtonText?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: '提示',
  message: '确定要执行此操作吗？',
  confirmButtonText: '确认',
  cancelButtonText: '取消'
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'confirm'): void
  (e: 'cancel'): void
}>()

const visible = ref(false)
const dialogRef = ref<HTMLElement | null>(null)

watch(() => props.modelValue, async (val) => {
  visible.value = val
  if (val) {
    // 等待DOM更新后聚焦到对话框
    await nextTick()
    dialogRef.value?.focus()
  }
})

// 处理ESC键关闭
const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && visible.value) {
    handleCancel()
  }
}

onMounted(() => {
  visible.value = props.modelValue
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})

function handleConfirm() {
  emit('confirm')
  emit('update:modelValue', false)
}

function handleCancel() {
  emit('cancel')
  emit('update:modelValue', false)
}
</script>

<style scoped>
.dialog-fade-enter-active {
  transition: all 0.3s ease-out;
}

.dialog-fade-leave-active {
  transition: all 0.3s ease-in;
}

.dialog-fade-enter-from,
.dialog-fade-leave-to {
  opacity: 0;
}

.dialog-fade-enter-from .dialog-container {
  transition: all 0.3s ease-out;
  transform: translate(-50%, -60%);
}
.dialog-fade-leave-to .dialog-container {
  transition: all 0.3s ease-out;
  transform: translate(-50%, -60%);
}
</style>
