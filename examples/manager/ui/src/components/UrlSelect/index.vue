<template>
  <el-select
    v-model="modelValue"
    v-bind="$attrs"
    :loading="loading"
  >
    <template v-for="(slot, slotName) in $slots" #[slotName]="slotProps" :key="slotName">
      <slot :name="slotName" v-bind="slotProps" />
    </template>

    <el-option
      v-for="item in Options"
      :key="item.value"
      :value="item.value"
      :label="item.label"
    />
  </el-select>
</template>

<script setup lang="ts">
import { get } from "@/utils/request";

interface UrlSelectProps {
  url: string;
  pk?: string;
  nameCol?: string;
  queryParams?: Record<string, any>;
}

const props = withDefaults(defineProps<UrlSelectProps>(), {
  pk: "code",
  nameCol: "name",
  queryParams: () => ({}),
});

interface option {
  value: string | number;
  label: string;
}

const Options = ref<option[]>([])
const loading = ref(false);

const modelValue = defineModel<string | number | string[] | undefined>({})

watch(
  () => [props.url, props.queryParams],
  async () => {
    if (!props.url) return;

    loading.value = true;
    try {
      const res: any[] = await get(props.url, props.queryParams);
      Options.value = res.map((item) => ({
        value: item[props.pk],
        label: item[props.nameCol] ?? String(item[props.pk] ?? ''),
      }));
    } catch (error) {
      console.error('UrlSelect fetch failed:', error);
      Options.value = [];
    } finally {
      loading.value = false;
    }
  },
  { deep: true, immediate: true }
);

</script>
