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
          path: '',
          name: 'Nodes',
          component: Nodes,
          meta: {
            title: 'Nodes'
          }
        },
      ]
    },
  ]
})

export default router