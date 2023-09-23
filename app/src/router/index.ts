import { createRouter, createWebHistory } from 'vue-router'
import AppLayout from '../layout/AppLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          name: 'dashboard',
          component: () => import('@/views/Dashboard.vue')
        },
        {
          path: '/agents',
          name: 'agents',
          component: () => import('@/views/AgentList.vue')
        },
        {
          path: '/agent/:uuid',
          name: 'agent',
          props: true,
          component: () => import('@/views/Agent.vue')
        },
        {
          path: '/deploy-auto',
          name: 'deployagentauto',
          component: () => import('@/views/DeployAgentAuto.vue')
        },
        {
          path: '/deploy-manual',
          name: 'deployagentmanual',
          component: () => import('@/views/DeployAgentManual.vue')
        }
      ]
    }
  ]
})

export default router
