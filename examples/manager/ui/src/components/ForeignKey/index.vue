<template>
  <slot v-if="value" v-bind="{ value }">
    {{ value }}
  </slot>
  <el-skeleton v-else-if="props.code" :rows="0" animated />
  <span v-else class="flex justify-center text-[var(--el-text-color-secondary)]">-</span>
</template>

<script setup lang="ts">
import { useFkStore } from "@/store";

const props = defineProps({
  url: {
    type: String,
    required: true,
  },
  code: {
    type: String,
    default: "",
  },
  field: {
    type: String,
    default: "name",
  },
});

const fkStore = useFkStore();

const value = computed(() => {
  const obj = fkStore.get(props.url, props.code)
  if (obj) {
    return obj[props.field]
  } else {
    return ""
  }
});

onMounted(() => {
  fkStore.subscribe(props.url, props.code);
})

onUnmounted(() => {
  fkStore.unsubscribe(props.url, props.code);
})

watch(() => props.code, (newCode, oldCode) => {
  fkStore.unsubscribe(props.url, oldCode);
  fkStore.subscribe(props.url, newCode);
}, {deep: false})

watch(() => props.url, (newCode, oldCode) => {
  fkStore.unsubscribe(oldCode, props.code);
  fkStore.subscribe(newCode, props.code);
}, {deep: false})

</script>
