import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { Host } from '@/service/AgentService'

export const useAgentStore = defineStore('counter', () => {
  const agents = ref<Record<string, Host>>({})
  return { agents }
})
