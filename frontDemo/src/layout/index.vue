<template>
  <div class="main">
    <Left class="left-component" />
    <router-view></router-view>
    <Right class="right-component" />
  </div>
</template>

<script setup lang="ts">

import Left from '@/layout/left.vue'
import Right from '@/layout/right.vue'
</script>

<style scoped lang="scss">
.main {
  display: flex;
  justify-content: space-between;
  .left-component {
    width: 30%;
  }
  .right-component {
    width: 30%;
  }
}
</style>