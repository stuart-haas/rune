<template>
  <h1 class="text-2xl font-bold mb-6">Node Detail</h1>
  
  <div v-if="node" class="space-y-4">
    <div class="flex">
      <strong class="w-40 font-medium">Hostname:</strong>
      <span>{{ node.Hostname }}</span>
    </div>
    
    <div class="flex">
      <strong class="w-40 font-medium">User:</strong>
      <span>{{ node.User }}</span>
    </div>
    
    <div class="flex">
      <strong class="w-40 font-medium">External ID:</strong>
      <span>{{ node.ExternalID }}</span>
    </div>
    
    <div class="flex">
      <strong class="w-40 font-medium">External Provider:</strong>
      <span>{{ node.ExternalProvider }}</span>
    </div>
    
    <div class="flex">
      <strong class="w-40 font-medium">Synced At:</strong>
      <span>{{ new Date(node.SyncedAt).toLocaleString() }}</span>
    </div>
    
    <div class="flex">
      <strong class="w-40 font-medium">Tags:</strong>
      <div class="flex flex-wrap gap-2">
        <span 
          v-for="tag in node.Tags" 
          :key="tag.Name" 
          class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
        >
          {{ tag.Name }}
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useNodesAPI } from '@/api/nodes'

const route = useRoute()
const nodeId = computed(() => route.params.id)

const nodesApi = useNodesAPI()
const { data: node } = nodesApi.fetchNode(nodeId)
</script>