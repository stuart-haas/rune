<template>
  <div class="layout">
    <!-- Header -->
    <div class="bg-gray-800 text-white mb-8">
      <div class="max-w-6xl mx-auto py-6 px-8">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <h1 class="text-4xl font-bold tracking-tight">Rune</h1>
            <span class="text-gray-500 text-3xl font-light">/</span>
            <h2 class="text-xl font-medium text-gray-300">{{ $route.meta.title || 'Dashboard' }}</h2>
          </div>
          
          <!-- User Status Icons -->
          <div class="flex items-center gap-4">
            <div class="text-gray-300 flex gap-4">
              <FontAwesomeIcon icon="user" class="h-5 w-5" />
              <FontAwesomeIcon icon="cog" class="h-5 w-5" />
            </div>
            <div class="h-2 w-2 rounded-full bg-green-500" title="Connected"></div>
          </div>
        </div>
      </div>
    </div>
    
    <router-view></router-view>
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