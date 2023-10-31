<template>
  <div class="menu-list">
    <ul>
      <el-link :underline="false" class="menu-item-a" @click="goPath('/')">
        <li>
          <el-icon :size="22" class="icon"><Watermelon /></el-icon>
          <span>1</span>
        </li>
      </el-link>
    </ul>
  </div>
</template>

<script setup lang="ts">

import { useRouter } from 'vue-router'
import { Watermelon } from '@element-plus/icons-vue'

const router = useRouter()
function goPath(e: string) {
  router.push(e)
}
</script>

<style scoped lang="scss">
.menu-list {
  padding-left: 10px;
  padding-top: 50px;

  ul {
    list-style: none !important;
    display: flex;
    flex-direction: column;
  }

  li {
    display: flex;
    align-items: center;
  }
}
.icon {
  margin-right: 10px;
}
.menu-item-a {
  margin-bottom: 20px;
  color: #646ebf;
}
.menu-item-a :hover {
  color: #b7bdfa;
}
a {
  text-decoration: none !important;
}
</style>