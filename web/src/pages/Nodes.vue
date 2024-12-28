<template>
  <div class="mb-4 flex justify-between">
    <Button variant="default" @click="addingNode = true">
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
        <CardTitle>{{ node.hostname }}</CardTitle>
        <CardDescription>User: {{ node.user }}</CardDescription>
        <CardDescription>Last Sync: {{ node.lastSync ? formatDate(node.lastSync) : 'Never' }}</CardDescription>
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
          <TableHead>Last Sync</TableHead>
          <TableHead class="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="node in nodes" :key="node.id">
          <TableCell>{{ node.Hostname }}</TableCell>
          <TableCell>{{ node.user }}</TableCell>
          <TableCell>{{ node.lastSync ? formatDate(node.lastSync) : 'Never' }}</TableCell>
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
        <FormField v-slot="{ value }" name="tags">
          <FormItem>
            <FormLabel>Tags</FormLabel>
            <FormControl>
              <TagsInput :model-value="value">
                <TagsInputItem v-for="item in value" :key="item" :value="item">
                  <TagsInputItemText />
                  <TagsInputItemDelete />
                </TagsInputItem>
                <TagsInputInput placeholder="Tags..." />
              </TagsInput>
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

  <Dialog :open="addingNode" @update:open="addingNode = false">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Add Node</DialogTitle>
      </DialogHeader>
      <form>
        <FormField v-slot="{ componentField }" name="Hostname">
          <FormItem>
            <FormLabel>Hostname</FormLabel>
            <FormControl>
              <Input type="text" placeholder="192.168.1.1" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="User">
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
        <Button variant="outline" @click="addingNode = false">Cancel</Button>
        <Button type="submit" @click="handleCreate">Create</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
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
import { TagsInput, TagsInputInput, TagsInputItem, TagsInputItemDelete, TagsInputItemText } from '@/components/ui/tags-input'
import { Input } from '@/components/ui/input'
import { FormField, FormItem, FormLabel, FormControl, FormMessage } from '@/components/ui/form'
import { useForm } from 'vee-validate'
import { array, object, string } from 'yup'

const formatDate = (dateString) => {
  try {
    const date = new Date(dateString)
    if (isNaN(date)) return 'Invalid Date'
    
    return new Intl.DateTimeFormat('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    }).format(date)
  } catch (error) {
    console.error('Error formatting date:', error)
    return 'Invalid Date'
  }
}

const formSchema = object({
  hostname: string().required('Hostname is required'),
  user: string().required('User is required'),
  tags: array().of(string()).nullable().default([]),
})

const form = useForm({
  validationSchema: formSchema,
})

const nodesApi = useNodesAPI()
const { data: nodes } = nodesApi.fetch()
const { mutateAsync: createNode } = nodesApi.create()
const { mutateAsync: updateNode } = nodesApi.update()
const { mutateAsync: deleteNode } = nodesApi.delete()

const isGridView = ref(true)
const addingNode = ref(false)
const editingNode = ref(null)

const handleEdit = (node) => {
  editingNode.value = node
  const formattedTags = node.tags ? node.tags.map(tag => tag.name) : []
  form.setValues({ ...node, tags: formattedTags })
}

const handleCreate = form.handleSubmit(async (payload) => {
  const formattedPayload = {
    ...payload,
    tags: payload.tags ? payload.tags.map(tag => ({ name: tag })) : []
  }
  await createNode(formattedPayload)
  addingNode.value = false
})

const handleUpdate = form.handleSubmit(async (payload) => {
  const formattedPayload = {
    ...payload,
    tags: payload.tags ? payload.tags.map(tag => ({ name: tag })) : [],
    id: editingNode.value.ID
  }
  await updateNode(formattedPayload)
  editingNode.value = null
})

const handleDelete = async (id) => {
  await deleteNode(id)
}
</script>