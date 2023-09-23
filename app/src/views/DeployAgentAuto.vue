<script setup lang="ts">
import { ref } from 'vue'

const host = ref('')

function deploy() {
    
}
</script>

<template>
  <div>
    <input-text type="text" :model-value="host" />
    <p-button label="Deploy" @click="deploy" />
  </div>
</template>
