export const useMermaidStore = defineStore('mermaid', () => {
  const mermaidInstance = ref()
  const isInitialized = ref(false)
  const error = ref(null)
  let initPromise: Promise<void> | null = null

  const initializeMermaid = async () => {
    if (isInitialized.value) return

    if (initPromise) {
      return initPromise
    }

    initPromise = (async () => {
      error.value = null

      try {
        const mermaidModule = await import('mermaid')
        mermaidInstance.value = mermaidModule.default

        mermaidInstance.value.initialize({
          startOnLoad: false,
          theme: 'default',
          securityLevel: 'loose',
          fontFamily: 'inherit',
          flowchart: {
            htmlLabels: true,
            curve: 'basis'
          }
        })

        isInitialized.value = true
      } catch (err: any) {
        error.value = err.message
        console.error('Mermaid 初始化失败:', err)
      }
    })()

    return initPromise
  }

  const renderChart = async (code: string, container: any) => {
    if (!isInitialized.value) {
      await initializeMermaid()
    }

    if (!mermaidInstance.value || !container) {
      throw new Error('Mermaid 未初始化或容器不存在')
    }

    try {
      const id = `mermaid-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`
      const { svg } = await mermaidInstance.value.render(id, code)
      container.innerHTML = svg

      // 返回 SVG 元素以便后续操作
      return container.querySelector('svg')
    } catch (err) {
      console.log('Mermaid 渲染错误:', err)
      // throw err
    }
  }

  return {
    mermaidInstance,
    isInitialized,
    error,
    initializeMermaid,
    renderChart
  }
})
