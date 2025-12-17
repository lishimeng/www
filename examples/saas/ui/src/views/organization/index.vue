<template>
  <div>
    <Members v-if="org" :organization="org"/>
    <Organizations v-else/>
  </div>
</template>

<script setup lang="ts">
import Organizations from "@/views/organization/components/Organizations.vue";
import Members from "@/views/organization/components/Members.vue";

const route = useRoute();
const org = computed(() => {
  const code = route.query.org;
  if (Array.isArray(code)) {
    // 过滤掉 null/undefined/空字符串，取第一个非空值
    return code
      .filter((item): item is string => typeof item === 'string' && item.trim() !== '')
      .map(str => str.trim())[0] || undefined; // 没找到返回 null
  } else if (typeof code === 'string' && code.trim() !== '') {
    return code.trim();
  } else {
    return undefined;
  }
});

</script>
