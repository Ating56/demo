<template>
  <div>
    <wave></wave>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">

import wave from '@/components/wave.vue'

export default {
  name: 'App',
  components: { wave }
}
</script>

<style>
#app {
  height: 100vh;
  width: 100%;
  max-width: none;
  margin: 0;
  padding: 0;
  border: 0;
}
</style>
