import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import Nodes from '@/pages/Nodes.vue'
import Settings from '@/pages/Settings.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: 'nodes',
          name: 'Nodes',
          component: Nodes,
          meta: {
            title: 'Nodes'
          }
        },
        {
          path: 'tasks',
          name: 'Tasks',
          component: Nodes,
          meta: {
            title: 'Tasks'
          }
        },
        {
          path: 'settings',
          name: 'Settings',
          component: Settings,
          meta: {
            title: 'Settings'
          }
        },
      ]
    },
  ]
})

export default router