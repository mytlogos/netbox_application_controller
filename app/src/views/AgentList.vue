<script setup lang="ts">
import AgentService, { type Host } from '@/service/AgentService'
import { Base } from '@/service/BasicService'
import { useAgentStore } from '@/stores/agent'
import type { DataTableRowClickEvent } from 'primevue/datatable'
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const agents = ref<Host[]>([])
const tableAgents = computed(() =>
  agents.value.map((item) => ({
    uuid: item.Uuid,
    state: item.Agent.State,
    version: item.Agent.Version,
    lastUpdate: item.Agent.LastUpdate,
    item
  }))
)
const agentService = new AgentService(Base)
const router = useRouter()
const agentStore = useAgentStore()

function loadAgents() {
  agentService.getAll().then((value) => {
    agents.value = value
    const newValue: any = {}
    for (const agent of value) {
      newValue[agent.Uuid] = agent
    }
    agentStore.agents = newValue
  })
}

const onRowClick = (event: DataTableRowClickEvent) => {
  router.push(`/agent/${event.data.uuid}`)
}

loadAgents()

const items = [
  {
    label: 'Deploy Agent',
    icon: 'pi pi-fw pi-plus',
    items: [
      {
        label: 'Automatic',
        route: '/deploy-auto',
        icon: 'pi pi-fw pi-cloud-upload'
      },
      {
        label: 'Manual',
        route: '/deploy-manual',
        icon: 'pi pi-fw pi-user'
      }
    ]
  },
  {
    label: 'Refresh',
    icon: 'pi pi-fw pi-refresh',
    command() {
      loadAgents()
    }
  }
]
</script>

<template>
  <div>
    <p-menubar :model="items">
      <template #item="{ label, item, props, root }">
        <router-link v-if="item.route" v-slot="routerProps" :to="item.route" custom>
          <a :href="routerProps.href" v-bind="props.action" @click="routerProps.navigate">
            <span v-bind="props.icon" />
            <span v-bind="props.label">{{ label }}</span>
          </a>
        </router-link>
        <a v-else :href="item.url" :target="item.target" v-bind="props.action">
          <span v-bind="props.icon" />
          <span v-bind="props.label">{{ label }}</span>
          <span
            :class="[root ? 'pi pi-fw pi-angle-down' : 'pi pi-fw pi-angle-right']"
            v-bind="props.submenuicon"
          />
        </a>
      </template>
    </p-menubar>
    <p-data-table :value="tableAgents" dataKey="uuid" @row-click="onRowClick">
      <p-column field="uuid" header="Uuid" />
      <p-column field="state" header="State" />
      <p-column field="version" header="Version" />
      <p-column field="lastUpdate" header="Last Update" />
    </p-data-table>
  </div>
</template>
