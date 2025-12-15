<template>
  <div class="flex w-full items-center justify-center mt-2 gap-4 text-sm text-gray-400">
    <template v-if="state.initialized">
      <span>{{ state.copyright }}</span>
      <span>Version: {{ state.version }}</span>
      <span>Commit: {{ state.commit }}</span>
      <span>Built at: {{ state.build }}</span>
    </template>
  </div>
</template>

<script setup lang="ts">
import { getVersionApi } from "@/api/system/version.api";

const state = reactive({
  copyright: "",
  version: "",
  commit: "",
  build: "",
  initialized: false,
});

onMounted(() => {
  getVersionApi().then((data) => {
    state.copyright = data.copyright;
    state.version = data.version;
    state.commit = data.commit;
    state.build = data.build;
    state.initialized = true;
  }).catch((err) => {
    console.log(err)
  })
})
</script>
