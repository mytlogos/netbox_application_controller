<script setup lang="ts">
import AgentService from '@/service/AgentService'
import { Base } from '@/service/BasicService'
import { ref } from 'vue'

const host = ref('')
const config = ref('')

const agentService = new AgentService(Base)

function downloadConfig() {
  const agent = agentService.deploy({
    Host: host.value,
    Type: 'manual'
  })
  agent
    .then((value) => {
      return agentService.getConfig(value.Uuid)
    })
    .then((value) => {
      config.value = JSON.stringify(value, null, 2)
    })
}

// Maybe use a multi step form: https://primevue.org/steps
</script>

<template>
  <p-card>
    <template #content>
      <div class="field">
        <label for="host">Host</label>
        <input-text id="host" type="text" :model-value="host" placeholder="Address of the Host" />
      </div>
      <p-card v-if="config">
        <template #content>
            {{ config }}
        </template>
      </p-card>
      <p-button label="Deploy" @click="downloadConfig" />
    </template>
  </p-card>
</template>
