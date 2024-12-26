<template>
  <div class="layout">
    <!-- Header -->
    <div class="bg-gray-800 text-white mb-8">
      <div class="max-w-6xl mx-auto py-6 px-8">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <h1 class="text-4xl font-bold tracking-tight">Rune</h1>
            
            <!-- Navigation Menu -->
            <nav class="ml-8">
              <ul class="flex gap-6">
                <li><router-link to="/" class="text-gray-300 hover:text-white">Dashboard</router-link></li>
                <li><router-link to="/nodes" class="text-gray-300 hover:text-white">Nodes</router-link></li>
                <li><router-link to="/tasks" class="text-gray-300 hover:text-white">Tasks</router-link></li>
                <li><router-link to="/settings" class="text-gray-300 hover:text-white">Settings</router-link></li>
              </ul>
            </nav>
          </div>
          
          <div class="flex items-center gap-4">
            <div class="text-gray-300 flex gap-4">
              <FontAwesomeIcon icon="user" class="h-5 w-5" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="max-w-6xl mx-auto px-8">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

export default {
  name: 'DefaultLayout',
  components: {
    FontAwesomeIcon
  }
}
</script>