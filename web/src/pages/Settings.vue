<template>
  <div class="p-8">
    <div class="space-y-6">
      <div>
        <h3 class="text-lg font-medium">OAuth Clients</h3>
        <p class="text-sm text-muted-foreground">
          Manage your OAuth clients and data synchronization
        </p>
      </div>

      <!-- Client List -->
      <div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <Card v-for="client in oauthClients" :key="client.id">
          <CardHeader>
            <CardTitle>{{ client.name }}</CardTitle>
            <CardDescription>{{ client.description }}</CardDescription>
          </CardHeader>
          <CardContent>
            <Badge :variant="client.isConnected ? 'default' : 'secondary'">
              {{ client.isConnected ? 'Connected' : 'Not Connected' }}
            </Badge>
          </CardContent>
          <CardFooter class="flex justify-end gap-2">
            <Button 
              @click="syncData(client)" 
              :disabled="!client.isConnected"
              variant="secondary"
            >
              <FontAwesomeIcon icon="sync" class="mr-2 h-4 w-4" />
              Sync Data
            </Button>
            <Button 
              v-if="client.isConnected"
              @click="disconnectClient(client)" 
              variant="destructive"
            >
              Disconnect
            </Button>
            <Button 
              v-else
              @click="connectClient(client)"
            >
              Connect
            </Button>
          </CardFooter>
        </Card>

        <!-- Add Client Card -->
        <Card class="border-dashed">
          <CardHeader>
            <CardTitle>Add New Client</CardTitle>
            <CardDescription>
              Configure a new OAuth client for data synchronization
            </CardDescription>
          </CardHeader>
          <CardContent class="flex items-center justify-center p-6">
            <Button variant="outline" @click="showAddClientDialog = true">
              <FontAwesomeIcon icon="plus" class="mr-2 h-4 w-4" />
              Add OAuth Client
            </Button>
          </CardContent>
        </Card>
      </div>
    </div>

    <!-- Add Client Dialog -->
    <Dialog v-model:open="showAddClientDialog">
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add OAuth Client</DialogTitle>
          <DialogDescription>
            Enter the details for your new OAuth client
          </DialogDescription>
        </DialogHeader>
        <form @submit.prevent="addClient">
          <div class="grid gap-4 py-4">
            <div class="grid gap-2">
              <Label for="name">Name</Label>
              <Input id="name" v-model="newClient.name" required />
            </div>
            <div class="grid gap-2">
              <Label for="description">Description</Label>
              <Textarea 
                id="description" 
                v-model="newClient.description"
                placeholder="Describe the purpose of this client"
              />
            </div>
          </div>
          <DialogFooter>
            <Button type="button" variant="secondary" @click="showAddClientDialog = false">
              Cancel
            </Button>
            <Button type="submit">Add Client</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Textarea } from '@/components/ui/textarea'
import { Badge } from '@/components/ui/badge'

const oauthClients = ref([])
const showAddClientDialog = ref(false)
const newClient = ref({
  name: '',
  description: '',
})

const addClient = () => {
  // TODO: Implement API call to add new client
  oauthClients.value.push({
    id: Date.now(),
    ...newClient.value,
    isConnected: false,
  })
  showAddClientDialog.value = false
  newClient.value = { name: '', description: '' }
}

const connectClient = async (client) => {
  // TODO: Implement OAuth connection flow
  client.isConnected = true
}

const disconnectClient = async (client) => {
  // TODO: Implement disconnect logic
  client.isConnected = false
}

const syncData = async (client) => {
  // TODO: Implement data sync logic
  console.log(`Syncing data for client: ${client.name}`)
}
</script>
