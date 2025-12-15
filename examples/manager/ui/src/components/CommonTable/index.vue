<template>
  <div
    class="flex flex-col"
    :class="{
    'full-screen': props.fullScreen,
    'p-3 bg-[var(--el-bg-color)] border-2 border-[var(--el-border-color-light)] mx-2 my-3': !props.noBorder
  }"
  >
    <slot name="title"></slot>

    <!--  筛选表单  -->
    <div v-if="hasSearchItems" class="mb-5 p-4 pb-0 bg-[var(--el-border-color-extra-light)]">
      <el-form :inline="true" :model="state.searchForm" label-position="right">
        <!--   如果queryParams中有同名字段，则自动隐藏对应的select    -->
        <template v-for="field in props.searchItems">
          <el-form-item
            v-if="!Object.hasOwn(queryParams, field.prop)"
            :key="field.prop"
            :label="field.label"
            :prop="field.prop"
          >
            <el-input
              v-if="field.component === 'input'"
              v-model="state.searchForm[field.prop]"
              :class="{'mb-4': props.inDialog}"
              style="width: 220px"
              clearable
              :placeholder="field.placeholder"
            />

            <el-select
              v-else-if="field.component === 'select'"
              v-model="state.searchForm[field.prop]"
              :class="{'mb-4': props.inDialog}"
              style="width: 220px"
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
              style="width: 220px"
              :class="{'mb-4': props.inDialog}"
              type="date"
              clearable
              :placeholder="field.placeholder"
            />

            <TableSelect
              v-else-if="field.component === 'table-select'"
              v-model="state.searchForm[field.prop]"
              style="width: 220px"
              :class="{'mb-4': props.inDialog}"
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
              style="width: 220px"
              :class="{'mb-4': props.inDialog}"
              :placeholder="field.placeholder"
              :dict-code="field.dictCode!"
              clearable
            />
          </el-form-item>
        </template>
        <el-form-item>
          <div class="flex gap-col-2" :class="{'mb-4': props.inDialog}">
            <div class="flex-1"/>
            <el-button type="primary" @click="doQuery">查询</el-button>
            <el-button type="primary" plain @click="reset">重置</el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>

    <slot name="head"></slot>

    <!--  表格主体  -->
    <el-table
      ref="tableRef"
      v-loading="props.loading ?? state.loading"
      class="flex-1"
      :data="props.data ?? state.pagination.tableData"
      border
      :row-key="props.pk"
      v-bind="$attrs"
      @sort-change="handleSortChange"
      @selection-change="selectChange"
    >
      <slot name="extra-column"></slot>
      <el-table-column v-if="props.autoIndex" fixed type="index" label="序号" width="80" />
      <el-table-column
        v-for="column in props.columns"
        :key="column.prop"
        :prop="column.prop"
        :label="column.label"
        :width="getWidth(column)"
        min-width="120"
        :show-overflow-tooltip="column.showOverflowTooltip !== false"
        :fixed="column.fixed"
        :sortable="column.sortable"
      >
        <template #default="scope">
          <template v-if="column.component === 'time'">
            <template v-if="scope.row[column.prop]">
              {{ formatDate(new Date(scope.row[column.prop]), "YYYY-mm-dd HH:MM:SS") }}
            </template>
            <span v-else class="flex justify-center text-[var(--el-text-color-secondary)]">-</span>
          </template>
          <template v-else-if="column.component === 'enum'">
            {{ column.enums?.find((e) => e.value === scope.row[column.prop])?.label }}
          </template>
          <template v-else-if="column.component === 'enum_tag'">
            <el-tag
              v-if="column.enums?.find((e) => e.value === scope.row[column.prop])"
              :type="column.enums?.find((e) => e.value === scope.row[column.prop])?.type"
            >
              {{ column.enums?.find((e) => e.value === scope.row[column.prop])?.label }}
            </el-tag>
          </template>
          <template v-else-if="column.component === 'enum_dict'">
            <DictText
              :item-code="scope.row[column.prop]"
              :dict-code="column.dictCode!"
            />
          </template>
          <template v-else-if="column.component === 'custom'">
            <slot :name="column.prop" v-bind="scope">
              {{ scope.row[column.prop] }}
            </slot>
          </template>
          <template v-else>
            <ForeignKey
              v-if="column.fk"
              :url="column.fk.url"
              :field="column.fk.field"
              :code="scope.row[column.prop]"
            />
            <span v-else>
              {{ scope.row[column.prop] }}
            </span>
          </template>
        </template>
      </el-table-column>
      <el-table-column
        v-if="showOperationsColumn"
        :width="operationWidth"
        fixed="right"
        label="操作"
      >
        <template #default="scope">
          <div>
            <slot name="operation" v-bind="scope"></slot>
          </div>
        </template>
      </el-table-column>
    </el-table>

    <div ref="operationRef" class="measure-container">
      <slot name="operation" :row="{ isMeasure: true } as Record<string,any>"></slot>
    </div>

    <!--  分页器  -->
    <div v-if="state.pagination.enabled" class="m-2 mt-4 flex justify-end">
      <el-pagination
        :current-page="state.pagination.pageNo"
        :page-size="state.pagination.pageSize"
        :page-sizes="state.pagination.pageSizes"
        :page-count="state.pagination.totalPage"
        :total="state.pagination.total"
        :layout="state.pagination.layout"
        background
        @size-change="sizeChange"
        @current-change="currentChange"
      />
    </div>

    <slot name="footer"></slot>
  </div>
</template>

<script setup lang="ts">
import { SearchItem, TableColumns } from "@/components/CommonTable/types";
import { fetchListApi } from "@/api/common/tabel.api";
import { formatDate } from "@/utils/formatTime";

interface CommonTableProps {
  url?: string;
  data? :any[];
  searchItems?: SearchItem[];
  columns: TableColumns[];
  pk?: string;
  autoIndex?: boolean;
  showOperationsColumn?: boolean;
  queryParams?: Record<string, any>;
  fullScreen?: boolean;
  noBorder?:boolean;
  inDialog?: boolean;  // 解决el-form在dialog内外渲染不统一的问题
  loading?: boolean;
}

const props = withDefaults(defineProps<CommonTableProps>(), {
  searchItems: () => [],
  pk: "id",
  autoIndex: false,
  showOperationsColumn: true,
  queryParams: () => ({}),
  fullScreen: true,
  inDialog: false,
  noBorder: false,
  loading: undefined,
});

const state = reactive({
  searchForm: {} as Record<string, any>,
  loading: true,
  pagination: {
    enabled: false,
    totalPage: undefined,
    total: undefined,
    pageNo: 1,
    pageSize: 10,
    pageSizes: [10, 20, 30, 50],
    tableData: [] as any[],
    layout: "prev, pager, next, sizes, jumper",
  },
});

const operationRef = ref();
const tableRef = ref();
const operationWidth = ref(0);
const sortField = ref("");
const sortOrder = ref("");

const emit = defineEmits(["selection-change", "loaded"]);

function selectChange(selected: any[]) {
  emit("selection-change", selected);
}

function toggleSelection(rows?: any[], ignoreSelectable?: boolean) {
  if (rows) {
    rows.forEach((row) => {
      tableRef.value!.toggleRowSelection(
        row,
        undefined,
        ignoreSelectable
      )
    })
  } else {
    tableRef.value!.clearSelection()
  }
}

function updateOperationWidth() {
  if (operationRef.value) {
    operationWidth.value = operationRef.value.clientWidth + 20;
    if (operationWidth.value < 120) {
      operationWidth.value = 120;
    }
  }
}

function fetchPageData() {
  if (!props.url) {
    return;
  }
  state.loading = true;
  const queryParams: any = {
    ...state.searchForm,
    ...props.queryParams,
  };
  if (state.pagination.enabled) {
    queryParams.pageNum = state.pagination.pageNo;
    queryParams.pageSize = state.pagination.pageSize;
  }
  if (sortField.value !== "") {
    queryParams.sortField = sortField.value;
    queryParams.sortOrder = sortOrder.value;
  }
  fetchListApi(props.url, queryParams)
    .then((data) => {
      if (Array.isArray(data) || !data) {
        // 不启用分页
        state.pagination.enabled = false;
        state.pagination.tableData = data;
      } else {
        state.pagination.enabled = true;
        state.pagination.tableData = data.items;
        if (data.totalPage != null) {
          state.pagination.totalPage = data.totalPage;
        }
        if (data.total != null) {
          state.pagination.total = data.total;
        }
        if (data.pageSize != null) {
          state.pagination.pageSize = data.pageSize;
        }
        if (data.pageNo != null) {
          state.pagination.pageNo = data.pageNo;
        }
        // 如果size不在state.pagination.pageSizes中，将其插入并维持顺序
        if (!state.pagination.pageSizes.includes(data.pageSize)) {
          state.pagination.pageSizes.push(data.pageSize);
          state.pagination.pageSizes.sort((a, b) => a - b);
        }
        if (data.total) {
          state.pagination.layout = "prev, pager, next, total, sizes, jumper";
        } else {
          state.pagination.layout = "prev, pager, next, sizes, jumper";
        }
      }
      state.loading = false;
      emit("loaded");
    })
    .catch((err) => {
      console.log(err);
      state.loading = false;
    });
}

function currentChange(val: number) {
  state.pagination.pageNo = val;
  fetchPageData();
}

function sizeChange(val: number) {
  state.pagination.pageSize = val;
  state.pagination.pageNo = 1;
  fetchPageData();
}

interface SortChangeParam {
  prop: any,
  order: string,
}

function handleSortChange(param: SortChangeParam) {
  sortField.value = param.prop;
  sortOrder.value = param.order === "ascending" ? "asc" : "desc";
  // 检查对应列的sortable属性是否为custom
  if (props.columns.find((column) => column.prop === param.prop)?.sortable === "custom") {
    state.pagination.pageNo = 1;
    fetchPageData();
  }
}

function doQuery() {
  state.pagination.pageNo = 1;
  fetchPageData();
}

function reset() {
  state.searchForm = {};
  doQuery();
}

function getWidth(col: TableColumns) {
  return col.width ?? (col.component === "time" ? "156px" : "auto");
}

function setCurrentRow(row:any) {
  tableRef.value.setCurrentRow(row);
}

function initSearchForm() {
  state.searchForm = {};
  for (const item of props.searchItems) {
    state.searchForm[item.prop] = "";
  }
}

onMounted(async () => {
  fetchPageData();
  await nextTick();
  updateOperationWidth();
});

watch(() => props.searchItems, () => {
  initSearchForm();
}, { immediate: true })

defineExpose({
  fetchPageData,
  setCurrentRow,
  toggleSelection,
});

const hasSearchItems = computed(() => {
  return props.searchItems && props.searchItems.length > 0;
});
</script>

<style scoped lang="scss">
/* 使用深度选择器，修复fixed列阴影不显示的问题 */
:deep(.el-table__fixed) {
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.15);
  z-index: 3;
}

:deep(.el-table__fixed-right) {
  box-shadow: -2px 0 8px rgba(0, 0, 0, 0.15);
  z-index: 3;
}

:deep(.el-table__fixed-right::before),
:deep(.el-table__fixed::before) {
  background-color: transparent;
}

.full-screen {
  //margin: 12px 8px;
  height: calc(-108px + 100vh);
}

.measure-container {
  position: absolute;
  top: -9999px;
  left: -9999px;
  visibility: hidden;
  white-space: nowrap;
}
</style>
