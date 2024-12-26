<template>
  <div class="max-w-6xl mx-auto px-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <Card v-for="node in nodes" :key="node.id" class="w-full">
        <CardHeader>
          <CardTitle>{{ node.Hostname }}</CardTitle>
          <CardDescription>User: {{ node.User }}</CardDescription>
          <CardDescription>Public Key: {{ node.PublicKey }}</CardDescription>
        </CardHeader>
        <CardFooter class="flex justify-end space-x-2">
          <Button variant="outline" @click="editNode(node)">
            <FontAwesomeIcon icon="pen-to-square" />
            Edit
          </Button>
          <Button variant="destructive" @click="deleteNode(node.id)">
            <FontAwesomeIcon icon="trash" />
            Delete
          </Button>
        </CardFooter>
      </Card>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useNodesAPI } from '@/api/nodes'
import {
  Card,
  CardHeader,
  CardFooter,
  CardTitle,
  CardDescription
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const isGridView = ref(true)
const showModal = ref(false)

const nodesApi = useNodesAPI()
const { data: nodes } = nodesApi.fetch()

const formatDate = (date) => {
  return new Date(date).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const editNode = (node) => {
  // TODO: Implement edit functionality
  console.log('Edit node:', node)
}

const deleteNode = (id) => {
  // TODO: Implement delete functionality
  console.log('Delete node:', id)
}
</script>