<template>
  <div class="max-w-6xl mx-auto px-8">
    <div class="mb-4 flex justify-end">
      <Button variant="outline" @click="isGridView = !isGridView">
        <FontAwesomeIcon :icon="isGridView ? 'table-list' : 'table-cells'" class="mr-2" />
        {{ isGridView ? 'Table View' : 'Grid View' }}
      </Button>
    </div>

    <div v-if="isGridView" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
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
          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button variant="destructive" size="sm">
                <FontAwesomeIcon icon="trash" />
                Delete
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Are you sure?</AlertDialogTitle>
                <AlertDialogDescription>
                  This action cannot be undone. This will permanently delete the node.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>Cancel</AlertDialogCancel>
                <AlertDialogAction class="bg-red-500" @click="deleteNode(node.id)">
                  <FontAwesomeIcon icon="trash" />
                  Continue
                </AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </CardFooter>
      </Card>
    </div>

    <div v-else class="rounded-md border">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Hostname</TableHead>
            <TableHead>User</TableHead>
            <TableHead>Public Key</TableHead>
            <TableHead class="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="node in nodes" :key="node.id">
            <TableCell>{{ node.Hostname }}</TableCell>
            <TableCell>{{ node.User }}</TableCell>
            <TableCell class="font-mono text-sm">{{ node.PublicKey }}</TableCell>
            <TableCell class="text-right space-x-2">
              <Button variant="outline" size="sm" @click="editNode(node)">
                <FontAwesomeIcon icon="pen-to-square" />
                Edit
              </Button>
              <AlertDialog>
                <AlertDialogTrigger asChild>
                  <Button variant="destructive" size="sm">
                    <FontAwesomeIcon icon="trash" />
                    Delete
                  </Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>Are you sure?</AlertDialogTitle>
                    <AlertDialogDescription>
                      This action cannot be undone. This will permanently delete the node.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <AlertDialogAction class="bg-red-500" @click="deleteNode(node.id)">
                      <FontAwesomeIcon icon="trash" />
                      Continue
                    </AlertDialogAction>
                  </AlertDialogFooter>
                </AlertDialogContent>
              </AlertDialog>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
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
import {
  Table,
  TableHeader,
  TableBody,
  TableHead,
  TableRow,
  TableCell,
} from '@/components/ui/table'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'

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

const nodeToDelete = ref(null)

const deleteNode = async (id) => {
  try {
    await nodesApi.delete(id)
    // Refresh the nodes list after deletion
    nodesApi.fetch()
  } catch (error) {
    console.error('Failed to delete node:', error)
    // You might want to show an error toast here
  }
}
</script>