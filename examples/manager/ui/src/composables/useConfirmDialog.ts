import { createApp, h, ref } from 'vue'
import ConfirmDialog from './../components/ConfirmDialog/index.vue'

interface ConfirmOptions {
  title?: string
  message: string
  confirmText?: string
  cancelText?: string
  type?: 'warning' | 'info' | 'error' | 'success'
}

const defaultOptions: Partial<ConfirmOptions> = {
  title: '提示',
  confirmText: '确认',
  cancelText: '取消',
  type: 'info'
}

export function useConfirmDialog() {
  const showConfirm = (options: ConfirmOptions): Promise<boolean> => {
    return new Promise((resolve) => {
      // 创建挂载容器
      const container = document.createElement('div')
      document.body.appendChild(container)

      // 合并选项
      const dialogOptions = { ...defaultOptions, ...options }

      // 创建应用实例
      const app = createApp({
        setup() {
          const visible = ref(true)

          const handleConfirm = () => {
            visible.value = false
            setTimeout(() => {
              app.unmount()
              document.body.removeChild(container)
            }, 300)
            resolve(true)
          }

          const handleCancel = () => {
            visible.value = false
            setTimeout(() => {
              app.unmount()
              document.body.removeChild(container)
            }, 300)
            resolve(false)
          }

          const handleClose = () => {
            visible.value = false
            setTimeout(() => {
              app.unmount()
              document.body.removeChild(container)
            }, 300)
            resolve(false)
          }

          return () => h(ConfirmDialog, {
            modelValue: visible.value,
            title: dialogOptions.title,
            message: dialogOptions.message,
            confirmButtonText: dialogOptions.confirmText,
            cancelButtonText: dialogOptions.cancelText,
            // type: dialogOptions.type,
            onConfirm: handleConfirm,
            onCancel: handleCancel,
            'onUpdate:visible': (value: boolean) => {
              if (!value) handleClose()
            }
          })
        }
      })

      // 挂载应用
      app.mount(container)
    })
  }

  return {
    confirm: showConfirm
  }
}

// 全局单例
let confirmInstance: ReturnType<typeof useConfirmDialog>

export function withConfirm(options: ConfirmOptions | string): Promise<boolean> {
  if (!confirmInstance) {
    confirmInstance = useConfirmDialog()
  }

  if (typeof options === 'string') {
    return confirmInstance.confirm({ message: options })
  }

  return confirmInstance.confirm(options)
}
