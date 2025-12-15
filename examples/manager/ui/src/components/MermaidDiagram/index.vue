<template>
  <div class="mermaid-container">
    <div v-show="isReady && !isNull" ref="mermaidRef" class="mermaid"></div>
    <div v-show="!isReady && !isNull" class="loading-state">
      <span>初始化中...</span>
    </div>
    <div v-show="isReady && isNull" class="loading-state">
      <span>无数据</span>
    </div>
    <div v-if="store.error" class="error-message">
      {{ store.error }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, nextTick } from 'vue'
import { useMermaidStore } from "@/store/index.js";

const props = defineProps({
  code: {
    type: String,
    required: true
  },
  interactive: {
    type: Boolean,
    default: true
  },
  onNodeClick: {
    type: Function,
    default: null
  }
})

const emit = defineEmits(['nodeClick', 'nodeHover', 'nodeDblClick'])

const store = useMermaidStore()
const mermaidRef = ref()
const isReady = ref(false)
const isNull = ref(false)
const selectedNode = ref(null)
// 添加事件监听器引用，用于清理
const eventListeners = ref(new Map())

onMounted(async () => {
  await initializeAndRender()
})

watch(() => props.code, async () => {
  await initializeAndRender()
})

// 在组件卸载时清理事件监听器
import { onUnmounted } from 'vue'
onUnmounted(() => {
  cleanupEventListeners()
})

const initializeAndRender = async () => {
  if (!props.code || typeof props.code !== 'string' || props.code.trim() === '') {
    isReady.value = true
    isNull.value = true
    return
  }
  isNull.value = false
  try {
    // 清理之前的事件监听器
    cleanupEventListeners()

    // 确保 Mermaid 已初始化
    await store.initializeMermaid()

    // 等待 DOM 更新
    await nextTick()

    // 渲染图表
    if (!mermaidRef.value) return
    const svgElement = await store.renderChart(props.code, mermaidRef.value)

    // 添加交互功能
    if (props.interactive && svgElement) {
      setupInteractions(svgElement)
    }

    isReady.value = true
  } catch (error) {
    console.error('渲染失败:', error)
    if (mermaidRef.value) {
      mermaidRef.value.innerHTML = `<div class="error">图表渲染失败: ${error.message}</div>`
    }
  }
}

const setupInteractions = (svgElement) => {
  if (!svgElement) return

  // 获取所有可交互的元素
  const nodes = svgElement.querySelectorAll('[id]')

  nodes.forEach(node => {
    const nodeId = node.getAttribute('id')
    const nodeType = getNodeType(node)
    const originalId = getOriginalNodeId(nodeId)

    // 创建事件处理器
    const clickHandler = (event) => {
      event.stopPropagation() // 阻止事件冒泡
      event.preventDefault() // 阻止默认行为
      handleNodeClick(event, node, originalId, nodeType)
    }

    const dblClickHandler = (event) => {
      event.stopPropagation()
      event.preventDefault()
      handleNodeDblClick(event, node, originalId, nodeType)
    }

    const mouseEnterHandler = (event) => {
      event.stopPropagation()
      handleNodeHover(event, node, originalId, nodeType, true)
    }

    const mouseLeaveHandler = (event) => {
      event.stopPropagation()
      handleNodeHover(event, node, originalId, nodeType, false)
    }

    // 添加样式和事件监听
    node.style.cursor = 'pointer'

    // 移除可能存在的旧监听器
    if (eventListeners.value.has(nodeId)) {
      const oldListeners = eventListeners.value.get(nodeId)
      Object.keys(oldListeners).forEach(eventType => {
        node.removeEventListener(eventType, oldListeners[eventType])
      })
    }

    // 添加新监听器
    node.addEventListener('click', clickHandler, { passive: false })
    node.addEventListener('dblclick', dblClickHandler, { passive: false })
    node.addEventListener('mouseenter', mouseEnterHandler)
    node.addEventListener('mouseleave', mouseLeaveHandler)

    // 保存监听器引用以便清理
    eventListeners.value.set(nodeId, {
      click: clickHandler,
      dblclick: dblClickHandler,
      mouseenter: mouseEnterHandler,
      mouseleave: mouseLeaveHandler
    })
  })
}

const getOriginalNodeId = (internalId) => {
  const match = internalId.match(/flowchart-(.+?)-/)
  if (match && match[1]) {
    return match[1]
  }

  return internalId // 回退到内部ID
}

const getNodeType = (node) => {
  if (node.tagName === 'rect') return 'process'
  if (node.tagName === 'circle' || node.tagName === 'ellipse') return 'start_end'
  if (node.tagName === 'polygon') return 'decision'
  if (node.tagName === 'path') return 'arrow'
  return 'unknown'
}

const handleNodeClick = (event, node, nodeId, nodeType) => {
  // 移除之前选中的节点样式
  // clearNodeSelection()

  // 添加选中样式
  // node.classList.add('mermaid-node-selected')

  // 获取节点文本
  const textElement = node.parentNode?.querySelector('.label') ||
    node.querySelector('.label') ||
    node.parentNode?.querySelector('text')
  const nodeText = textElement?.textContent || '未知'

  const nodeInfo = {
    id: nodeId,
    type: nodeType,
    text: nodeText,
    element: node,
    event: event
  }

  selectedNode.value = nodeInfo

  // 只触发一次 emit - 使用防抖确保
  if (props.onNodeClick) {
    props.onNodeClick(nodeInfo)
  }
  // emit('nodeClick', nodeInfo)
}

const handleNodeDblClick = (event, node, nodeId, nodeType) => {
  const textElement = node.parentNode?.querySelector('.label') ||
    node.querySelector('.label') ||
    node.parentNode?.querySelector('text')
  const nodeText = textElement?.textContent || '未知'

  const nodeInfo = {
    id: nodeId,
    type: nodeType,
    text: nodeText,
    element: node,
    event: event
  }

  emit('nodeDblClick', nodeInfo)
}

const handleNodeHover = (event, node, nodeId, nodeType, isEnter) => {
  if (nodeType === "arrow") {
    return
  }
  if (isEnter) {
    // node.classList.add('mermaid-node-hover')
    clearNodeSelection()

    // 添加选中样式
    node.classList.add('mermaid-node-selected')

    const textElement = node.parentNode?.querySelector('.label') ||
      node.querySelector('.label') ||
      node.parentNode?.querySelector('text')
    const nodeText = textElement?.textContent || '未知'

    emit('nodeHover', {
      id: nodeId,
      type: nodeType,
      text: nodeText,
      element: node,
      event: event,
      isHover: true
    })
  } else {
    node.classList.remove('mermaid-node-hover')
    clearNodeSelection()

    emit('nodeHover', {
      id: nodeId,
      type: nodeType,
      text: '',
      element: node,
      event: event,
      isHover: false
    })
  }
}

const clearNodeSelection = () => {
  const selectedNodes = mermaidRef.value?.querySelectorAll('.mermaid-node-selected')
  selectedNodes?.forEach(node => {
    node.classList.remove('mermaid-node-selected')
  })
}

const cleanupEventListeners = () => {
  eventListeners.value.forEach((listeners, nodeId) => {
    const node = mermaidRef.value?.querySelector(`[id="${nodeId}"]`)
    if (node) {
      Object.keys(listeners).forEach(eventType => {
        node.removeEventListener(eventType, listeners[eventType])
      })
    }
  })
  eventListeners.value.clear()
}

// 暴露方法给父组件
defineExpose({
  clearSelection: clearNodeSelection,
  getSelectedNode: () => selectedNode.value
})
</script>

<style scoped>
.mermaid-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  color: #666;
}

.error-message {
  color: #f56c6c;
  padding: 10px;
  background: #fef0f0;
  border-radius: 4px;
}

.empty-message {
  color: #999;
  font-style: italic;
  padding: 1rem;
}

/* 节点交互样式 */
:deep(.mermaid-node-selected) {
  stroke: #409eff !important;
  filter: drop-shadow(0 0 4px rgba(64, 158, 255, 0.6));
}

:deep(.mermaid-node-hover) {
  stroke: #67c23a !important;
  cursor: pointer;
}

:deep(.mermaid-node-selected rect),
:deep(.mermaid-node-hover rect) {
  stroke: inherit !important;
  stroke-width: inherit !important;
}

:deep(.mermaid-node-selected circle),
:deep(.mermaid-node-hover circle) {
  stroke: inherit !important;
  stroke-width: inherit !important;
}

:deep(.mermaid-node-selected polygon),
:deep(.mermaid-node-hover polygon) {
  stroke: inherit !important;
  stroke-width: inherit !important;
}
</style>
