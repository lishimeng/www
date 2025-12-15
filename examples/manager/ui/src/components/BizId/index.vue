<template>
  <el-input
    v-bind="$attrs"
    :model-value="props.modelValue"
    :readonly="autoGenerate"
    :placeholder="placeholder"
    :class="{ 'disabled': !props.isCreate || autoGenerate }"
    @update:model-value="handleInput"
  >
    <template #prefix>
      <el-icon v-if="loading" class="is-loading">
        <Loading />
      </el-icon>
    </template>
    <template #suffix>
      <el-checkbox
        v-if="props.isCreate"
        v-model="autoGenerate"
        :disabled="props.fixed"
        label="自动生成"
        style="white-space: nowrap"
        @change="handleAutoGenerateChange"
      />
    </template>
  </el-input>
</template>

<script setup lang="ts">
import { createBizId } from "@/api/system/bizId.api";
import { Loading } from "@element-plus/icons-vue";

const props = defineProps({
  modelValue: {
    type: String,
    required: true,
  },
  idCode: {
    type: String,
    required: true,
  },
  auto: {
    type: Boolean,
    default: true,
  },
  fixed: {
    type: Boolean,
    default: false,
  },
  bizData: {
    type: Object,
    default: () => ({}),
  },
  isCreate: {
    type: Boolean,
    required: true,
  },
});

const emit = defineEmits(["update:modelValue"]);

const loading = ref(false);
const autoGenerate = ref(false);

const generate = (data: object) => {
  loading.value = true;
  emit("update:modelValue", "");
  createBizId(props.idCode, data)
    .then((code) => {
      emit("update:modelValue", code);
    })
    .finally(() => {
      loading.value = false;
    });
};

// 透传 v-model
const handleInput = (value: string) => {
  emit("update:modelValue", value);
};

const handleAutoGenerateChange = (value: string | number | boolean) => {
  if (props.isCreate && value) {
    generate(props.bizData);
  }
};

const placeholder = computed(() => {
  return autoGenerate.value ? "正在生成编码" : "请输入编码";
});

const refresh = (data: object, isCreate: boolean) => {
  if (isCreate && autoGenerate.value) {
    generate(data);
  }
};

onMounted(() => {
  autoGenerate.value = props.auto;
  // refresh(props.bizData, props.isCreate);
});

defineExpose({
  generate,
  refresh,
});
</script>

<style scoped lang="scss">
.disabled {
  :deep(.el-input__wrapper) {
    background-color: var(--el-fill-color-light, #f5f7fa);
    color: var(--el-text-color-disabled, #c0c4cc);
    .el-input__inner {
      cursor: not-allowed;
    }
  }
}
</style>
