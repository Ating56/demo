<template>
  <div class="right-switch">
    <div></div>
    <el-switch
      v-model='themeMode'
      :active-action-icon="Sunny"
      :inactive-action-icon="Moon"
      @change="handleChange"
      style="--el-switch-on-color: rgb(111, 122, 212);"
    />
  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'
import { Sunny, Moon } from '@element-plus/icons-vue'
import { useAppStore } from '@/store/app'

const appStore = useAppStore()
const themeMode = ref(localStorage.getItem('themeType') ? localStorage.getItem('themeType') === 'true' : false)

async function handleChange(val: any) {
  if(val) {
    appStore.changeTheme('#292c4f')
    appStore.toggleThemeMode('dark', '#292c4f')
    localStorage.setItem('themeType', 'true')
  } else {
    appStore.changeTheme('#fff')
    appStore.toggleThemeMode('day', '#fff')
    localStorage.setItem('themeType', 'false')
  }
}
</script>

<style scoped lang="scss">
.right-switch {
  padding-top: 30px;
  display: flex;
  justify-content: space-around;
}
</style>
