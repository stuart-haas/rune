<template>
  <div class="page">
    <div class="mb-4 flex justify-between">
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
                <AlertDialogAction class="bg-red-500" @click="handleDelete(node.ID)">
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
            <TableHead class="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="node in nodes" :key="node.id">
            <TableCell>{{ node.Hostname }}</TableCell>
            <TableCell>{{ node.User }}</TableCell>
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
                    <AlertDialogAction class="bg-red-500" @click="handleDelete(node.ID)">
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
        <form>
          <FormField v-slot="{ componentField }" name="hostname">
            <FormItem>
              <FormLabel>Hostname</FormLabel>
              <FormControl>
                <Input type="text" placeholder="192.168.1.1" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="user">
            <FormItem>
              <FormLabel>User</FormLabel>
              <FormControl>
                <Input type="text" placeholder="user" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>
        <DialogFooter>
          <Button variant="outline" @click="editingNode = null">Cancel</Button>
          <Button type="submit" @click="handleUpdate">Update</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>

    <Dialog :open="showNewNodeDialog" @update:open="showNewNodeDialog = false">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add Node</DialogTitle>
        </DialogHeader>
        <form>
          <FormField v-slot="{ componentField }" name="hostname">
            <FormItem>
              <FormLabel>Hostname</FormLabel>
              <FormControl>
                <Input type="text" placeholder="192.168.1.1" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="user">
            <FormItem>
              <FormLabel>User</FormLabel>
              <FormControl>
                <Input type="text" placeholder="user" v-bind="componentField" />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
        </form>
        <DialogFooter>
          <Button variant="outline" @click="showNewNodeDialog = false">Cancel</Button>
          <Button type="submit" @click="handleCreate">Create</Button>
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
import { FormField, FormItem, FormLabel, FormControl, FormMessage } from '@/components/ui/form'
import { useForm } from 'vee-validate'
import { object, string } from 'yup'

const formSchema = object({
  hostname: string().required('Hostname is required'),
  user: string().required('User is required'),
})

const form = useForm({
  validationSchema: formSchema,
})

const nodesApi = useNodesAPI()
const { data: nodes } = nodesApi.fetch()
const { mutateAsync: createNode } = nodesApi.create()
const { mutateAsync: updateNode } = nodesApi.update()
const { mutateAsync: deleteNode } = nodesApi.remove()

const isGridView = ref(true)
const showNewNodeDialog = ref(false)
const editingNode = ref(null)

const handleEdit = (node) => {
  editingNode.value = node
  form.setValues({
    hostname: node.Hostname,
    user: node.User,
  })
}

const handleCreate = form.handleSubmit(async (payload) => {
  try {
    await createNode(payload)
    showNewNodeDialog.value = false
  } catch (error) {
    console.error('Failed to create node:', error)
  }
})

const handleUpdate = form.handleSubmit(async (payload) => {
  console.log('Updating node:', payload)
  try {
    await updateNode({ ...payload, id: editingNode.value.ID })
    editingNode.value = null
  } catch (error) {
    console.error('Failed to update node:', error)
  }
})

const handleDelete = async (id) => {
  try {
    await deleteNode(id)
  } catch (error) {
    console.error('Failed to delete node:', error)
  }
}
</script>