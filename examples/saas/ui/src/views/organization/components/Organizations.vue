<template>
  <div>
    <CommonTable
      ref="tableRef"
      :url="orgUrl"
      name="组织"
      pk="code"
      :columns="orgColumns"
      full-screen
    >
      <template #head>
        <div class="flex mb-3">
          <el-button type="primary" @click="createRef.open()">新建组织</el-button>
        </div>
      </template>
      <template #operation="scope">
        <div class="flex">
          <el-button
            type="text" size="small"
            @click="router.push({name: 'Organizations', query: {org: scope.row.code}})"
          >成员列表</el-button>
        </div>
      </template>
    </CommonTable>

    <CreateOrg ref="createRef" @submit="tableRef.fetchPageData()" />
  </div>
</template>

<script setup lang="ts">
import {orgColumns, orgUrl} from "@/types/tables/orgs.table";
import CreateOrg from "@/views/organization/components/CreateOrg.vue";

const tableRef = ref();
const createRef = ref();
const router = useRouter();

</script>
