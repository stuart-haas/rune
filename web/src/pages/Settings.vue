<template>
  <Card>
    <CardHeader>
      <CardTitle>Tailscale Integration</CardTitle>
      <CardDescription v-if="devices">
        {{ devices.length }} {{ devices.length === 1 ? 'device' : 'devices' }} not synced
      </CardDescription>
    </CardHeader>
    <CardContent>
      <Skeleton v-if="checkDevicesLoading" class="w-full h-10 bg-gray-200" />
      <div v-else class="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
            <TableHead>ID</TableHead>
            <TableHead>Hostname</TableHead>
            <TableHead class="text-right">Actions</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody v-if="devices">
          <TableRow v-for="device in devices" :key="device.id">
            <TableCell>{{ device.id }}</TableCell>
            <TableCell>{{ device.hostname }}</TableCell>
            <TableCell class="text-right space-x-2">
              <Button variant="outline" size="sm" @click="syncDevice(device.id)">
                <FontAwesomeIcon icon="sync" />
                Sync
              </Button>
            </TableCell>
          </TableRow>
        </TableBody>
        <TableFooter v-else>
          <TableRow>
            <TableCell colspan="4">All devices synced</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
      </div>
    </CardContent>
    <CardFooter class="flex flex-row gap-2">
      <Button @click="syncDevices" :disabled="!devices || syncDevicesLoading">
        <FontAwesomeIcon icon="sync" class="mr-2" :class="{ 'fa-spin': syncDevicesLoading }" />
        Sync Devices
      </Button>
    </CardFooter>
  </Card>
</template>

<script setup>
import { Button } from '@/components/ui/button'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { Card, CardHeader, CardTitle, CardContent, CardDescription, CardFooter } from '@/components/ui/card'
import {
  Table,
  TableHeader,
  TableBody,
  TableHead,
  TableRow,
  TableCell,
  TableFooter,
} from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { useTailscaleAPI } from '@/api/tailscale'

const tailscaleApi = useTailscaleAPI()

const { data: devices, isPending: checkDevicesLoading } = tailscaleApi.checkDevices()
const { mutateAsync: syncDevices, isPending: syncDevicesLoading } = tailscaleApi.syncDevices()
const { mutateAsync: syncDevice, isPending: syncDeviceLoading } = tailscaleApi.syncDevice()
</script>