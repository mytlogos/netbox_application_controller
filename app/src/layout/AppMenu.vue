<script setup lang="ts">
import { ref } from 'vue'

const model = ref<
  Array<{ label: string; items?: any[]; separator?: boolean; icon?: string; route?: string }>
>([
  {
    label: 'Home',
    icon: 'pi pi-fw pi-home',
    route: '/'
  },
  {
    label: 'Agents',
    icon: 'pi pi-fw pi-desktop',
    route: '/agents'
  }
])
</script>

<template>
  <p-menu :model="model">
    <template #item="{ label, item, props }">
      <router-link v-if="item.route" v-slot="routerProps" :to="item.route" custom>
        <a
          :href="routerProps.href"
          v-bind="props.action"
          @click="routerProps.navigate"
          :class="{ 'p-menuitem-active': routerProps.isActive }"
        >
          <span v-bind="props.icon" />
          <span v-bind="props.label">{{ label }}</span>
        </a>
      </router-link>
      <a v-else :href="item.url" v-bind="props.action">
        <span v-bind="props.icon" />
        <span v-bind="props.label">{{ label }}</span>
      </a>
    </template>
  </p-menu>
</template>
