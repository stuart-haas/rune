<template>
  <div class="mb-4 flex justify-between">
    <Button variant="default" @click="handleAdd">
      <FontAwesomeIcon icon="plus" class="mr-2" />
      New Task
    </Button>
  </div>

  <div class="rounded-md border">
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Command</TableHead>
          <TableHead>Created At</TableHead>
          <TableHead class="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="task in tasks" :key="task.id">
          <TableCell>{{ task.Command }}</TableCell>
          <TableCell>{{ task.CreatedAt ? formatDate(task.CreatedAt) : 'Never' }}</TableCell>
          <TableCell class="text-right">
            <AlertDialog>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" size="sm">
                    <FontAwesomeIcon icon="ellipsis-vertical" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent>
                  <DropdownMenuItem @click="handleEdit(task)">
                    <FontAwesomeIcon icon="pen-to-square" class="mr-2" />
                    Edit
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
                    This action cannot be undone. This will permanently delete the task.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction class="bg-red-500" @click="handleDelete(task.ID)">
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

  <Dialog :open="!!editingTask" @update:open="editingTask = null">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Edit Task</DialogTitle>
      </DialogHeader>
      <form>
        <FormField v-slot="{ componentField }" name="Command">
          <FormItem>
            <FormLabel>Command</FormLabel>
            <FormControl>
              <Input type="text" placeholder="ls -l" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
      <DialogFooter>
        <Button variant="outline" @click="editingTask = null">Cancel</Button>
        <Button type="submit" @click="handleUpdate">Update</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>

  <Dialog :open="addingTask" @update:open="addingTask = false">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Add Task</DialogTitle>
      </DialogHeader>
      <form>
        <FormField v-slot="{ componentField }" name="Command">
          <FormItem>
            <FormLabel>Command</FormLabel>
            <FormControl>
              <Input type="text" placeholder="ls -l" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </form>
      <DialogFooter>
        <Button variant="outline" @click="addingTask = false">Cancel</Button>
        <Button type="submit" @click="handleCreate">Create</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

<script setup>
import { ref } from 'vue'
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
import { formatDate } from '@/lib/utils'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { useTasksAPI } from '@/api/tasks'

const formSchema = object({
  Command: string().required('Command is required'),
})

const form = useForm({
  validationSchema: formSchema,
})

const tasksApi = useTasksAPI()
const { data: tasks } = tasksApi.fetch()
const { mutateAsync: createTask } = tasksApi.create()
const { mutateAsync: updateTask } = tasksApi.update()
const { mutateAsync: deleteTask } = tasksApi.delete()

const addingTask = ref(false)
const editingTask = ref(null)

const handleAdd = () => {

  addingTask.value = true
}

const handleEdit = (task) => {
  editingTask.value = task
  form.setValues(task)
}

const handleCreate = form.handleSubmit(async (payload) => {
  await createTask(payload)
  addingTask.value = false
})

const handleUpdate = form.handleSubmit(async (payload) => {
  await updateTask({ ...payload, ID: editingTask.value.ID })
  editingTask.value = null
})

const handleDelete = async (id) => {
  await deleteTask(id)
}
</script>