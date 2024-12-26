import { createRouter, createWebHistory } from 'vue-router'
import Nodes from '@/pages/Nodes.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'

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
          component: Nodes,
          meta: {
            title: 'Settings'
          }
        },
      ]
    },
  ]
})

export default router