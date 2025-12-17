import { withConfirm } from '@/composables/useConfirmDialog'
import type { App } from "vue";

// 扩展到 Vue 实例
declare module 'vue' {
  interface ComponentCustomProperties {
    $withConfirm: typeof withConfirm
  }
}

// 定义插件对象
const confirmDialogPlugin = {
  install(app: App) {
    app.config.globalProperties.$withConfirm = withConfirm
  }
}

export function setupConfirmDialog(app: App) {
  app.use(confirmDialogPlugin)
}
