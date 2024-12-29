<template>
  <div class="mb-4 flex justify-between">
    <Button variant="default" @click="handleAdd">
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
        <CardDescription>User: {{ node.User || 'Not specified' }}</CardDescription>
        <CardDescription v-if="node.ExternalID">External Provider: {{ node.ExternalProvider }}</CardDescription>
        <CardDescription v-if="node.SyncedAt">Synced At: {{ formatDate(node.SyncedAt) }}</CardDescription>
        <div class="flex flex-wrap gap-2 mt-2">
          <span 
            v-for="tag in node.Tags" 
            :key="tag.Name"
            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
          >
            {{ tag.Name }}
          </span>
        </div>
      </CardHeader>
      <CardFooter class="flex justify-end space-x-2">
        <AlertDialog>
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" size="sm">
                <FontAwesomeIcon icon="ellipsis-vertical" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem @click="handleEdit(node)">
                <FontAwesomeIcon icon="pen-to-square" class="mr-2" />
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem v-if="node.ExternalID" @click="handleSync(node.ExternalID)">
                <FontAwesomeIcon icon="sync" class="mr-2" />
                Sync
              </DropdownMenuItem>
              <DropdownMenuItem>
                <AlertDialogTrigger>
                  <FontAwesomeIcon icon="trash" class="mr-2" />
                  <span class="text-red-600">Delete</span>
                </AlertDialogTrigger>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
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
          <TableHead>External Provider</TableHead>
          <TableHead>Synced At</TableHead>
          <TableHead>Tags</TableHead>
          <TableHead class="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="node in nodes" :key="node.id">
          <TableCell>{{ node.Hostname }}</TableCell>
          <TableCell>{{ node.User }}</TableCell>
          <TableCell>{{ node.ExternalID ? node.ExternalProvider : '-' }}</TableCell>
          <TableCell>{{ node.SyncedAt ? formatDate(node.SyncedAt) : 'Never' }}</TableCell>
          <TableCell>
            <div class="flex flex-wrap gap-2">
              <span 
                v-for="tag in node.Tags" 
                :key="tag.Name"
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
              >
                {{ tag.Name }}
              </span>
            </div>
          </TableCell>
          <TableCell class="text-right">
            <AlertDialog>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" size="sm">
                    <FontAwesomeIcon icon="ellipsis-vertical" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent>
                  <DropdownMenuItem @click="handleEdit(node)">
                    <FontAwesomeIcon icon="pen-to-square" class="mr-2" />
                    Edit
                  </DropdownMenuItem>
                  <DropdownMenuItem v-if="node.ExternalID" @click="handleSync(node.ExternalID)">
                    <FontAwesomeIcon icon="sync" class="mr-2" />
                    Sync
                  </DropdownMenuItem>
                  <DropdownMenuItem>
                    <AlertDialogTrigger>
                      <FontAwesomeIcon icon="trash" class="mr-2" />
                      <span class="text-red-600">Delete</span>
                    </AlertDialogTrigger>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
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
        <FormField v-slot="{ value }" name="Tags">
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
        <FormField v-slot="{ value }" name="Tags">
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
import { formatDate } from '@/lib/utils'
import { useTailscaleAPI } from '@/api/tailscale'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

const formSchema = object({
  Hostname: string().required('Hostname is required'),
  User: string().optional(),
  Tags: array().of(string()).nullable().default([]),
})

const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    Tags: [],
  }
})

const nodesApi = useNodesAPI()
const { data: nodes } = nodesApi.fetch()
const { mutateAsync: createNode } = nodesApi.create()
const { mutateAsync: updateNode } = nodesApi.update()
const { mutateAsync: deleteNode } = nodesApi.delete()

const tailscaleApi = useTailscaleAPI()
const { mutateAsync: syncDevice } = tailscaleApi.syncDevice()

const isGridView = ref(true)
const addingNode = ref(false)
const editingNode = ref(null)

const handleAdd = () => {
  form.setValues({ Tags: [] })
  addingNode.value = true
}

const handleEdit = (node) => {
  editingNode.value = node
  const formattedTags = node.Tags ? node.Tags.map(tag => tag.Name) : []
  form.setValues({ ...node, Tags: formattedTags })
}

const handleCreate = form.handleSubmit(async (payload) => {
  const formattedPayload = {
    ...payload,
    Tags: payload.Tags ? payload.Tags.map(tag => ({ Name: tag })) : []
  }
  await createNode(formattedPayload)
  addingNode.value = false
})

const handleUpdate = form.handleSubmit(async (payload) => {
  const formattedPayload = {
    ...payload,
    ID: editingNode.value.ID,
    Tags: payload.Tags ? payload.Tags.map(tag => ({ Name: tag })) : [],
  }
  await updateNode(formattedPayload)
  editingNode.value = null
})

const handleDelete = async (id) => {
  await deleteNode(id)
}

const handleSync = async (externalId) => {
  await syncDevice(externalId)
}
</script>