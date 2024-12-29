import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import Nodes from '@/pages/Nodes.vue'
import Tasks from '@/pages/Tasks.vue'
import Settings from '@/pages/Settings.vue'
import Schedules from '@/pages/Schedules.vue'
import NodeDetail from '@/pages/NodeDetail.vue'
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
          },
        },
        {
          path: 'nodes/:id',
          name: 'NodeDetail',
          component: NodeDetail,
          meta: {
            title: 'Node Detail'
          }
        },
        {
          path: 'tasks',
          name: 'Tasks',
          component: Tasks,
          meta: {
            title: 'Tasks'
          }
        },
        {
          path: 'schedules',
          name: 'Schedules',
          component: Schedules,
          meta: {
            title: 'Schedules'
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