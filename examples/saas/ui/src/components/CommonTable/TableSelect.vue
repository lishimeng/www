<template>
  <div class="w-full">
    <el-input
      :model-value="rendered"
      class="w-full"
      :input-style="{ cursor: props.disabled ? 'not-allowed' : 'pointer' }"
      v-bind="$attrs"
      readonly
      role="button"
      :placeholder="loading ? '加载中...' : props.placeholder"
      :disabled="props.disabled"
      @click="openDialog"
      @keydown.enter="openDialog"
    >
      <!--   插槽:清空按钮   -->
      <template #suffix>
        <el-icon
          v-if="props.clearable && props.modelValue"
          class="cursor-pointer"
          style="width: 1em; height: 1em; display: inline-flex; align-items: center; justify-content: center;"
          @click.stop="clear"
        >
          <CircleClose />
        </el-icon>
      </template>
    </el-input>
    <el-dialog
      v-model="dialogVisible"
      :title="props.title"
      append-to-body
      width="80%"
    >
      <CommonTable
        ref="tableRef"
        class="flex-1 h-[60vh]"
        :url="props.url"
        :columns="props.columns"
        :search-items="props.searchItems"
        :query-params="props.queryParams"
        :full-screen="false"
        :show-operations-column="false"
        highlight-current-row
        :pk="props.pk"
        @current-change="onSelect"
      >
        <!--    透传插槽    -->
        <template v-for="(slot, slotName) in $slots" #[slotName]="slotProps" :key="slotName">
          <slot :name="slotName" v-bind="slotProps" />
        </template>
      </CommonTable>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">

import { SearchItem, TableColumns } from "@/components/CommonTable/types";
import { CircleClose } from "@element-plus/icons-vue";
import { fetchOneApi } from "@/api/common/tabel.api";

interface TableSelectProps {
  modelValue: any;
  nameCol?: string;
  url: string;
  searchItems?: SearchItem[];
  queryParams?: Record<string, any>;
  columns: TableColumns[];
  pk?: string;
  clearable?: boolean;
  disableAutoName?: boolean;
  title?: string;
  name?: string;
  disabled?: boolean;
  placeholder?: string;
}

const props = withDefaults(defineProps<TableSelectProps>(), {
  pk: "code",
  title: "",
  nameCol: "name",
  name: "",
  clearable: false,
  disableAutoName: false,
  disabled: false,
  placeholder: "",
});

const emit = defineEmits(['update:modelValue', 'update:name', 'select']);

const tableRef = ref()
const rowCode = ref("")
const rowName = ref("")
const dialogVisible = ref(false);

const loading = ref(true);

const rendered = computed(() => {
  if (loading.value) return "";
  return props.name || rowName.value || props.modelValue;
})

async function openDialog() {
  dialogVisible.value = true;
  await nextTick();
  tableRef.value?.fetchPageData();
}

function onSelect(currentRow: Record<string, any>) {
  if (currentRow) {
    if (props.modelValue === currentRow[props.pk]) {
      return
    }
    emit("update:modelValue", currentRow[props.pk]);
    emit("update:name", currentRow[props.nameCol]);
    emit("select", currentRow);
    rowName.value = currentRow[props.nameCol];
    rowCode.value = currentRow[props.pk];
    dialogVisible.value = false;
  }
}

function clear() {
  emit("update:modelValue", "");
  emit("update:name", "");
  emit("select", {});
  rowName.value = '';
  rowCode.value = '';
  tableRef.value?.setCurrentRow()
}

// 监听props.modelValue
watch(() => props.modelValue, (newValue) => {
  loading.value = true;
  if (newValue === rowCode.value) {
    loading.value = false;
    return;
  }
  if (!newValue) {
    rowCode.value = '';
    rowName.value = '';
    tableRef.value?.setCurrentRow()
    loading.value = false;
    return;
  }
  rowCode.value = newValue;
  if (props.disableAutoName || props.name !== '') {
    loading.value = false;
    return;
  }
  rowName.value = '';
  fetchOneApi(props.url, newValue).then((res: Record<string, any>) => {
    rowName.value = res[props.nameCol];
    loading.value = false;
    emit("select", res);
  }).catch((err) => {
    console.log("Fetch one error: ", err)
    rowName.value = newValue;
    loading.value = false;
  })
}, {immediate: true})

</script>
