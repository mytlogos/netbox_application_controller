<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppFooter from './AppFooter.vue'
import AppSidebar from './AppSidebar.vue'
</script>

<template>
  <div class="flex gap-2">
    <div>
      <app-sidebar></app-sidebar>
    </div>
    <div class="flex-grow-1">
      <div class="layout-main">
        <router-view></router-view>
      </div>
      <app-footer></app-footer>
    </div>
  </div>
</template>
