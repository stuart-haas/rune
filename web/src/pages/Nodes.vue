<template>
  <div class="max-w-6xl mx-auto px-8">
    <div class="mb-4 flex justify-end space-x-2">
      <Button variant="default" @click="showNewNodeDialog = true">
        <FontAwesomeIcon icon="plus" class="mr-2" />
        New Node
      </Button>
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
          <Button variant="outline" @click="handleEdit(node)">
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
                <AlertDialogAction class="bg-red-500" @click="handleDelete(node.id)">
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
              <Button variant="outline" size="sm" @click="handleEdit(node)">
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
                    <AlertDialogAction class="bg-red-500" @click="handleDelete(node.id)">
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

    <Dialog :open="!!editingNode" @update:open="editingNode = null">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Node</DialogTitle>
        </DialogHeader>
        <div class="grid gap-4 py-4">
          <div class="grid gap-2">
            <label for="hostname">Hostname</label>
            <Input id="hostname" v-model="editForm.hostname" />
          </div>
          <div class="grid gap-2">
            <label for="user">User</label>
            <Input id="user" v-model="editForm.user" />
          </div>
          <div class="grid gap-2">
            <label for="publicKey">Public Key</label>
            <Input id="publicKey" v-model="editForm.publicKey" />
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" @click="editingNode = null">Cancel</Button>
          <Button @click="handleSave">Save changes</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <Dialog :open="showNewNodeDialog" @update:open="showNewNodeDialog = false">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add New Node</DialogTitle>
        </DialogHeader>
        <div class="grid gap-4 py-4">
          <div class="grid gap-2">
            <label for="new-hostname">Hostname</label>
            <Input id="new-hostname" v-model="newNodeForm.hostname" />
          </div>
          <div class="grid gap-2">
            <label for="new-user">User</label>
            <Input id="new-user" v-model="newNodeForm.user" />
          </div>
          <div class="grid gap-2">
            <label for="new-publicKey">Public Key</label>
            <Input id="new-publicKey" v-model="newNodeForm.publicKey" />
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" @click="showNewNodeDialog = false">Cancel</Button>
          <Button @click="handleCreateNode">Create</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
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
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from '@/components/ui/dialog'
import { Input } from '@/components/ui/input'

const nodesApi = useNodesAPI()
const { data: nodes } = nodesApi.fetch()
const { mutateAsync: createNode } = nodesApi.create()
const { mutateAsync: updateNode } = nodesApi.update()
const { mutateAsync: deleteNode } = nodesApi.remove()

const isGridView = ref(true)
const showNewNodeDialog = ref(false)
const editingNode = ref(null)

const newNodeForm = ref({
  hostname: '',
  user: '',
  publicKey: ''
})

const editForm = ref({
  hostname: '',
  user: '',
  publicKey: ''
})

const handleEdit = (node) => {
  console.log('Editing node:', node)
  editingNode.value = node
  editForm.value = {
    hostname: node.Hostname,
    user: node.User,
    publicKey: node.PublicKey
  }
}

const handleSave = async () => {
  try {
    await updateNode({ ...editForm.value, id: editingNode.value.ID })
    editingNode.value = null
  } catch (error) {
    console.error('Failed to update node:', error)
  }
}

const handleDelete = async (id) => {
  try {
    await deleteNode(id)
  } catch (error) {
    console.error('Failed to delete node:', error)
  }
}

const handleCreateNode = async () => {
  try {
    await createNode(newNodeForm.value)
    showNewNodeDialog.value = false
    newNodeForm.value = { hostname: '', user: '', publicKey: '' }
  } catch (error) {
    console.error('Failed to create node:', error)
  }
}
</script>